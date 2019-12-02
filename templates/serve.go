package templates

const ServeGo = `//+build serve

package main

import (
	"{{ .ProjectName }}/cmd"
)

func main() {
	cmd.GOCOMU()
}
`
