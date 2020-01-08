package cli

import (
	"testing"
)

func TestEmbed(t *testing.T) {
	t.Run("Run embed()", func(t *testing.T) {
		err := embedAudio()
		if err != nil {
			t.Error(err)
		}
	})
}
