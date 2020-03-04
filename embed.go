package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gocomu/cli/templates"
)

// Embed is used to embded all wav/aiff files under 'embed' dir
func embedAudio() error {
	fmt.Println("Embedding")
	fmt.Println("This might take a while, depending on the size of audio files")
	// generate embed.go
	templates.CreateFile(filepath.Join(dir, "cmd"), "gocomuEmbed.go", templates.EmbedGo, &templates.Data{})
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	// run go run -tags embed ./cmd/embed.go
	cmd, err := exec.Command("go", "run", "-tags", "embed", filepath.Join(dir, "cmd", "gocomuEmbed.go")).Output()
	if err != nil {
		return err
	}

	fmt.Println(string(cmd))
	// delete embed.go
	os.Remove(filepath.Join(dir, "cmd", "gocomuEmbed.go"))
	fmt.Println("Done!")
	return nil
}
