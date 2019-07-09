package flmake

import (
	"io"
	"os"
)

// Image ...
type Image string

// Open ...
func (img Image) Open() (io.ReadCloser, error) {
	return os.Open(string(img))
}
