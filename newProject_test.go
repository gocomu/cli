package cli

import (
	"testing"
	"time"
)

func TestProject(t *testing.T) {
	// dir, _ := os.Getwd()
	t.Run("Run newProject()", func(t *testing.T) {
		newProject(Gui, "test", PortAudio)
		time.Sleep(300 * time.Millisecond)
	})

	t.Run("Run newProject()", func(t *testing.T) {
		newProject(Cli, "test", PortAudio)
		time.Sleep(300 * time.Millisecond)
	})
	// os.RemoveAll(dir + "/test")
}
