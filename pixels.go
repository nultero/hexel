package hexel

import (
	"image"
	"image/color"
)

type Point struct {
	X, Y uint32
}

type Pixel struct {
	R, G, B, A uint8
}

func toPix(r, g, b, a uint32) Pixel {
	return Pixel{
		R: uint8(r / 257),
		G: uint8(g / 257),
		B: uint8(b / 257),
		A: uint8(a / 257),
	}
}

func (p *Pixel) getColor() color.RGBA {
	return color.RGBA{
		R: uint8(p.R),
		G: uint8(p.G),
		B: uint8(p.B),
		A: uint8(p.A),
	}
}

func getPixels(im image.Image) [][]Pixel {

	pix := [][]Pixel{}

	b := im.Bounds()
	x, y := b.Max.X, b.Max.Y

	for j := 0; j < y; j++ {
		px := []Pixel{}
		for i := 0; i < x; i++ {
			px = append(px, toPix(im.At(i, j).RGBA()))
		}

		pix = append(pix, px)
	}

	return pix
}
