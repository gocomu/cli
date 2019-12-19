package templates

// CliTemplateMainGo holds cli template main.go
const CliTemplateMainGo = `package main

import (
	"flag"
	"fmt"
	"{{ .ProjectName }}/cmd"

	gocomu "github.com/gocomu/cli"
)

func main() {
	// a boolean flag to start playing the song
	// if no flag is present when cli starts it will print out an err/help message
	play := flag.Bool("play", false, "Start playing the demo song")
	flag.Parse()

	// load project info details from gocomu.yml located at proejct root
	yamlData, err := gocomu.Yaml()
	if err != nil {
		fmt.Println(err)
		return
	}

	if *play == true {
		cmd.GOCOMU()
	} else {
		fmt.Println("Project: ", yamlData.Name, " Description: ", yamlData.Description)
		fmt.Println("\nyou need to pass a '-play' flag")
	}
}
`
