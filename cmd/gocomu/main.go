package main

import (
	"fmt"

	gocomu "github.com/gocomu/cli"
	"github.com/gookit/color"
	"github.com/leaanthony/clir"
)

func main() {
	// Create new cli
	cli := clir.NewCli("GOCOMU", "GOCOMU CLI", "v0.0.13")

	gocomu.ClirActions(cli)

	// Run!
	err := cli.Run()
	if err != nil {
		color.Warn.Println(fmt.Sprint(err))
		return
	}
}
