package cli

import (
	"fmt"
	"os/exec"
)

// BuildApp .
func BuildApp() error {
	fmt.Println("Started Building")

	yamlData, _ := Yaml()
	cmd := exec.Command("go", "build", "-o", "output/"+yamlData.Name, "./cmd/"+yamlData.Name)
	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Printf(`Sucess! File %s can be found inside "output/" directory
`, yamlData.Name)

	return nil
}
