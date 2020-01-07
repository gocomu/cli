package cli

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gocomu/cli/templates"
	"gopkg.in/yaml.v2"
)

// fsnotify
var watcher *fsnotify.Watcher

// reaload chan
var trigger = make(chan bool)

// data holder for templates arguments
var yamlData GocomuYaml

func projectServe() error {
	// check if gocomu.yml is present inside current dir
	dir, _ := os.Getwd()

	data, err := ioutil.ReadFile(dir + "/gocomu.yml")
	if err != nil {
		data, err = ioutil.ReadFile(dir + "/gocomu.yaml")
		if err != nil {
			return errors.New("wrong directory. gocomu.yml file not found")
		}

	}

	yaml.Unmarshal(data, &yamlData)
	// do stuff with yamlData
	// then continue
	timeStarted := time.Now()
	// creates a new file watcher
	watcher, _ = fsnotify.NewWatcher()
	defer watcher.Close()
	// starting at the root of the project,
	// walk each file/directory searching for directories
	if err := filepath.Walk(dir, watchDir); err != nil {
		return err
	}

	// remove cmd/projectName dir from watcher
	// TODO: the way it is now you seems u can't edit files inside cmd dir, fix or rationalise
	watcher.Remove(dir + "/cmd/" + yamlData.Name)
	watcher.Remove(dir + "/go.mod")
	watcher.Remove(dir + "/go.sum")

	// create a blocking channel
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, os.Kill)
	block := make(chan bool)

	// generate serve.go
	// templates.CreateFile("cmd/"+yamlData.Name, "/gocomuServe.go", templates.ServeGo, &templates.Data{ProjectName: yamlData.Name})

	fmt.Printf(`
Started serving %s
at %s

`, yamlData.Name, timeStarted)

	go func() {
		// init reload()
		reload(yamlData.Name)
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					trigger <- true
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					trigger <- true
				}
				if event.Op&fsnotify.Rename == fsnotify.Rename {
					trigger <- true
				}
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					trigger <- true
				}
			case err, ok := <-watcher.Errors:
				fmt.Println("error:", err)
				if !ok {
					return
				}

			case <-done:
				// delete serve.go
				// dir, _ := os.Getwd()
				// os.Remove(dir + "/cmd/" + yamlData.Name + "/gocomuServe.go")

				fmt.Printf(`

stoped serving
time elapsed %s 

`, time.Now().Sub(timeStarted))

				os.Exit(0)
				return
			}
		}
	}()

	// block
	<-block

	return nil
}

// watchDir gets run as a walk func, searching for directories to add watchers to
func watchDir(path string, fi os.FileInfo, err error) error {
	// since fsnotify can watch all the files in a directory, watchers only need
	// to be added to each nested directory
	if fi.Mode().IsDir() {
		return watcher.Add(path)
	}

	return nil
}

// arg: yamlData.Name
func reload(name string) {
	blockingChan := make(chan bool)
	go func() {
		for {
			// generate gocomuServe.go
			templates.CreateFile("cmd/"+name, "/gocomuServe.go", templates.ServeGo, &templates.Data{ProjectName: name})

			// run 'go run -tags serve ./cmd/{{projectname}}/gocomuServe.go'
			cmd := exec.Command("go", "run", "-tags", "serve", "./cmd/"+name+"/gocomuServe.go")
			go func() {
				time.Sleep(1500 * time.Millisecond)
				dir, _ := os.Getwd()
				os.Remove(dir + "/cmd/" + name + "/gocomuServe.go")

				<-trigger

				fmt.Println("reloading...")
				// kill it
				if err := cmd.Process.Kill(); err != nil {
					//fmt.Println("wasn't running")
				}
				blockingChan <- true
			}()
			cmd.Run()
			<-blockingChan
		}
	}()
}
