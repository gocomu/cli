package main

import (
	gocomu "github.com/gocomu/cli"
	"github.com/leaanthony/clir"
)

func main() {
	// Create new cli
	cli := clir.NewCli("GOCOMU", "GOCOMU CLI", "v0.0.1")

	gocomu.ClirActions(cli)

	// Run!
	cli.Run()
}
