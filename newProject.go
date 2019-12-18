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
	"gopkg.in/yaml.v3"
)

const cliEndNote = `Note: Cli support is provided by Clir library https://github.com/leaanthony/clir
Documantion on how to use it can be found at https://clir.leaanthony.com/
However you can opt to use any other CLI library like cobra, urfave/cli etc..
Examples and demos live at https://github.com/gocomu/cli/examples`

const guiEndNote = `Note: Gui support is provided by Fyne library https://github.com/fyne-io/fyne
Documantion on how to use it can be found at https://fyne.io/
Examples and demos live at https://github.com/gocomu/cli/examples`

func newProject(projectType ProjectType, projectName string, rtOut RTout) error {
	yellow := color.FgYellow.Render
	green := color.FgGreen.Render
	magenta := color.FgWhite.Render

	var endNote string
	var ptype string
	var output string

	switch projectType {
	case Cli:
		ptype = "cli"
		endNote = cliEndNote
	case Gui:
		ptype = "gui"
		endNote = guiEndNote
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
	err := ioutil.WriteFile(projectName+"/gocomu.yml", data, 0755)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// generate sine.go
	templates.CreateFile(
		// path
		projectName,
		// file name
		"/sine.go",
		// template
		templates.SineGo,
		// data
		&templates.Data{
			ProjectName: projectName,
		},
	)

	// cmd
	// create cmd folder
	os.Mkdir(projectName+"/cmd", 0755)

	// generate cmd/gocomu.go
	templates.CreateFile(
		// path
		projectName+"/cmd",
		// file name
		"/gocomu.go",
		// template
		templates.GOCOMUGo,
		// data
		&templates.Data{
			ProjectName: projectName,
		},
	)

	// create folder projectName/cmd/projectName
	os.Mkdir(projectName+"/cmd/"+projectName, 0755)

	switch projectType {
	case Cli:
		// generate main.go
		templates.CreateFile(
			// path
			projectName+"/cmd/"+projectName,
			// file name
			"/main.go",
			// template
			templates.CliTemplateMainGo,
			// data
			&templates.Data{
				ProjectName: projectName,
			},
		)

	case Gui:
		// generate main.go
		templates.CreateFile(
			// path
			projectName+"/cmd/"+projectName,
			// file name
			"/main.go",
			// template
			templates.GuiMainGo,
			// data
			&templates.Data{
				ProjectName: projectName,
			},
		)
	}

	// embed
	// create embed folder
	os.Mkdir(projectName+"/embed", 0755)
	// generate embed.go
	embedgo, _ := os.Create(dir + "/" + projectName + "/embed/fs.go")
	defer embedgo.Close()
	t := template.Must(template.New("fsgo").Parse(templates.FSGo))
	t.Execute(embedgo, &templates.Data{
		ProjectName: projectName,
	})
	embedgo.Sync()

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

More guides and tutorials on usage of your newly created project
exist on comu library's wiki at https://github.com/gocomu/comu/wiki/cli

%s

`, green(ptype), green(projectName), green(output), yellow("cd "+projectName+"/"), yellow("gocomu serve"), yellow("gocomu -help"), green("gocomu"), magenta(endNote))
	return nil
}
