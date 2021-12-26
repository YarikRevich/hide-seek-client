package physics

import (
	"math"
	"time"

	"github.com/YarikRevich/hide-seek-client/internal/core/keycodes"
	"github.com/YarikRevich/hide-seek-client/internal/core/world"
)

type Jump struct {
	ticker *time.Ticker
}

//Calculates jump init values and
//runs further planned execution
func (j *Jump) Calculate() {
	p := world.UseWorld().GetPC()
	p.SaveAnimationStartPosition()

	if len(p.Physics.Jump) != 0 {
		return
	}

	t := int(math.Round((math.Sqrt((2*p.ModelCombination.Modified.Size.Height)/p.Parent.ModelCombination.Modified.Physics.G) / 2)))

	for t%2 != 0 {
		t++
	}

	for i := 0; i < int(t)*3; i++ {
		p.Physics.Jump = append(p.Physics.Jump, keycodes.UP)
	}

	for i := 0; i < int(t)*3; i++ {
		p.Physics.Jump = append(p.Physics.Jump, keycodes.DOWN)
	}

	j.ticker = time.NewTicker(time.Millisecond * 20)
	go j.execute()
}

func (j *Jump) execute() {
	p := world.UseWorld().GetPC()

	if len(p.Physics.Jump) == 0 {
		return
	}

	select {
	case <-j.ticker.C:
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

func NewJump() *Jump {
	return new(Jump)
}
