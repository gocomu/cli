package cli

import (
	"fmt"

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

	// real-time record
	record := cli.NewSubCommand("record", "Render the output as wav/aiff")
	record.Action(func() error {
		// err := Render(dur, form)
		// return err
		println(`** Under Construction **`)
		return nil
	})

	// offline render
	var dur int
	var form int
	offline := cli.NewSubCommand("offline", "Render the output as wav/aiff")
	offline.IntFlag("duration", "Duration in minutes of render", &dur)
	offline.IntFlag("format", "0 = wav, 1 = aiff", &form)
	offline.Action(func() error {
		// err := Render(dur, form)
		// return err
		println(`** Under Construction **`)
		return nil
	})

	// embed
	embed := cli.NewSubCommand("embed", "Embed all *.wav/*.aiff files as []byte")
	embed.Action(func() error {
		err := embedAudio()
		return err
	})

	// build a binary
	build := cli.NewSubCommand("build", "Build a stand-alone binary")
	build.Action(func() error {
		err := buildApp()
		return err
	})

	// update gocomu
	update := cli.NewSubCommand("update", "Update GOCOMU to latest version")
	update.Action(func() error {
		err := updateGocomu()
		return err
	})

	// print version
	version := cli.NewSubCommand("version", "Returns binary's version")
	version.Action(func() error {
		fmt.Println(cli.Version())
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
