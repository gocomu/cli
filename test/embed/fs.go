package embed

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"sync"
)

// storage holds a map (key/value) which corresponds
// to filepath (string) "./embed/filename" and file data (string)
// when app inits all embedded audio files are loaded here
var storage = make(map[string]string)
var mutex = &sync.Mutex{}

// List returns all audio files (wav/aiff)
// stored in binary via 'gocomu embed'
// for example you can use embed.List() to range over available
// stored files and subsequently use embed.Open to open them
func List() []string {
	var list []string
	for v := range storage {
		list = append(list, v)
	}
	return list
}

// Open a file (*bytes.Reader) stored in binary.
// It is safe to be used with goroutines.
// Conveniently is can be used along os.Open("path/to/file")
// and when you are ready to build a standalone app
// replace os.Open with embed.Open
func Open(filepath string) (*bytes.Reader, error) {
	var val string
	var ok bool

	mutex.Lock()
	// check if filepath exists in storage
	if val, ok = storage[filepath]; !ok {
		// if not a gracefull error returns
		return nil, fmt.Errorf("File '%s' doesn't exist/wrong filepath", filepath)
	}
	mutex.Unlock()

	// decode from base64 to []byte
	data, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		return nil, err
	}

	// create a new reader and load the data ([]byte)
	datareader := bytes.NewReader(data)

	// ungzip the data ([]byte)
	gzipReader, err := gzip.NewReader(datareader)
	if err != nil {
		return nil, err
	}
	defer gzipReader.Close()

	// load the reader to a new []byte
	bt, _ := ioutil.ReadAll(gzipReader)

	// re-assign a new reader to the decoded & unzgziped data ([]byte)
	datareader = bytes.NewReader(bt)
	return datareader, nil
}

// add is the function embedded audio files call when app is init'ing
func add(filepath string, data string) {
	mutex.Lock()
	storage[filepath] = data
	mutex.Unlock()
}
