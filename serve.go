package cli

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gookit/color"
	"gopkg.in/yaml.v2"
)

var watcher *fsnotify.Watcher

func projectServe() error {
	// check if gocomu.yml is present inside current dir
	dir, _ := os.Getwd()

	data, err := ioutil.ReadFile(dir + "/gocomu.yml")
	if err != nil {
		data, err = ioutil.ReadFile(dir + "/gocomu.yaml")
		if err != nil {
			color.Warn.Printf(`
Wrong directory.
No yaml file found!

`)
			return err
		}

	}

	var yamlData GocomuYaml
	yaml.Unmarshal(data, &yamlData)

	// do stuff with yamlData
	// then continue

	timeStarted := time.Now()
	fmt.Printf(`
Started serving %s
at %s

`, yamlData.Name, timeStarted)

	// creates a new file watcher
	watcher, _ = fsnotify.NewWatcher()
	defer watcher.Close()

	// starting at the root of the project,
	// walk each file/directory searching for directories
	if err := filepath.Walk(dir, watchDir); err != nil {
		fmt.Println("ERROR", err)
		return nil
	}

	// create a blocking channel
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, os.Kill)
	blcok := make(chan bool)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				fmt.Printf("RELOADING!")
				fmt.Println("event:", event)
				// if event.Op&fsnotify.Write == fsnotify.Write {
				// 	fmt.Println("modified file:", event.Name)
				// }
				// if event.Op&fsnotify.Remove == fsnotify.Remove {
				// 	fmt.Println("removed file:", event.Name)
				// }
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
