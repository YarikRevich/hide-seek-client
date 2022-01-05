package render

import (
	"errors"

	screenhistory "github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/hajimehoshi/ebiten/v2"
)

var instance IRender

type render struct {
	toRender                 []func(*ebiten.Image)
	permitToGetNumOfRendered bool
	numOfRendered            int
}

type IRender interface {
	SetToRender(func(*ebiten.Image))
	Render()

	CleanRenderPool()
}

func (r *render) SetToRender(c func(*ebiten.Image)) {
	r.numOfRendered = 0
	r.permitToGetNumOfRendered = false
	r.toRender = append(r.toRender, c)
}

func (r *render) Render() {
	screen := screenhistory.UseScreen().GetScreen()
	for _, v := range r.toRender {
		v(screen)
		r.numOfRendered++
	}
	r.permitToGetNumOfRendered = true
}

//Used for testing purpose
func (r *render) NumOfRenderer() (int, error) {
	if r.permitToGetNumOfRendered {
		return r.numOfRendered, nil
	}
	return 0, errors.New("NumOfRendered should be called after Render")
}

func (r *render) CleanRenderPool() {
	r.toRender = r.toRender[:0]
}

func UseRender() IRender {
	if instance == nil {
		instance = &render{
			toRender: make([]func(*ebiten.Image), 0, 2000),
		}
	}
	return instance
}
