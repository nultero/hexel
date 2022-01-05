package hexel

import (
	"fmt"
	"image/color"
)

func ColorToPx(c color.Color) Pixel {
	r, g, b, a := c.RGBA()
	return Pixel{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(a),
	}
}

type Palette struct{}

func (p *Palette) Black() {
	fmt.Println("")
}
