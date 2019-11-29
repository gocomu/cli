package cli

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/gocomu/cli/templates"
)

type RTout int

const (
	PortAudio RTout = iota
	Oto
)

type ProjectType int

const (
	Cli = ProjectType(iota)
	Cui
	Gui
)

type TemplatesVariables struct {
	ProjectName string
	ProjectType ProjectType
	Output      RTout
}

func newProject(projectType ProjectType, projectName string, rtOut RTout) error {
	dir, _ := os.Getwd()

	// root
	// check if directory exists
	if _, err := os.Stat(projectName); !os.IsNotExist(err) {
		return errors.New("directory already exists")
	}
	// create project's root folder
	os.Mkdir(projectName, 0755)

	// create gocomu.yml
	yml, _ := os.Create(projectName + "/gocomu.yml")
	defer yml.Close()
	t := template.Must(template.New("gocomuyaml").Parse(templates.GocomuYaml))
	t.Execute(yml, &TemplatesVariables{
		ProjectName: projectName,
	})
	yml.Sync()

	// // cmd
	// create cmd folder
	os.Mkdir(projectName+"/cmd", 0755)
	// create a folder named projectName
	os.Mkdir(projectName+"/cmd/"+projectName, 0755)
	// generate main.go
	switch projectType {
	case Cli:
		maingo, _ := os.Create(projectName + "/cmd/" + projectName + "/main.go")
		defer maingo.Close()
		t = template.Must(template.New("maingo").Parse(templates.CliTemplateMainGo))
		t.Execute(maingo, &TemplatesVariables{
			ProjectName: projectName,
		})
		maingo.Sync()
	case Cui:
	case Gui:
	}

	// embed
	// create embed folder
	os.Mkdir(projectName+"/embed", 0755)
	// generate embed.go & emdedded.go
	embedgo, _ := os.Create(projectName + "/embed/embed.go")
	defer embedgo.Close()
	t = template.Must(template.New("embedgo").Parse(templates.EmbedGo))
	t.Execute(embedgo, &TemplatesVariables{
		ProjectName: projectName,
	})
	embedgo.Sync()

	embeddedgo, _ := os.Create(projectName + "/embed/embeddedgo.go")
	defer embeddedgo.Close()
	t = template.Must(template.New("embeddedgo").Parse(templates.EmbeddedGo))
	t.Execute(embeddedgo, &TemplatesVariables{
		ProjectName: projectName,
	})
	embeddedgo.Sync()

	// create output folder
	os.Mkdir(projectName+"/output", 0755)

	// run go mod init projectName
	projectDir := filepath.Join(dir, projectName)
	os.Chdir(projectDir)
	cmd := exec.Command("go", "mod", "init", projectName)
	err := cmd.Run()
	if err != nil {
		return err
	}

	fmt.Println(`New project created!`)
	return nil
}
