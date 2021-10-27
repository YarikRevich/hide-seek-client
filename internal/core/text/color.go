package text

import (
	"image/color"

	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/models"
)

type colorSession struct{
	color models.FontColor
}

type IColorSession interface {
	GetColor() color.RGBA
}

//Compares color stated in metadata
//to available colors and returns rgba 
//representing if color is correct else
//it returns transparent one
func (c *colorSession) GetColor()color.RGBA{
	switch c.color{
	case models.Black:
		return color.RGBA{0, 0, 0, 255}
	case models.White:
		return color.RGBA{255, 255, 255, 255}
	}
	return color.RGBA{}
}

func NewColorSession(color models.FontColor)IColorSession{
	return &colorSession{color: color}
}