package camera

import (
	// "fmt"

	"fmt"
	// "math"

	// "github.com/YarikRevich/HideSeek-Client/internal/core/keycodes"
	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	"github.com/hajimehoshi/ebiten/v2"
)

var instance *Camera

type Hero struct {
	Followed       *objects.PC
	followedMatrix ebiten.GeoM
}

func (h *Hero) GetMatrixFor(p *objects.PC) ebiten.GeoM {
	if p.ID == h.Followed.ID {
		return h.followedMatrix
	}
	g := ebiten.GeoM{}

	w := objects.UseObjects().World()
	m := w.GetMetadata().Modified
	g.Scale(p.GetZoomForSkin(m.Camera.Zoom))
	g.Translate(p.GetZoomedRawPos(w.GetZoomedMapScale()))

	return g
}

func (h *Hero) UpdateMatrix() {
	w := objects.UseObjects().World()
	wm := w.GetMetadata().Modified

	h.Followed.GetMovementRotation()
	h.followedMatrix.Scale(h.Followed.GetMovementRotation(), 1)
	h.followedMatrix.Scale(h.Followed.GetZoomForSkin(wm.Camera.Zoom))

	// fmt.Println(h.Followed.IsTranslationMovementBlocked())

	if !h.Followed.TranslationMovementXBlocked && w.IsAxisXCrossedBy(h.Followed) {
		h.Followed.SetTranslationXMovementBlocked(true)
	}

	if !h.Followed.TranslationMovementYBlocked && w.IsAxisYCrossedBy(h.Followed) {
		h.Followed.SetTranslationYMovementBlocked(true)
	}

	h.followedMatrix.Translate(h.Followed.GetZoomedRawPosForCamera(w.GetZoomedMapScale()))
}

type Map struct {
	matrix   ebiten.GeoM
	Followed *objects.PC
}

func (m *Map) GetMatrix() ebiten.GeoM {
	return m.matrix
}

func (m *Map) UpdateMatrix() {
	w := objects.UseObjects().World()
	wm := w.GetMetadata().Modified

	pm := m.Followed.GetMetadata().Origin

	wsx, wsy := w.GetZoomedMapScale()
	m.matrix.Scale(wsx, wsy)

	fmt.Println(m.Followed.IsTranslationMovementBlocked())
	if m.Followed.IsTranslationMovementBlocked() {
		x, y := m.Followed.GetZoomedRawPos(w.GetZoomedMapScale())
		ax, ay := m.Followed.GetZoomedRawPosForCamera(w.GetZoomedMapScale())

		dy := y - ay
		dx := x - ax
		sy := y + ay
		sx := x + ax
		if m.Followed.TranslationMovementXBlocked {
			if (dx) < 0 && m.Followed.IsDirectionLEFT() {
				m.Followed.SetTranslationXMovementBlocked(false)
				w.SetZoomedAttachedPosX(0)
			} else if (sx + pm.Size.Width/2) > wm.Size.Width*wsx && m.Followed.IsDirectionRIGHT() {
				m.Followed.SetTranslationXMovementBlocked(false)
				w.SetZoomedAttachedPosX(-((dx - pm.Size.Width/2) - (sx - wm.Size.Width*wsx)))
			} else {
				if dx < 0 {
					m.matrix.Translate(0, 0)
				} else {
					m.matrix.Translate(-dx, 0)
				}
			}
		}

		if m.Followed.TranslationMovementYBlocked {
			fmt.Println(m.Followed.Direction)
			if (dy < 0) && m.Followed.IsDirectionUP() {
				m.Followed.SetTranslationYMovementBlocked(false)
				w.SetZoomedAttachedPosY(0)

				
			} else if ((sy + pm.Size.Height/2) > wm.Size.Height*wsy) && m.Followed.IsDirectionDOWN() {
				m.Followed.SetTranslationYMovementBlocked(false)
				w.SetZoomedAttachedPosY(-((dy - pm.Size.Height/2) - (sy - wm.Size.Height*wsy)))

			} else {
				// if dy < 0 {
				// 	m.matrix.Translate(0, 0)
				// }else if ((sy + pm.Size.Height/2) > wm.Size.Height*wsy){
				// 	m.matrix.Translate(0, -((dy - pm.Size.Height / 2) - (sy - wm.Size.Height*wsy)))
				// } else {
				if dy < 0 {
					m.matrix.Translate(0, 0)
				} else {
					m.matrix.Translate(0, -dy)
				}
			}
		}
	}

	ax, ay := w.GetZoomedAttachedPos()

	if !m.Followed.TranslationMovementXBlocked && !m.Followed.TranslationMovementYBlocked {
		m.matrix.Translate(ax, ay)
	}

	if !m.Followed.TranslationMovementXBlocked && m.Followed.TranslationMovementYBlocked {
		m.matrix.Translate(ax, 0)
	}

	if m.Followed.TranslationMovementXBlocked && !m.Followed.TranslationMovementYBlocked {
		m.matrix.Translate(0, ay)
	}
}

type Camera struct {
	Hero
	Map
}

func (c *Camera) Follow(p *objects.PC) {
	c.Hero.Followed = p
	c.Map.Followed = p
}

//Updates camera properties
func (c *Camera) UpdateMatrices() {
	c.Map.matrix.Reset()
	c.Hero.followedMatrix.Reset()

	c.Hero.UpdateMatrix()
	c.Map.UpdateMatrix()
}

//Increments zoom property
func (c *Camera) ZoomIn() {
	w := objects.UseObjects().World()
	m := w.GetMetadata().Modified
	if m.Camera.Zoom < m.Camera.MaxZoom {
		m.Camera.Zoom++
	}
}

//Decrements zoom property
func (c *Camera) ZoomOut() {
	w := objects.UseObjects().World()
	m := w.GetMetadata().Modified
	if m.Camera.Zoom > m.Camera.MinZoom {
		m.Camera.Zoom--
	}
}

//Uses or creates a new instance of camera
func UseCamera() *Camera {
	if instance == nil {
		instance = new(Camera)
	}
	return instance
}
