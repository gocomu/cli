package cli

import (
	"fmt"
	"testing"

	"github.com/leaanthony/clir"
)

func TestCli(t *testing.T) {
	cli := clir.NewCli("GOCOMU", "GOCOMU CLI", "v0.0.3")

	t.Run("cli tests", func(t *testing.T) {
		ClirActions(cli)
	})

	t.Run("cli customBanner()", func(t *testing.T) {
		dat := customBanner(cli)
		fmt.Println(dat)
	})

	t.Run("cli selectedOutputHelper", func(t *testing.T) {
		dat := selectedOutputHelper(0)
		if dat == PortAudio {
			fmt.Println("succ")
		}
	})

	t.Run("cli selectedOutputHelper", func(t *testing.T) {
		dat := selectedOutputHelper(1)
		if dat == Oto {
			fmt.Println("succ")
		}
	})

	t.Run("cli selectedOutputHelper with error", func(t *testing.T) {
		dat := selectedOutputHelper(0)
		if dat == Oto {
			fmt.Println("succ")
		}
	})
}
