package cli

import (
	"fmt"
	"os/exec"
)

func updateGocomu() error {
	fmt.Println("updating")
	// run go mod tidy
	cmd := exec.Command("go", "get", "-u", "github.com/gocomu/cli")
	err := cmd.Run()
	if err != nil {
		return err
	}
	fmt.Println("Done!")
	return nil
}
