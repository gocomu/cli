package cli

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/gocomu/cli/templates"
)

func Embed() error {
	fmt.Println("Embedding started")
	// generate embed.go
	templates.CreateFile("cmd/", "gocomuEmbed.go", templates.EmbedGo, &templates.Data{})
	dir, _ := os.Getwd()
	// run go run -tags embed ./cmd/embed.go
	cmd, _ := exec.Command("go", "run", "-tags", "embed", dir+"/cmd/gocomuEmbed.go").Output()
	fmt.Println(string(cmd))
	// delete embed.go
	os.Remove(dir + "/cmd/gocomuEmbed.go")
	fmt.Println("Finished")
	return nil
}
