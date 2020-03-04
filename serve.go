package cli

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gocomu/cli/templates"
)

// fsnotify watcher
var watcher *fsnotify.Watcher

// reload chan
var trigger = make(chan bool, 1)

// watchDir gets run as a walk func, searching for directories to add watchers to
func watchDir(path string, fi os.FileInfo, err error) error {
	// since fsnotify can watch all the files in a directory, watchers only need
	// to be added to each nested directory
	if fi.Mode().IsDir() {
		return watcher.Add(path)
	}

	return nil
}

// reload is a for loop that generates, runs and immediately erases `gocomuServe.go`
// then blocks (<-trigger channel) for fsnotify event to take place,
// at which point kills the `go run` command and repeats the process.
func reload(name string) {
	for {
		// generate gocomuServe.go
		templates.CreateFile(filepath.Join(dir, "cmd", name), "gocomuServe.go", templates.ServeGo, &templates.Data{ProjectName: name})

		// run 'go run -tags serve ./cmd/{{projectname}}/gocomuServe.go'
		cmd := exec.Command("go", "run", "-tags", "serve", filepath.Join(dir, "cmd", name, "gocomuServe.go"))
		cmd.Start()
		time.Sleep(1000 * time.Millisecond)
		os.Remove(filepath.Join(dir, "cmd", name, "gocomuServe.go"))

		<-trigger

		fmt.Println("reloading...")
		// kill it
		if err := cmd.Process.Kill(); err != nil {
			fmt.Println(err)
		}
	}
}

func events(timeStarted time.Time) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, os.Kill)

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			if event.Op&fsnotify.Create == fsnotify.Create || event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Rename == fsnotify.Rename || event.Op&fsnotify.Remove == fsnotify.Remove {
				trigger <- true
			}
		case err, ok := <-watcher.Errors:
			fmt.Println("error:", err)
			if !ok {
				return
			}

		case <-done:
			fmt.Printf(`

Stopped
time elapsed: %s 

`, time.Now().Sub(timeStarted))

			os.Exit(0)
			return
		}
	}
}

// projectServe .
func projectServe() error {
	yamlData, _ := Yaml()

	timeStarted := time.Now()
	// creates a new file watcher
	watcher, _ = fsnotify.NewWatcher()
	defer watcher.Close()

	// starting at the root of the project,
	// walk each file/directory searching for directories
	if err := filepath.Walk(dir, watchDir); err != nil {
		fmt.Println(err)
		return err
	}

	// remove cmd/projectName dir from watcher
	// TODO: the way it is now, seems u can't edit files inside cmd dir, fix or rationalize
	watcher.Remove(filepath.Join(dir, "cmd", yamlData.Name))
	// watcher.Remove(dir + "/cmd/" + yamlData.Name + "/gocomuServe.go")
	watcher.Remove(filepath.Join(dir, "go.mod"))
	watcher.Remove(filepath.Join(dir, "go.sum"))

	fmt.Printf(`
Serving %s
started: %s

`, yamlData.Name, timeStarted)

	// init reload()
	go reload(yamlData.Name)
	time.Sleep(1500 * time.Millisecond)

	// fsnotify events
	go events(timeStarted)

	// create a blocking channel
	block := make(chan bool)
	<-block

	return nil
}
