package colormasked

func (c *Image) ToFloat() [][]float64 {
	bounds := c.Bounds()

	floats := make([][]float64, 0, bounds.Max.X-bounds.Min.X)
	for x := bounds.Min.X; x <= bounds.Max.X; x++ {
		agg := make([]float64, 0, bounds.Max.Y-bounds.Min.Y)
		for y := bounds.Min.Y; y <= bounds.Max.Y; y++ {
			v := float64(c.GrayValueAt(x, y)) / float64(math.MaxUint16)
			agg = append(agg, v)
		}
		floats = append(floats, agg)
	}
	
	return floats
}
