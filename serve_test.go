package cli

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"gopkg.in/yaml.v3"
)

func TestServe(t *testing.T) {
	// prepare gocomu.yml
	dir, _ := os.Getwd()
	yamlData := &GocomuYaml{
		Name:        "test",
		Description: "t",
		Version:     "0.0",
		Type:        "cli",
		ServeOutput: "0",
	}
	data, _ := yaml.Marshal(yamlData)
	ioutil.WriteFile("gocomu.yml", data, 0755)
	os.Chdir(dir + "/test")
	t.Run("Run proejctServe()", func(t *testing.T) {
		go projectServe()
		time.Sleep(1 * time.Second)
	})

	os.Chdir(dir)
	os.RemoveAll(dir + "/test")

	t.Run("Run reload()", func(t *testing.T) {
		reload("test")
	})
}
