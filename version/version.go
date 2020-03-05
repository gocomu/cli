// +build

package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/google/go-github/v29/github"
)

func main() {
	client := github.NewClient(nil)
	release, _, err := client.Repositories.GetLatestRelease(context.Background(), "gocomu", "cli")
	if err != nil {
		fmt.Println(err)
		return
	}

	data := []byte(`package cli

const Version = "` + release.GetTagName() + `" 
`)
	_ = ioutil.WriteFile("version.go", data, 0644)
}
