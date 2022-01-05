package hexel

import "image"

type ImageContext struct {
	Img      image.Image
	FilePath string
	Grid     Grid
}

func ImgCtxFromBounds(x, y int) ImageContext {

	pxls := [][]Pixel{}
	for j := 0; j < y; j++ {
		rw := []Pixel{}
		for i := 0; i < x; i++ {
			rw = append(rw, Pixel{0, 0, 0, 255})
		}

		pxls = append(pxls, rw)
	}

	return ImageContext{Grid: Grid{pxls}}
}
