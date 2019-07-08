package flmake

import (
	"fmt"
	"image"
	"image/draw"

	"github.com/golang/freetype"
)

// Builder ...
type Builder struct {
	Config *Config
}

// Build ...
func (b *Builder) Build() error {
	dest, err := b.Config.Images[0].Open()
	if err != nil {
		return err
	}
	defer dest.Close()
	img, _, err := image.Decode(dest)
	if err != nil {
		return err
	}
	drawable, ok := img.(draw.Image)
	if !ok {
		return fmt.Errorf("failed to convert image to a drawable")
	}
	ctx := freetype.NewContext()
	ctx.SetDst(drawable)
	return nil
}
