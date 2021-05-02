package astar

import (
	"image/color"
)

var (
	BackgroundColor = color.RGBA{0xfa, 0xf8, 0xef, 0xff}
	FrameColor      = color.RGBA{0xbb, 0xad, 0xa0, 0xff}
	White           = color.RGBA{0x00, 0x00, 0x00, 0x00}
	Black           = color.RGBA{0x13, 0x13, 0x13, 0xff}
	Green           = color.RGBA{0x1e, 0x82, 0x4c, 0xff}
	Red             = color.RGBA{0x96, 0x28, 0x1b, 0xff}
	Orange          = color.RGBA{0xff, 0xa5, 0x00, 0xff}
	LightBlue       = color.RGBA{0xad, 0xd8, 0xe6, 0xff}
)

func colorToScale(clr color.Color) (float64, float64, float64, float64) {
	r, g, b, a := clr.RGBA()
	rf := float64(r) / 0xffff
	gf := float64(g) / 0xffff
	bf := float64(b) / 0xffff
	af := float64(a) / 0xffff
	// Convert to non-premultiplied alpha components.
	if 0 < af {
		rf /= af
		gf /= af
		bf /= af
	}
	return rf, gf, bf, af
}
