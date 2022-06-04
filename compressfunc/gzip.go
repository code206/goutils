package compressfunc

import (
	"bytes"
	"compress/gzip"
)

func GZipBytes(data []byte) ([]byte, error) {
	var input bytes.Buffer
	g, err := gzip.NewWriterLevel(&input, 9)
	if err != nil {
		return nil, err
	}
	g.Write(data)
	g.Flush()
	g.Close()
	return input.Bytes(), nil
}
