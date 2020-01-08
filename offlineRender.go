package cli

// import (
// 	"io"

// 	"github.com/go-audio/aiff"
// 	"github.com/go-audio/audio"
// 	"github.com/go-audio/wav"
// )

// // Render .
// func Render(dur, format int, buf audio.Buffer, w io.WriteSeeker) error {
// 	switch format {
// 	case 0:
// 		e := wav.NewEncoder(w,
// 			buf.PCMFormat().SampleRate,
// 			16,
// 			buf.PCMFormat().NumChannels, 1)
// 		if err := e.Write(buf.AsIntBuffer()); err != nil {
// 			return err
// 		}
// 		return e.Close()
// 	case 1:
// 		e := aiff.NewEncoder(w,
// 			buf.PCMFormat().SampleRate,
// 			16,
// 			buf.PCMFormat().NumChannels)
// 		if err := e.Write(buf.AsIntBuffer()); err != nil {
// 			return err
// 		}
// 		return e.Close()
// 		// default:
// 		// 	return errors.New("unknown format")
// 	}

// 	return nil
// }
