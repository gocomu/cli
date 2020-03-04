package cli

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

// BuildApp .
func buildApp() error {
	fmt.Println("Building")
	yamlData, _ := Yaml()
	cmd := exec.Command("go", "build", "-o", filepath.Join(dir, "output", yamlData.Name), filepath.Join(dir, "cmd", yamlData.Name))
	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Printf(`Binary %s built successfully, check "output/" directory
`, yamlData.Name)

	return nil
}
