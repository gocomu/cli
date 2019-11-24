package main

import (
	"fmt"

	gocomu "github.com/gocomu/cli"
	"github.com/leaanthony/clir"
)

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

func main() {

	// Create new cli
	cli := clir.NewCli("GOCOMU", "GOCOMU CLI", "v0.0.1")

	// Set the custom banner
	cli.SetBannerFunction(customBanner)

	// new
	new := cli.NewSubCommand("new", "Create New Project")

	// new cli
	var projectName string
	newCLI := new.NewSubCommand("cli", "New CLI Project")
	newCLI.StringFlag("name", "Project name", &projectName)
	newCLI.Action(func() error {
		if projectName == "" {
			fmt.Println(`Please fill the project name flag
   ie. gocomu new cli -name sampleProject
			`)
			return nil
		}
		err := gocomu.NewProject(projectName, "cli", 0)
		if err != nil {
			fmt.Println(err)
		}
		return nil
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

	// Run!
	cli.Run()

}
