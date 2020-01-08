package cli

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestBuildApp(t *testing.T) {
	// prepare gocomu.yml
	yamlData := &GocomuYaml{}
	data, _ := yaml.Marshal(yamlData)
	err := ioutil.WriteFile("gocomu.yml", data, 0755)
	if err != nil {
		t.Error()
	}

	t.Run("Run buildApp()", func(t *testing.T) {
		err := buildApp()
		if err != nil {
			fmt.Println(err)
		}
	})

	os.Remove("gocomu.yml")
}
