package jump

import (
	"math"
	"time"

	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	"github.com/YarikRevich/HideSeek-Client/internal/core/keycodes"
)

var JumpGap = time.NewTicker(time.Millisecond * 20)

//Returns the trace for jump due to the elementHeight
//and current world G metric
func CalculateJump(p *objects.Object) {
	if len(p.Physics.Jump) == 0 {
		t := int(math.Round((math.Sqrt((2*p.Metadata.Size.Height)/objects.UseObjects().World().Skin.Metadata.Physics.G) / 2)))

		for t%2 != 0 {
			t++
		}

		for i := 0; i < int(t)*3; i++ {
			p.Physics.Jump = append(p.Physics.Jump, keycodes.UP)
		}

		for i := 0; i < int(t)*3; i++ {
			p.Physics.Jump = append(p.Physics.Jump, keycodes.DOWN)
		}
	}
}

//Process jump for the object
func ProcessJump(p *objects.Object){
	if len(p.Physics.Jump) != 0 {
		select {
		case <-JumpGap.C:
			j := p.Physics.Jump[0]

			if j == keycodes.UP {
				p.RawPos.Y -= 2
			}

			if j == keycodes.DOWN {
				p.RawPos.Y += 2
			}

			p.Physics.Jump = p.Physics.Jump[1:]
		default:
		}
	}
}
