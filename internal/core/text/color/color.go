package color

import (
	"image/color"

	"github.com/YarikRevich/HideSeek-Client/internal/core/sources"
)

type Color struct{}

//Returns RGBA representation of
//color saved in metadata
func (c *Color) GetColor(fc sources.FontColor) color.RGBA {
	switch fc {
	case sources.Black:
		return color.RGBA{0, 0, 0, 255}
	case sources.White:
		return color.RGBA{255, 255, 255, 255}
	}
	return color.RGBA{}
}

func NewColor() *Color {
	return new(Color)
}
