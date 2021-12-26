package game

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/animation"
	"github.com/YarikRevich/hide-seek-client/internal/core/world"
)

// "github.com/YarikRevich/hide-seek-client/internal/core/world"

func Exec() {
	pcs := world.UseWorld().GetPCs()
	for _, v := range pcs {
		animation.Animate(&v.Base)
	}
}
