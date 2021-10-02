package jump

import (
	"github.com/YarikRevich/HideSeek-Client/internal/direction"
	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/pc"
	"github.com/YarikRevich/HideSeek-Client/internal/world"
	"math"
)

func CalculateJump(){
	p := pc.GetPC()
	w := world.UseWorld()

	t := (math.Sqrt((2 * p.Metadata.Size.Height)/w.Metadata.Physics.G)/2)
	for i := 0; i < int(t); i++{
		p.Physics.Jump = append(p.Physics.Jump, direction.UP)
	}

	for i := 0; i < int(t); i++{
		p.Physics.Jump = append(p.Physics.Jump, direction.DOWN)
	}
}