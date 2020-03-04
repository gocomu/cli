package cli

import (
	"fmt"
	"os"
)

// current working directory
var dir string

func init() {
	dir, _ = os.Getwd()

	_, err := Yaml()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
