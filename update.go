package cli

import (
	"fmt"
	"os/exec"
)

func updateGocomu() error {
	fmt.Println("updating")
	// cmd := exec.Command("go", "get", "-u", "github.com/gocomu/cli")
	// cmd.Env = os.Environ()
	// cmd.Env = append(cmd.Env, "GO111MODULE=off")
	// if err := cmd.Run(); err != nil {
	// 	return err
	// }

	// dir, _ := os.UserHomeDir()
	// os.Chdir(dir)
	// os.Mkdir(".gocomutmp", 0755)
	// os.Chdir(".gocomutpm")
	cmd := exec.Command("gocomuUpdate")
	// cmd.Env = os.Environ()
	// cmd.Env = append(cmd.Env, "GO111MODULE=off")
	if err := cmd.Run(); err != nil {
		return err
	}
	// os.Chdir(dir)
	// os.RemoveAll(".gocomutmp")
	fmt.Println("Done!")
	return nil
}
