package cli

import (
	"fmt"
	"os"
	"os/exec"
)

func updateGocomu() error {
	fmt.Println("updating")
	dir, _ := os.Getwd()
	os.Mkdir(dir+"/.gocomutmp", 0755)
	cmd := exec.Command("go", "get", "-u", "github.com/gocomu/cli")
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "GO111MODULE=off")
	if err := cmd.Run(); err != nil {
		return err
	}
	os.RemoveAll(dir + "/.gocomutmp")
	fmt.Println("Done!")
	return nil
}
