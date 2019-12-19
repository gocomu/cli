package templates

// CliTemplateMainGo holds cli template main.go
const CliTemplateMainGo = `package main

import (
	"{{ .ProjectName }}/cmd"

	gocomu "github.com/gocomu/cli"
	"github.com/leaanthony/clir"
)

func main() {
	// load project info details from gocomu.yml located at proejct root
	yamlData, _ := gocomu.Yaml()
	clir := clir.NewCli(yamlData.Name, yamlData.Description, yamlData.Version)
	play := clir.NewSubCommand("play", "Start making sound")
	play.Action(func() error {

		cmd.GOCOMU()
		return nil
	})

	clir.Run()
    
}
`
