package physics

import (
	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/pc"
	"github.com/YarikRevich/HideSeek-Client/internal/physics/jump"
)

func ProcessAnimation(p *pc.PC){
	jump.ProcessJump(p)
}