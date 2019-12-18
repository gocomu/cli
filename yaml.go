package cli

import (
	"fmt"
	"io/ioutil"

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
func Yaml() (*GocomuYaml, error) {
	data, err := ioutil.ReadFile("gocomu.yml")
	if err != nil {
		return nil, fmt.Errorf("Are you in the correct directory? \n Error: %v", err)
	}

	yamlData := &GocomuYaml{}
	yaml.Unmarshal(data, &yamlData)
	return yamlData, nil
}
