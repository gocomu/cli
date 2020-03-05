package cli

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/google/go-github/v29/github"
)

// fetch latest release from github
func fetchLatestRelease() (string, error) {
	client := github.NewClient(nil)
	release, _, err := client.Repositories.GetLatestRelease(context.Background(), "gocomu", "cli")
	if err != nil {
		return "", err
	}

	return release.GetTagName(), nil
}

// updateGocomu checks against github's latest release
// and if there is a newer tag updates the binary
func updateGocomu() error {
	// get current version
	version, err := exec.Command("gocomu", "version").Output()
	if err != nil {
		return err
	}

	// fetch github latest release tag
	latestVersion, err := fetchLatestRelease()
	if err != nil {
		return err
	}

	// if versions are equal stop the process
	if strings.TrimSuffix(string(version), "\n") == string(latestVersion) {
		return errors.New("Already up to date")
	}

	// else move on with the update
	fmt.Println("Updating")
	// for the update, very conveniently, we use `go get`
	cmd := exec.Command("go", "get", "-u", "-ldflags", "'-X main.version="+string(latestVersion)+"'", "github.com/gocomu/cli/cmd/gocomu")
	cmd.Env = os.Environ()
	// we need to run the command with GO111MODULE env off
	cmd.Env = append(cmd.Env, "GO111MODULE=off")
	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Println("New version")
	// print/check new version
	newVersion, err := exec.Command("gocomu", "version").Output()
	if err != nil {
		return err
	}
	fmt.Println(strings.TrimSuffix(string(newVersion), "\n"))

	fmt.Println("Done!")
	return nil
}
