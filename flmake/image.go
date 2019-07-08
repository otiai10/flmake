package flmake

import "io"

// Image ...
type Image string

// Open ...
func (img Image) Open() (io.ReadCloser, error) {
	return nil, nil
}
