package cli

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/leaanthony/clir"
)

// ClirActions cli handlers
func ClirActions(cli *clir.Cli) {
	// Set the custom banner
	cli.SetBannerFunction(customBanner)

	// new
	new := cli.NewSubCommand("new", "Create New Project")

	// new cli
	var projectName string
	var out int
	newCLI := new.NewSubCommand("cli", "New CLI Project")
	newCLI.StringFlag("name", "Project name", &projectName)
	newCLI.IntFlag("out", "Choose sound output. 0 = PortAudio, 1 = Oto", &out)
	newCLI.Action(func() error {

		selectedOut := selectedOutputHelper(out)

		err := newProject(Cli, projectName, selectedOut)
		return err
	})

	// new gui
	newGUI := new.NewSubCommand("gui", "New GUI Project")
	newGUI.StringFlag("name", "Project name", &projectName)
	newGUI.Action(func() error {
		selectedOut := selectedOutputHelper(out)

		err := newProject(Gui, projectName, selectedOut)
		return err
	})

	// serve
	serve := cli.NewSubCommand("serve", "Hot load your composition while working")
	serve.Action(func() error {
		err := projectServe()
		return err
	})

	// offline render
	offline := cli.NewSubCommand("offline", "Render the output as wav/aiff")
	offline.Action(func() error {
		println(`** Under Construction **`)
		return nil
	})

	// embed
	embed := cli.NewSubCommand("embed", "Embed all *.wav/*.aiff files as []byte")
	embed.Action(func() error {
		println(`** Under Construction **`)
		return nil
	})

	// build a binary
	build := cli.NewSubCommand("build", "Build a stand-alone binary")
	build.Action(func() error {
		fmt.Println("Starting Building")
		binary, lookErr := exec.LookPath("go")
		if lookErr != nil {
			panic(lookErr)
		}

		yamlData := Yaml()
		args := []string{"go", "build", "-o", "output/" + yamlData.Name, "./cmd/" + yamlData.Name}
		env := os.Environ()
		execErr := syscall.Exec(binary, args, env)
		if execErr != nil {
			panic(execErr)
		}

		fmt.Printf(`
Sucess! File %s can be found inside output/ directory
`, yamlData.Name)
		return nil
	})
}

func customBanner(cli *clir.Cli) string {

	return `
      ::::::::       ::::::::       ::::::::       ::::::::         :::   :::      :::    ::: 
    :+:    :+:     :+:    :+:     :+:    :+:     :+:    :+:       :+:+: :+:+:     :+:    :+:  
   +:+            +:+    +:+     +:+            +:+    +:+      +:+ +:+:+ +:+    +:+    +:+   
  :#:            +#+    +:+     +#+            +#+    +:+      +#+  +:+  +#+    +#+    +:+    
 +#+   +#+#     +#+    +#+     +#+            +#+    +#+      +#+       +#+    +#+    +#+     
#+#    #+#     #+#    #+#     #+#    #+#     #+#    #+#      #+#       #+#    #+#    #+#      
########       ########       ########       ########       ###       ###     ########        

  ` + cli.Version() + " - " + cli.ShortDescription()
}

func selectedOutputHelper(out int) RTout {
	if out == int(PortAudio) {
		return PortAudio
	}
	return Oto
}
