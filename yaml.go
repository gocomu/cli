package cli

import (
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
func Yaml() *GocomuYaml {
	data, _ := ioutil.ReadFile("gocomu.yml")
	yamlData := &GocomuYaml{}
	yaml.Unmarshal(data, &yamlData)
	return yamlData
}
