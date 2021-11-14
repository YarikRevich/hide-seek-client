package game

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/animation"
	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
)

func Exec(){
	w := objects.UseObjects().World()
	
	for _, v := range w.PCs{
		animation.Animate(&v.Object)
	}
}