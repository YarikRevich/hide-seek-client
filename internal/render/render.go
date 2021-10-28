package render

import (
	screenhistory "github.com/YarikRevich/HideSeek-Client/internal/core/screen"
	"github.com/YarikRevich/HideSeek-Client/internal/render/middlewares"
	"github.com/hajimehoshi/ebiten/v2"
)

var instance IRender

type callback func(*ebiten.Image)

type render struct {
	screen *ebiten.Image

	renderList     []callback
	postRenderList []callback
}

type IRender interface {
	SetToRender(callback)
	SetToPostRender(callback)

	rawRender([]callback)
	Render()
	PostRender()

	CleanRenderPool()
}

func (r *render) UpdateScreen(screen *ebiten.Image) {
	r.screen = screen
}

func (r *render) SetToRender(c callback) {
	r.renderList = append(r.renderList, c)
}

func (r *render) SetToPostRender(c callback) {
	r.postRenderList = append(r.postRenderList, c)
}

func (r *render) rawRender(lc []callback) {
	screen := screenhistory.GetScreen()
	for _, v := range lc {
		v(screen)
	}
}

func (r *render) Render() {
	r.rawRender(r.renderList)
}

func (r *render) PostRender() {
	r.rawRender(r.postRenderList)

	middlewares.UseRenderMiddlewares()
}

func (r *render) CleanRenderPool() {
	r.renderList = r.renderList[:0]
	r.postRenderList = r.postRenderList[:0]
}

func UseRender() IRender {
	if instance == nil {
		instance = &render{
			renderList:     make([]callback, 0, 100),
			postRenderList: make([]callback, 0, 100),
		}
	}
	return instance
}
