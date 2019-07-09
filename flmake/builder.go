package flmake

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

// Builder ...
type Builder struct {
	Config *Config
}

// Build ...
func (b *Builder) Build() error {

	// Prepare base image
	src, err := b.Config.Images[0].Open()
	if err != nil {
		return err
	}
	defer src.Close()
	img, _, err := image.Decode(src)
	if err != nil {
		return err
	}

	canvas := image.NewRGBA(img.Bounds())
	draw.Draw(canvas, img.Bounds(), img, image.Point{}, 0)

	// Prepare drawing labels
	ctx := freetype.NewContext()
	ctx.SetDst(canvas)

	fontfile, err := os.Open("./Raleway-Thin.ttf")
	if err != nil {
		return err
	}
	defer fontfile.Close()
	ttf, err := ioutil.ReadAll(fontfile)
	if err != nil {
		return err
	}
	font, err := truetype.Parse(ttf)
	if err != nil {
		return err
	}
	ctx.SetDPI(72)
	ctx.SetFont(font)
	ctx.SetFontSize(32)
	ctx.SetClip(canvas.Bounds())
	ctx.SetDst(canvas)
	ctx.SetSrc(image.NewUniform(color.RGBA{255, 255, 255, 255}))
	point := freetype.Pt(100, 180+int(ctx.PointToFixed(16)>>6))
	// point := freetype.Pt(8, 8)

	if _, err := ctx.DrawString("Flyer Auto Generation Test", point); err != nil {
		return err
	}

	out, err := os.Create("./out.png")
	if err != nil {
		return err
	}
	if err := png.Encode(out, canvas); err != nil {
		return err
	}
	return nil
}
