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
	"gopkg.in/yaml.v3"
)

const guiEndNote = `Note: Gui support is provided by Fyne library https://github.com/fyne-io/fyne
Documentation on how to use it can be found at https://fyne.io/
Examples and demos live at https://github.com/gocomu/cli/examples`

// newProject .
func newProject(projectType ProjectType, projectName string, rtOut RTout) error {
	var endNote string
	var ptype string
	var output string

	switch projectType {
	case Cli:
		ptype = "cli"
		endNote = ""
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
		return errors.New("-name flag can not be empty! try ie: \ngocomu new cli -name sampleProject")
	}

	// root
	// check if directory exists
	if _, err := os.Stat(projectName); !os.IsNotExist(err) {
		return errors.New("directory already exists")
	}
	// create project's root folder
	os.Mkdir(filepath.Join(dir, projectName), 0755)

	// create gocomu.yml
	data, err := yaml.Marshal(&GocomuYaml{
		Name:        projectName,
		Description: "Demo Project",
		Version:     "v0.0.0",
		Type:        ptype,
		ServeOutput: output,
	})
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Join(dir, projectName, "gocomu.yml"), data, 0755)
	if err != nil {
		return err
	}

	// generate sine.go
	templates.CreateFile(
		// path
		filepath.Join(dir, projectName),
		// file name
		"sine.go",
		// template
		templates.SineGo,
		// data
		&templates.Data{
			ProjectName: projectName,
		},
	)

	// cmd
	// create cmd folder
	os.Mkdir(filepath.Join(dir, projectName, "cmd"), 0755)

	// generate cmd/gocomu.go
	templates.CreateFile(
		// dir
		filepath.Join(dir, projectName, "cmd"),
		// file name
		"gocomu.go",
		// template
		templates.GOCOMUGo,
		// data
		&templates.Data{
			ProjectName: projectName,
		},
	)

	// create folder projectName/cmd/projectName
	os.Mkdir(filepath.Join(dir, projectName, "cmd", projectName), 0755)

	switch projectType {
	case Cli:
		// generate main.go
		templates.CreateFile(
			// path
			filepath.Join(dir, projectName, "cmd", projectName),
			// file name
			"main.go",
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
			filepath.Join(dir, projectName, "cmd", projectName),
			// file name
			"main.go",
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
	os.Mkdir(filepath.Join(dir, projectName, "embed"), 0755)
	// generate embed.go
	embedgo, _ := os.Create(filepath.Join(dir, projectName, "embed", "fs.go"))
	defer embedgo.Close()
	t := template.Must(template.New("fsgo").Parse(templates.FSGo))
	t.Execute(embedgo, &templates.Data{
		ProjectName: projectName,
	})
	embedgo.Sync()

	// create output folder
	os.Mkdir(filepath.Join(dir, projectName, "output"), 0755)

	// run go mod init projectName
	projectDir := filepath.Join(dir, projectName)
	os.Chdir(projectDir)
	cmd := exec.Command("go", "mod", "init", projectName)
	err = cmd.Run()
	if err != nil {
		return err
	}

	// run go mod tidy
	cmd = exec.Command("go", "mod", "tidy")
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

`, ptype, projectName, output, "cd "+projectName+"/", "gocomu serve", "gocomu -help", "gocomu", endNote)
	return nil
}
