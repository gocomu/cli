package cli

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/gocomu/cli/templates"
	"github.com/gookit/color"
	"gopkg.in/yaml.v2"
)

// RTout type helps cli's -out flag
type RTout int

const (
	// PortAudio output
	PortAudio RTout = iota
	// Oto output
	Oto
)

// ProjectType type helps cli's `cli` subcommand
type ProjectType int

const (
	// Cli project type
	Cli = ProjectType(iota)
	// Gui project type
	Gui
)

// TemplatesVariables holds values to be substituted
// when creating/executing biolerplate templates
type TemplatesVariables struct {
	ProjectName string
	ProjectType ProjectType
	Output      RTout
}

func newProject(projectType ProjectType, projectName string, rtOut RTout) error {
	yellow := color.FgYellow.Render
	green := color.FgGreen.Render
	magenta := color.FgLightMagenta.Render

	var ptype string
	var output string

	switch projectType {
	case Cli:
		ptype = "cli"
	case Gui:
		ptype = "gui"
	}

	switch rtOut {
	case PortAudio:
		output = "PortAudio"
	case Oto:
		output = "Oto"
	}

	// check if projectName flag is empty
	if projectName == "" {
		fmt.Printf(`
-name flag can not be empty! try ie:`)
		color.Green.Printf(`
gocomu new %s -name sampleProject

`, ptype)
		fmt.Println(`if you need to see all available flags run`)
		color.Green.Printf(`gocomu new %s -help

`, ptype)
		return nil
	}

	dir, _ := os.Getwd()

	// root
	// check if directory exists
	if _, err := os.Stat(projectName); !os.IsNotExist(err) {
		return errors.New("directory already exists")
	}
	// create project's root folder
	os.Mkdir(projectName, 0755)

	// create gocomu.yml
	data, _ := yaml.Marshal(&GocomuYaml{
		Name:        projectName,
		Description: "Demo Project",
		Version:     "v0.0.0",
		Type:        ptype,
		ServeOutput: output,
	})
	err := ioutil.WriteFile("/gocomu.yml", data, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// cmd
	// create cmd folder
	os.Mkdir(projectName+"/cmd", 0755)
	// create a folder named projectName
	os.Mkdir(projectName+"/cmd/"+projectName, 0755)
	// generate main.go
	switch projectType {
	case Cli:
		maingo, _ := os.Create(projectName + "/cmd/" + projectName + "/main.go")
		defer maingo.Close()
		t := template.Must(template.New("maingo").Parse(templates.CliTemplateMainGo))
		t.Execute(maingo, &TemplatesVariables{
			ProjectName: projectName,
		})
		maingo.Sync()

		ptype = "CLI"
	case Gui:
	}

	// embed
	// create embed folder
	os.Mkdir(projectName+"/embed", 0755)
	// generate embed.go & emdedded.go
	embedgo, _ := os.Create(projectName + "/embed/embed.go")
	defer embedgo.Close()
	t := template.Must(template.New("embedgo").Parse(templates.EmbedGo))
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
	err = cmd.Run()
	if err != nil {
		return err
	}

	fmt.Printf(`

New %s project created!

Project Name: %s
Sound Output: %s

To serve the project and get hot reloading everytime 
you create, save or remove files from project folder
change into project directory "%s" and run:

%s

To see all available commands run

%s

or visit https://github.com/gocomu/cli/ and go through 
the README for extensive documentation on how to use %s 

More guides and tutorials on usage of your newly craeted project
exist on comu library's wiki at https://github.com/gocomu/comu/wiki/cli

%s

`, green(ptype), green(projectName), green(output), yellow("cd "+projectName+"/"), yellow("gocomu serve"), yellow("gocomu -help"), green("gocomu"), magenta(`Note: Cli support is provided by Clir library https://github.com/leaanthony/clir
Documantion on how to use it can be found at https://clir.leaanthony.com/
However you can opt to use any other CLI library like cobra, urfave/cli etc..
Check out more available libraries at awesome go https://awesome-go.com/#command-line
Examples and demos live at https://github.com/gocomu/cli/examples`))
	return nil
}
