package templates

const CliTemplateMainGo = `package main

import (
	demo "{{ .ProjectName }}"

	gocomu "github.com/gocomu/cli"
	comu "github.com/gocomu/comu/cio"
	"github.com/leaanthony/clir"
)

func main() {
	// load project info details from gocomu.yml located at proejct root
	yamlData := gocomu.Yaml()

	clir := clir.NewCli(yamlData.Name, yamlData.Description, yamlData.Version)

	play := clir.NewSubCommand("play", "Start making sound")
	play.Action(func() error {
		// initialize an output
		// NewAudioIO() arguments are
		// 1: cio.RTout (real-time output)
		// option 1, cio.PortAduio
		// option 2, cio.Oto (wip)
		// 2: number of channels, 1-n
		// 3: buffer size, if you listen click try using a bigger buffer
		// keep in mind bigger buffer size equals to more latency
		comuIO := comu.NewAudioIO(comu.PortAudio, 2, comu.BS2048)

		demo.Sine(comuIO)
		return nil
	})

	clir.Run()
}
`

const CliTemplateSineGo = `package {{ .ProjectName }}

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-audio/audio"
	"github.com/go-audio/generator"
	"github.com/gocomu/comu"
	"github.com/gocomu/comu/cio"
)

// Sine is a 440.0 sine wave that runs for 8 bars and then stops
func Sine(comuIO *cio.AudioIO) {
	// prepare a buffer
	buf := &audio.FloatBuffer{
		Data:   make([]float64, cio.BS2048),
		Format: audio.FormatStereo44100,
	}

	// start a sine wave generator
	osc := generator.NewOsc(generator.WaveSine, 440.0, buf.Format.SampleRate)
	osc.Amplitude = 0.5

	// start a new clock
	tempo := comu.NewClock(120.0)

	for {
		// populate the out buffer
		if err := osc.Fill(buf); err != nil {
			log.Printf("error filling up the buffer")
		}

		// pass populated buffer to PortAudio's stream
		comuIO.PortAudioOut(buf)

		switch {
		// exit after 8 bars
		case tempo.BarCounter == 8:
			fmt.Println("Exiting Demo at bar nu. ", tempo.BarCounter)
			fmt.Println("Time elapsed ", time.Now().Sub(tempo.TimeStarted))
			os.Exit(0)
		default:
		}
	}
}
`
