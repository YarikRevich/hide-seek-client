package camera

import (
	"fmt"

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

	fmt.Println(h.Followed.IsTranslationMovementBlocked())

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

	wsx, wsy := w.GetZoomedMapScale()
	m.matrix.Scale(wsx, wsy)

	if m.Followed.IsTranslationMovementBlocked() {
		x, y := m.Followed.GetZoomedRawPos(w.GetZoomedMapScale())
		ax, ay := m.Followed.GetZoomedRawPosForCamera(w.GetZoomedMapScale())

		fmt.Println(ax, ay)
		if m.Followed.TranslationMovementXBlocked {
			m.matrix.Translate(-(x - ax), 0)

			if (x - ax) < 0 {
				m.Followed.SetTranslationXMovementBlocked(false)
				w.SetZoomedAttachedPosX(0)
			}

			if (x + ax) > wm.Size.Width*wsx {
				m.Followed.SetTranslationXMovementBlocked(false)
				w.SetZoomedAttachedPosY(-((x - ax) - (x + ax - wm.Size.Width*wsx)))
			}
		}

		if m.Followed.TranslationMovementYBlocked {
			m.matrix.Translate(0, -(y - ay))

			// fmt.Println((y - ay < 0), (y + ay) > wm.Size.Height*wsy)
			fmt.Println(y - ay)
			if (y - ay) < 0 {
				m.Followed.SetTranslationYMovementBlocked(false)
				w.SetZoomedAttachedPosY(0)
			}

			if (y + ay) > wm.Size.Height*wsy {
				m.Followed.SetTranslationYMovementBlocked(false)
				// fmt.Println("HERE", y, ay, (y - ay), (y + ay - wm.Size.Height*wsy))

				w.SetZoomedAttachedPosY(-((y - ay) - (y + ay - wm.Size.Height*wsy)))
				// m.Followed.SetZoomedRawPosY(ay)
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

	c.Map.UpdateMatrix()
	c.Hero.UpdateMatrix()
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
