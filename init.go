package cli

import (
	"os"
)

// current working directory
var dir string

func init() {
	// assign global var dir to current working directory
	dir, _ = os.Getwd()

	// check if dir is a GOCOMU project top folder
	// _, err := Yaml()
	// if not print error and exit
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
}
