package camera

import (
	"fmt"

	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	// "github.com/YarikRevich/HideSeek-Client/internal/core/screen"
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
	// g.Translate(p.GetZoomedRawPos(w.GetZoomedMapScale()))

	return g
}

func (h *Hero) UpdateMatrix() {
	w := objects.UseObjects().World()
	wm := w.GetMetadata().Modified

	h.followedMatrix.Scale(h.Followed.GetMovementRotation(), 1)
	h.followedMatrix.Scale(h.Followed.GetZoomForSkin(wm.Camera.Zoom))

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

	zoomedMapScaleX, zoomedMapScaleY := w.GetZoomedMapScale()
	m.matrix.Scale(zoomedMapScaleX, zoomedMapScaleY)

	// // fmt.Println(m.Followed.IsTranslationMovementBlocked())
	// if m.Followed.IsTranslationMovementBlocked() {
	// 	x, y := m.Followed.GetZoomedRawPos(w.GetZoomedMapScale())
	// 	ax, ay := m.Followed.GetZoomedRawPosForCamera(w.GetZoomedMapScale())

	// 	dy := y - ay
	// 	dx := x - ax
	// 	sy := y + ay
	// 	// sx := x + ax
	// 	if m.Followed.TranslationMovementXBlocked {
	// 		if m.IsOuttaAvailableAreaX(){
	// 			m.AttachMapX()

	// 		// if (dx) < 0 && m.Followed.IsDirectionLEFT() {
	// 		// 	m.Followed.SetTranslationXMovementBlocked(false)
	// 		// 	w.SetZoomedAttachedPosX(0)
	// 			// } else if (sx + pm.Size.Width/2) > wm.Size.Width*wsx && m.Followed.IsDirectionRIGHT() {
	// 			// 	m.Followed.SetTranslationXMovementBlocked(false)
	// 			// 	w.SetZoomedAttachedPosX(-((dx - pm.Size.Width/2) - (sx - wm.Size.Width*wsx)))
	// 		} else {
	// 			if dx < 0 {
	// 				m.matrix.Translate(0, 0)
	// 			} else {
	// 				m.matrix.Translate(-dx, 0)
	// 			}
	// 		}
	// 	}

	// 	if m.Followed.TranslationMovementYBlocked {
	// 		// fmt.Println(m.Followed.Direction)
	// 		// fmt.Println(sy+pm.Size.Height*2 > wm.Size.Height*wsy)
	// 		if (dy < 0) && m.Followed.IsDirectionUP() {
	// 			m.Followed.SetTranslationYMovementBlocked(false)
	// 			w.SetZoomedAttachedPosY(0)

	// 		} else if ((sy + pm.Size.Height*2) > wm.Size.Height*wsy) && m.Followed.IsDirectionDOWN() {
	// 			m.Followed.SetTranslationYMovementBlocked(false)
	// 			fmt.Println(((dy) - (sy - wm.Size.Height*wsy)))
	// 			w.SetZoomedAttachedPosY(-((dy - pm.Size.Height*2) - (sy - wm.Size.Height*wsy)))

	// 		} else {
	// 			// if dy < 0 {
	// 			// 	m.matrix.Translate(0, 0)
	// 			// }else if ((sy + pm.Size.Height/2) > wm.Size.Height*wsy){
	// 			// 	m.matrix.Translate(0, -((dy - pm.Size.Height / 2) - (sy - wm.Size.Height*wsy)))
	// 			// } else {
	// 			if dy < 0 {
	// 				m.matrix.Translate(0, 0)
	// 			} else {
	// 				m.matrix.Translate(0, -dy)
	// 			}
	// 		}
	// 	}
	// }

	co := objects.UseObjects().Camera()
	cZoomedX, cZoomedY := co.GetZoomedRawPos(zoomedMapScaleX, zoomedMapScaleY)
	pZoomedOffsetX, _ := m.Followed.GetZoomedRawPosForCamera(zoomedMapScaleX, zoomedMapScaleY)

	if m.Followed.TranslationMovementYBlocked {
		if cZoomedY <= 0 && m.Followed.IsDirectionUP() {
			m.Followed.SetTranslationYMovementBlocked(false)
		}
		if cZoomedY >= wm.Size.Height*zoomedMapScaleY && m.Followed.IsDirectionDOWN() {
			m.Followed.SetTranslationYMovementBlocked(false)
		}
	}

	if m.Followed.TranslationMovementXBlocked {
		if co.GetRawX()+pZoomedOffsetX >= wm.Size.Width*zoomedMapScaleX && m.Followed.IsDirectionRIGHT() {
			m.Followed.SetTranslationXMovementBlocked(false)
		}
		if cZoomedX <= 0 && m.Followed.IsDirectionLEFT() {
			m.Followed.SetTranslationXMovementBlocked(false)
		}
	}

	fmt.Println(co.GetRawX(), m.Followed.RawPosForCamera.X, wm.Size.Width*zoomedMapScaleX)

	m.matrix.Translate(-co.RawPos.X, -co.RawPos.Y)
	// ax, ay := w.GetZoomedAttachedPos()
	// co := objects.UseObjects().Camera()
	// fmt.Println(wm.Size.Width, co.RawPos)

	// if !m.Followed.TranslationMovementXBlocked && !m.Followed.TranslationMovementYBlocked {
	// m.matrix.Translate(ax, ay)
	// }

	// if !m.Followed.TranslationMovementXBlocked && m.Followed.TranslationMovementYBlocked {
	// 	m.matrix.Translate(ax, 0)
	// }

	// if m.Followed.TranslationMovementXBlocked && !m.Followed.TranslationMovementYBlocked {
	// 	m.matrix.Translate(0, ay)
	// }
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

	c.Hero.Followed.SetTranslationXMovementBlocked(false)
	c.Hero.Followed.SetTranslationYMovementBlocked(false)
}

//Decrements zoom property
func (c *Camera) ZoomOut() {
	w := objects.UseObjects().World()
	m := w.GetMetadata().Modified

	// wsx, _ := w.GetZoomedMapScale()
	// czx, _ := c.Hero.Followed.GetZoomedRawPosForCamera(w.GetZoomedMapScale())
	// zx, _ := c.Hero.Followed.GetZoomedRawPos(w.GetZoomedMapScale())
	// fmt.Println(zx, czx, m.Size.Width * wsx)
	// co := objects.UseObjects().Camera()
	// fmt.Println(co.RawPos.X, czx)
	co := objects.UseObjects().Camera()
	fmt.Println(co.RawPos.X+1 < 1113)
	// if co.RawPos.X+1 < 1113 {
	if m.Camera.Zoom > m.Camera.MinZoom {
		m.Camera.Zoom--
	}
	// }

	c.Hero.Followed.SetTranslationXMovementBlocked(false)
	c.Hero.Followed.SetTranslationYMovementBlocked(false)
}

//Uses or creates a new instance of camera
func UseCamera() *Camera {
	if instance == nil {
		instance = new(Camera)
	}
	return instance
}
