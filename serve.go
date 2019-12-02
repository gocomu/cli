package cli

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gocomu/cli/templates"
	"github.com/gookit/color"
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
			color.Warn.Printf(`
Wrong directory.
gocomu.yml file not found!

`)
			return err
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
		fmt.Println("ERROR", err)
		return nil
	}

	// remove cmd/projectName dir from watcher
	watcher.Remove(dir + "/cmd/" + yamlData.Name)

	// create a blocking channel
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, os.Kill)
	blcok := make(chan bool)

	fmt.Printf(`
	Started serving %s
	at %s
	
	`, yamlData.Name, timeStarted)

	go func() {
		go reload(yamlData.Name)
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
				if !ok {
					return
				}
				fmt.Println("error:", err)

			case <-done:
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
	<-blcok

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
	for {
		// generate serve.go
		templates.CreateFile("cmd/"+name, "/serve.go", templates.ServeGo, &templates.Data{ProjectName: name})
		// run go run -tags serve ./cmd/{{projectname}}
		// Start a process:
		cmd := exec.Command("go", "run", "-tags", "serve", "./cmd/serve.go")
		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}
		// delete serve.go
		os.Remove("cmd/" + name + "/serve.go")
		<-trigger
		fmt.Println("reloading...")
		// Kill it:
		if err := cmd.Process.Kill(); err != nil {
			log.Fatal("failed to kill process: ", err)
		}
	}
}

// go run
// cmd := exec.Command("go", "run", "./cmd/"+yamlData.Name, "play")
// var stdout, stderr bytes.Buffer
// cmd.Stdout = &stdout
// cmd.Stderr = &stderr
// err = cmd.Run()
// if err != nil {
// 	log.Fatalf("cmd.Run() failed with %s\n", err)
// }
// outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
// fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)

// binary, lookErr := exec.LookPath("go")
// if lookErr != nil {
// 	panic(lookErr)
// }

// args := []string{"go", "run", "./cmd/" + yamlData.Name, "play"}
// env := os.Environ()
// execErr := syscall.Exec(binary, args, env)
// if execErr != nil {
// 	panic(execErr)
// }

// go run
// func goRun(projectName string) {
// 	cmd := exec.Command("go", "run", "./cmd/"+projectName, "play")
// 	var stdout, stderr bytes.Buffer
// 	cmd.Stdout = &stdout
// 	cmd.Stderr = &stderr
// 	err := cmd.Run()
// 	if err != nil {
// 		log.Fatalf("cmd.Run() failed with %s\n", err)
// 	}
// 	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
// 	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
// }
