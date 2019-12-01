package cli

import (
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

	// embed
	embed := cli.NewSubCommand("embed", "Embed all *.wav/*.aiff files as []byte")
	embed.Action(func() error {
		println(`** Under Construction **`)
		return nil
	})

	// serve
	serve := cli.NewSubCommand("serve", "Hot load your composition after save")
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

	// build a binary
	build := cli.NewSubCommand("build", "Build a stand-alone binary")
	build.Action(func() error {
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

func selectedOutputHelper(out int) RTout {
	if out == int(PortAudio) {
		return PortAudio
	}
	return Oto
}
