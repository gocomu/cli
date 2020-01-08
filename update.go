package cli

import (
	"fmt"
	"os"
	"os/exec"
)

func updateGocomu() error {
	fmt.Println("updating")
	os.Setenv("GO111MODULE", "off")
	cmd := exec.Command("go", "get", "-u", "github.com/gocomu/cli")
	err := cmd.Run()
	if err != nil {
		return err
	}
	fmt.Println("Done!")
	return nil
}
