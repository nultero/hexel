package hexel

import (
	"errors"
	"image"
	"image/color"
	_ "image/jpeg" // these have to be in here or else image.Decode doesn't register these formats properly
	"image/png"
	"os"
	"strings"

	"github.com/fogleman/gg"
)

func LoadImg(path string) ImageContext {
	f, err := os.Open(path)
	if err != nil {
		blame(LoadImg, err)
	}

	defer f.Close()

	img, _, err := image.Decode(f)

	if err != nil {
		blame(LoadImg, err)
	}

	pixels := getPixels(img)

	return ImageContext{
		Img:      img,
		FilePath: path,
		Grid:     Grid{pixels},
	}
}

func (ic *ImageContext) SaveImg() error {

	if len(ic.FilePath) == 0 {
		return errors.New("cannot save image context without filepath")

	} else if len(ic.Grid.Pixels) == 0 {
		return errors.New("cannot save image context; no content to save in pixel grid")
	}

	newImg := ic.imgFromGrid()

	f, err := os.Create(ic.FilePath)
	if err != nil {
		return err
	}
	defer f.Close()
	return png.Encode(f, newImg)
}

func (ic *ImageContext) imgFromGrid() image.Image {
	rect := image.Rectangle{}

	rect.Min = image.Point{0, 0}
	maxX := len(ic.Grid.Pixels[0])
	maxY := len(ic.Grid.Pixels)
	rect.Max = image.Point{maxX, maxY}

	newImg := image.NewRGBA(rect)

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			c := ic.Grid.Pixels[y][x].getColor()
			newImg.Set(x, y, c)
		}
	}

	return newImg
}

func (ic *ImageContext) ModPath(modification string) error {
	spl := strings.Split(ic.FilePath, ".")
	if len(spl) != 2 {
		return errors.New("wanted filename like `___.png`, got " + ic.FilePath)
	}

	name := spl[0]
	name += "_" + modification
	name += ".png"

	ic.FilePath = name
	return nil
}

func (ic *ImageContext) PlanetCircleTemplate(radius int, color color.Color) {
	ctx := gg.NewContextForImage(ic.imgFromGrid())
	ctx.SetColor(color)

	ctr := ic.Grid.Center()
	x, y := float64(ctr.X), float64(ctr.Y)
	ctx.DrawCircle(x, y, float64(radius))

	ctx.Fill()

	ic.Img = ctx.Image()
	ic.Grid = Grid{getPixels(ic.Img)}
}
