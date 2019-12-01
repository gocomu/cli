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

// Yaml .
func Yaml() *GocomuYaml {
	// create gocomu.yml
	data, _ := ioutil.ReadFile("/gocomu.yml")
	yamlData := &GocomuYaml{}
	yaml.Unmarshal(data, &yamlData)
	return yamlData
}
