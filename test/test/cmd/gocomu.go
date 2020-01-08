package cmd

import (
	demo "test"

	"github.com/gocomu/comu/cio"
)

// GOCOMU .
func GOCOMU() {
	// initialize an output
	// NewAudioIO() arguments are
	// 1: cio.RTout (real-time output)
	// option 1, cio.PortAduio
	// option 2, cio.Oto (wip)
	// 2: number of channels, 1-n
	// 3: buffer size, if you listen click try using a bigger buffer
	// keep in mind bigger buffer size equals to more latency
	comuIO := cio.NewAudioIO(cio.PortAudio, 2, cio.BS2048)

	demo.Sine(comuIO)
}
