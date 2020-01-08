package cli

import (
	"fmt"
	"os"
	"os/exec"
)

func updateGocomu() error {
	fmt.Println("updating")
	cmd := exec.Command("go", "get", "-u", "github.com/gocomu/cli/cmd/gocomu")
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "GO111MODULE=off")
	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Println("Done!")
	return nil
}
