package cli

import (
	"fmt"
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestYaml(t *testing.T) {
	t.Run("Run Yaml() with error", func(t *testing.T) {
		_, err := Yaml()
		if err != nil {
			fmt.Println(err)
		}
	})

	t.Run("Run Yaml()", func(t *testing.T) {
		//dir, _ := os.Getwd()
		// prepare gocomu.yml
		yamlData := &GocomuYaml{}
		data, _ := yaml.Marshal(yamlData)
		f, _ := os.Create("gocomu.yml")
		f.Write(data)
		defer f.Close()
		// err := ioutil.WriteFile("gocomu.yml", data, 0755)
		// if err != nil {
		// fmt.Println(err)
		// t.Error()
		// }

		_, err := Yaml()
		if err != nil {
			//t.Error()
			fmt.Println(err)
		}

		os.Remove("gocomu.yml")
	})

}
