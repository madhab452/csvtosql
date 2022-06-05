package reader

import (
	"encoding/csv"
	"os"
)

type chunk [][]string

// Reader represents reader.
type Reader struct {
	f  *os.File
	cr *csv.Reader
}

// NewReader returns an instance of reader
func NewReader(f *os.File) *Reader {
	return &Reader{
		f:  f,
		cr: csv.NewReader(f),
	}
}

// ReadChunks return header, chunks, and error
func (r *Reader) ReadChunks(size int) ([]string, []chunk, error) {
	var chunks []chunk
	var chunk [][]string
	counter := 1

	header, err := r.cr.Read()
	if err != nil {
		return []string{}, nil, nil
	}

	for {
		row, err := r.cr.Read()
		if row == nil {
			break
		}
		if err != nil {
			return nil, nil, err
		}
		if counter%size == 0 {
			chunks = append(chunks, chunk)
			chunk = [][]string{}
		}
		chunk = append(chunk, row)
		counter++
	}
	chunks = append(chunks, chunk) // last element
	return header, chunks, nil
}
