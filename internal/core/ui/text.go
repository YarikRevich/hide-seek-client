package ui

import (
	"image/color"

	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
)

type TextOpts struct {
	sources.Align

	Text     string
	RowWidth float64
	Position types.Vec2
	Font     *sources.Font
	Color    color.Color
}
