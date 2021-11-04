package physics

import (
	"math"
	"time"

	"github.com/YarikRevich/HideSeek-Client/internal/core/keycodes"
	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
)

type Jump struct {
	ticker *time.Ticker
}

//Calculates jump init values and
//runs further planned execution
func (j *Jump) Calculate() {
	p := objects.UseObjects().PC()

	if len(p.Physics.Jump) != 0 {
		return
	}

	pm := objects.UseObjects().PC().GetMetadata()
	wm := objects.UseObjects().World().GetMetadata()
	t := int(math.Round((math.Sqrt((2*pm.Modified.Size.Height)/wm.Modified.Physics.G) / 2)))

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
	p := objects.UseObjects().PC()

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
