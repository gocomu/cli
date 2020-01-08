package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("updating")
	// dir, _ := os.UserHomeDir()
	// os.Chdir(dir)
	// os.Mkdir(".gocomutmp", 0755)
	// os.Chdir(".gocomutpm")
	cmd := exec.Command("go", "get", "-u", "github.com/gocomu/cli")
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "GO111MODULE=off")
	if err := cmd.Run(); err != nil {
		fmt.Println("Aborted")
		return
	}
	// os.Chdir(dir)
	// os.RemoveAll(".gocomutmp")
	fmt.Println("Done!")
}
