package main

import (
	"fmt"

	gocomu "github.com/gocomu/cli"
	"github.com/gookit/color"
	"github.com/leaanthony/clir"
)

func main() {
	// Create new cli
	cli := clir.NewCli("GOCOMU", "GOCOMU CLI", "v0.0.1")

	gocomu.ClirActions(cli)

	// Run!
	err := cli.Run()
	if err != nil {
		color.Warn.Printf(fmt.Sprint(err))
		return
	}
}
