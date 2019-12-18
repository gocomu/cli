package cli

import (
	"log"
	"os"
	"os/exec"

	"github.com/gocomu/cli/templates"
)

func Embed() error {
	embed()
	return nil
}

// arg: yamlData.Name
func embed() {
	for {
		// generate embed.go
		templates.CreateFile("cmd/embed/", "main.go", templates.EmbedGo, &templates.Data{})
		// run go run -tags embed ./cmd/embed
		// Start a process:
		cmd := exec.Command("go", "run", "-tags", "embed", "./cmd/embed")
		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}
		// delete embed.go
		os.Remove("cmd/embed/main.go")
	}
}
