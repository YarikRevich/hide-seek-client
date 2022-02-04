package ui

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
)

type Label struct {
	text string
}

func (l *Label) Render(screen screen.Screen) {
	for _, c := range l.text {
		screen.RenderTextCharachter(c, types.Vec2{10, 10})
	}
}

func NewLabel(text string) Component {
	return &Label{text}
}
