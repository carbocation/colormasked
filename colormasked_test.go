package colormasked

import (
	"image"
	"image/color"
	"image/draw"
	"testing"
)

var (
	white, black color.Color = color.RGBA{255, 255, 255, 255}, color.RGBA{0, 0, 0, 255}
)

func newImg() draw.Image {
	img := image.NewRGBA(image.Rect(0, 0, 40, 40))
	bounds := img.Bounds()

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			if y%2 == 0 {
				img.Set(x, y, white)
				continue
			}

			img.Set(x, y, black)
		}
	}

	return img
}

func TestToFloat(t *testing.T) {
	img := newImg()

	cm := Image{Img: img, Mod: color.RGBAModel}

	floats := cm.ToFloat()
	for _, row := range floats {
		for _, v := range row {
			_ = v
		}
	}
}

func TestAt(t *testing.T) {
	cm := Image{Img: newImg(), Mod: color.RGBAModel}

	if cm.At(0, 0) != white {
		t.Fatalf("Expected %v, got %v", white, cm.At(0, 0))
	}
}

func TestMod(t *testing.T) {
	cm := Image{Img: newImg(), Mod: color.Gray16Model}

	if cm.ColorModel() != color.Gray16Model {
		t.Fatalf("Expected %v, got %v", color.Gray16Model, cm.ColorModel())
	}
}
