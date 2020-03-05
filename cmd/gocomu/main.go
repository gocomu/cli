package main

import (
	"fmt"

	gocomu "github.com/gocomu/cli"
	// release "github.com/gocomu/version"
	"github.com/leaanthony/clir"
)

func main() {
	// Create new cli
	cli := clir.NewCli("GOCOMU", "GOCOMU CLI", gocomu.Version)

	gocomu.ClirActions(cli)

	// Run!
	err := cli.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
