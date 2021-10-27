package physics

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	"github.com/YarikRevich/HideSeek-Client/internal/physics/jump"
)

func ProcessAnimation(o *objects.Object){
	jump.ProcessJump(o)
}