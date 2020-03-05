package main

import (
	"fmt"

	gocomu "github.com/gocomu/cli"
	"github.com/leaanthony/clir"
)

// Version .
var version = "v0.0.1"

func main() {
	// Create new cli
	cli := clir.NewCli("GOCOMU", "GOCOMU CLI", version)

	gocomu.ClirActions(cli)

	// Run!
	err := cli.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
