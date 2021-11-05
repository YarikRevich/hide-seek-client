package render

import (
	screenhistory "github.com/YarikRevich/HideSeek-Client/internal/core/screen"
	"github.com/hajimehoshi/ebiten/v2"
)

var instance IRender

type render struct {
	renderList     []func(*ebiten.Image)
}

type IRender interface {
	SetToRender(func(*ebiten.Image))
	Render()

	CleanRenderPool()
}

func (r *render) SetToRender(c func(*ebiten.Image)) {
	r.renderList = append(r.renderList, c)
}

func (r *render) Render() {
	screen := screenhistory.GetScreen()
	for _, v := range r.renderList {
		v(screen)
	}
}

func (r *render) CleanRenderPool() {
	r.renderList = r.renderList[:0]
}

func UseRender() IRender {
	if instance == nil {
		instance = &render{
			renderList:     make([]func(*ebiten.Image), 0, 100),
		}
	}
	return instance
}
