package cli

import (
	"errors"
	"fmt"

	"github.com/leaanthony/clir"
)

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
	newCLI.IntFlag("out", "Choose between Port-audio and Oto for audio output", &out)
	newCLI.Action(func() error {
		err := prepareNewProject("cli", projectName, out)
		if err != nil {
			fmt.Println(err)
		}
		return err
	})

	// new cui
	newCUI := new.NewSubCommand("cui", "New CUI Project")
	newCUI.StringFlag("name", "Project name", &projectName)
	newCUI.Action(func() error {
		println(`** Under Construction **`)
		return nil
	})

	// new gui
	newGUI := new.NewSubCommand("gui", "New GUI Project")
	newGUI.StringFlag("name", "Project name", &projectName)
	newCUI.Action(func() error {
		println(`** Under Construction **`)
		return nil
	})

	// embed
	embed := cli.NewSubCommand("embed", "Embed all *.wav/*.aiff files as []byte")
	embed.Action(func() error {
		println(`** Under Construction **`)
		return nil
	})

	// serve
	serve := cli.NewSubCommand("serve", "Hot load your composition after save")
	serve.Action(func() error {
		println(`** Under Construction **`)
		return nil
	})

	// offline render
	offline := cli.NewSubCommand("offline", "Render the output as wav/aiff")
	offline.Action(func() error {
		println(`** Under Construction **`)
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

func prepareNewProject(projectType, projectName string, out int) error {
	// check if projectName flag is empty
	if projectName == "" {
		// if so, print instruction and return
		return errors.New(`Please fill the project name flag
		ie. gocomu new cli -name sampleProject
				`)
	}

	// assign ProjectType to subcommand
	var ptype ProjectType
	switch projectType {
	case "cli":
		ptype = Cli
	case "cui":
		ptype = Cui
	case "gui":
		ptype = Gui
	}

	// assign RTout to flag
	var output RTout
	switch out {
	case 0:
		output = PortAudio
	case 1:
		output = Oto
	}

	err := newProject(ptype, projectName, output)
	return err
}
