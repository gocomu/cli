package templates

// ServeGo holds the gocomu serve template
const ServeGo = `//+build serve

package main

import (
	"{{ .ProjectName }}/cmd"
)

func main() {
	cmd.GOCOMU()
}
`
