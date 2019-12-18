package templates

// SineGo holds the hello world template
const SineGo = `package {{ .ProjectName }}

import (
	"fmt"
	"log"
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
			return
		default:
		}
	}
}
`
