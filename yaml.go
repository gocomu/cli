package cli

import (
	"errors"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// GocomuYaml .
type GocomuYaml struct {
	Name        string
	Description string
	Version     string
	Type        string
	ServeOutput string
}

// Yaml reads & returns content from gocomu.yml
// a secondary use is to infer wether we are inside project dir
func Yaml() (*GocomuYaml, error) {
	// check if gocomu.yml is present inside current dir
	gocomuYml := filepath.Join(dir, "gocomu.yml")
	gocomuYaml := filepath.Join(dir, "gocomu.yaml")

	data, err := ioutil.ReadFile(gocomuYaml)
	if err != nil {
		data, err = ioutil.ReadFile(gocomuYml)
		if err != nil {
			return nil, errors.New("Wrong directory. `gocomu.yml` file not found")
		}
	}

	yamlData := &GocomuYaml{}
	yaml.Unmarshal(data, &yamlData)

	return yamlData, nil
}
