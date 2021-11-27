package game

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/animation"
	"github.com/YarikRevich/HideSeek-Client/internal/core/world"
)

// "github.com/YarikRevich/HideSeek-Client/internal/core/world"

func Exec() {
	pcs := world.UseWorld().GetPCs()
	for _, v := range pcs {
		animation.Animate(&v.Base)
	}
}
