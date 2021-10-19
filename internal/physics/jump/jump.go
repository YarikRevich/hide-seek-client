package jump

import (
	"math"
	"time"

	"github.com/YarikRevich/HideSeek-Client/internal/direction"
	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/pc"
	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/world"
)

var JumpGap = time.NewTicker(time.Millisecond * 20)

//Returns the trace for jump due to the elementHeight
//and current world G metric
func CalculateJump(p *pc.PC) {
	if len(p.Physics.Jump) == 0 {
		t := int(math.Round((math.Sqrt((2*p.Metadata.Size.Height)/world.UseWorld().Location.Metadata.Physics.G) / 2)))

		for t%2 != 0 {
			t++
		}

		for i := 0; i < int(t)*3; i++ {
			p.Physics.Jump = append(p.Physics.Jump, direction.UP)
		}

		for i := 0; i < int(t)*3; i++ {
			p.Physics.Jump = append(p.Physics.Jump, direction.DOWN)
		}
	}
}

//Process jump for the object
func ProcessJump(p *pc.PC){
	if len(p.Physics.Jump) != 0 {
		select {
		case <-JumpGap.C:
			j := p.Physics.Jump[0]

			if j == direction.UP {
				p.Y -= 2
			}

			if j == direction.DOWN {
				p.Y += 2
			}

			p.Physics.Jump = p.Physics.Jump[1:]
		default:
		}
	}
}
