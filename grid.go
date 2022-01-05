package hexel

import "image/color"

type Grid struct {
	Pixels [][]Pixel
}

func (gd *Grid) Bounds() (int, int) {
	x := len(gd.Pixels[0])
	y := len(gd.Pixels)
	return x, y
}

func (gd *Grid) Center() Point {
	x := len(gd.Pixels[0]) / 2
	y := len(gd.Pixels) / 2
	return Point{uint32(x), uint32(y)}
}

// Completely whites out grid canvas.
func (gd *Grid) TabulaRasa() {
	w := ColorToPx(color.White)
	for y, pxRow := range gd.Pixels {
		for x := range pxRow {
			gd.Pixels[y][x] = w
		}
	}
}
