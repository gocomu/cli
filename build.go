package cli

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// BuildApp .
func BuildApp() error {
	fmt.Println("Starting Building")
	binary, lookErr := exec.LookPath("go")
	if lookErr != nil {
		panic(lookErr)
	}

	yamlData, _ := Yaml()
	args := []string{"go", "build", "-o", "output/" + yamlData.Name, "./cmd/" + yamlData.Name}
	env := os.Environ()
	err := syscall.Exec(binary, args, env)
	if err != nil {
		return err
	}

	fmt.Printf(`
Sucess! File %s can be found inside output/ directory
`, yamlData.Name)

	return nil
}
