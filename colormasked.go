package colormasked

import (
	"image"
	"image/color"
)

// From Hjulle on http://stackoverflow.com/a/17076395/199475

type Image struct {
	Img image.Image
	Mod color.Model
}

// We return the new color model...
func (c *Image) ColorModel() color.Model {
	return c.Mod
}

// ... but the original bounds
func (c *Image) Bounds() image.Rectangle {
	return c.Img.Bounds()
}

// At forwards the call to the original image and
// then asks the color model to convert it.
func (c *Image) At(x, y int) color.Color {
	return c.Mod.Convert(c.Img.At(x, y))
}

// Extracts the numerical 16-bit grayscale value of the point
func (c *Image) GrayValueAt(x, y int) uint16 {
	return color.Gray16Model.Convert(c.Img.At(x, y)).(color.Gray16).Y
}
