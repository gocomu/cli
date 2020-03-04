package cli

import (
	"fmt"
	"os/exec"
)

func updateGocomu() error {
	fmt.Println("Current version")
	version, err := exec.Command("gocomu", "version").Output()
	if err != nil {
		return err
	}
	fmt.Println(string(version))

	fmt.Println("Updating")
	// cmd := exec.Command("go", "get", "-u", "github.com/gocomu/cli/cmd/gocomu")
	// cmd.Env = os.Environ()
	// cmd.Env = append(cmd.Env, "GO111MODULE=off")
	// if err := cmd.Run(); err != nil {
	// 	return err
	// }

	fmt.Println("New version")
	newVersion, err := exec.Command("gocomu", "version").Output()
	if err != nil {
		return err
	}
	fmt.Println(string(newVersion))

	fmt.Println("Done!")
	return nil
}
