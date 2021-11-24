package game

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/animation"
	"github.com/YarikRevich/HideSeek-Client/internal/core/world"
)

func Exec() {
	for _, v := range world.UseWorld().GetPCs() {
		animation.Animate(&v.Base)
	}
}
