package cli

import (
	"os"
)

// current working directory
var dir string

func init() {
	// assign global var dir to current working directory
	dir, _ = os.Getwd()
}
