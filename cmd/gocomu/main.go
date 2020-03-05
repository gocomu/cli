package main

import (
	"fmt"

	gocomu "github.com/gocomu/cli"
	"github.com/leaanthony/clir"
)

// Version .
var Version string

func main() {
	// Create new cli
	cli := clir.NewCli("GOCOMU", "GOCOMU CLI", Version)

	gocomu.ClirActions(cli)

	// Run!
	err := cli.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
