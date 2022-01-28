// package camera

// import (
// 	"github.com/YarikRevich/hide-seek-client/internal/core/objects"
// 	// "github.com/YarikRevich/hide-seek-client/internal/core/screen"
// 	"github.com/hajimehoshi/ebiten/v2"
// )

// var instance *Camera

// type Hero struct {
// 	Followed       *objects.PC
// 	followedMatrix ebiten.GeoM
// }

// func (h *Hero) GetMatrixFor(p *objects.PC) ebiten.GeoM {
// 	if p.ID == h.Followed.ID {
// 		return h.followedMatrix
// 	}
// 	g := ebiten.GeoM{}

// 	// w := objects.UseObjects().World()
// 	// m := w.GetMetadata().Modified
// 	// g.Scale(p.GetZoomForSkin(m.Camera.Zoom))
// 	// // g.Translate(p.GetZoomedRawPos(w.GetZoomedMapScale()))

// 	return g
// }

// func (h *Hero) UpdateMatrix() {
// 	// w := objects.UseObjects().World()
// 	// wm := w.GetMetadata().Modified

// 	// h.followedMatrix.Scale(h.Followed.GetMovementRotation(), 1)
// 	// h.followedMatrix.Scale(h.Followed.GetZoomForSkin(wm.Camera.Zoom))

// 	// h.followedMatrix.Translate(h.Followed.GetZoomedRawPosForCamera(w.GetZoomedMapScale()))
// }

// type Map struct {
// 	matrix   ebiten.GeoM
// 	Followed *objects.PC
// }

// // if m := mc.Modified.Scale.X / 100 * c.Zoom; m != mc.RuntimeDefined.ZoomedScale.X {
// // 	mc.RuntimeDefined.ZoomedScale.X = m
// // }

// // if m := mc.Modified.Scale.Y / 100 * c.Zoom; m != mc.RuntimeDefined.ZoomedScale.Y {
// // 	mc.RuntimeDefined.ZoomedScale.Y = m
// // }

// func (m *Map) GetMatrix() ebiten.GeoM {
// 	return m.matrix
// }

// func (m *Map) UpdateMatrix() {
// 	// // fmt.Println(m.Followed.IsTranslationMovementBlocked())
// 	// if m.Followed.IsTranslationMovementBlocked() {
// 	// 	x, y := m.Followed.GetZoomedRawPos(w.GetZoomedMapScale())
// 	// 	ax, ay := m.Followed.GetZoomedRawPosForCamera(w.GetZoomedMapScale())

// 	// 	dy := y - ay
// 	// 	dx := x - ax
// 	// 	sy := y + ay
// 	// 	// sx := x + ax
// 	// 	if m.Followed.TranslationMovementXBlocked {
// 	// 		if m.IsOuttaAvailableAreaX(){
// 	// 			m.AttachMapX()

// 	// 		// if (dx) < 0 && m.Followed.IsDirectionLEFT() {
// 	// 		// 	m.Followed.SetTranslationXMovementBlocked(false)
// 	// 		// 	w.SetZoomedAttachedPosX(0)
// 	// 			// } else if (sx + pm.Size.Width/2) > wm.Size.Width*wsx && m.Followed.IsDirectionRIGHT() {
// 	// 			// 	m.Followed.SetTranslationXMovementBlocked(false)
// 	// 			// 	w.SetZoomedAttachedPosX(-((dx - pm.Size.Width/2) - (sx - wm.Size.Width*wsx)))
// 	// 		} else {
// 	// 			if dx < 0 {
// 	// 				m.matrix.Translate(0, 0)
// 	// 			} else {
// 	// 				m.matrix.Translate(-dx, 0)
// 	// 			}
// 	// 		}
// 	// 	}

// 	// 	if m.Followed.TranslationMovementYBlocked {
// 	// 		// fmt.Println(m.Followed.Direction)
// 	// 		// fmt.Println(sy+pm.Size.Height*2 > wm.Size.Height*wsy)
// 	// 		if (dy < 0) && m.Followed.IsDirectionUP() {
// 	// 			m.Followed.SetTranslationYMovementBlocked(false)
// 	// 			w.SetZoomedAttachedPosY(0)

// 	// 		} else if ((sy + pm.Size.Height*2) > wm.Size.Height*wsy) && m.Followed.IsDirectionDOWN() {
// 	// 			m.Followed.SetTranslationYMovementBlocked(false)
// 	// 			fmt.Println(((dy) - (sy - wm.Size.Height*wsy)))
// 	// 			w.SetZoomedAttachedPosY(-((dy - pm.Size.Height*2) - (sy - wm.Size.Height*wsy)))

// 	// 		} else {
// 	// 			// if dy < 0 {
// 	// 			// 	m.matrix.Translate(0, 0)
// 	// 			// }else if ((sy + pm.Size.Height/2) > wm.Size.Height*wsy){
// 	// 			// 	m.matrix.Translate(0, -((dy - pm.Size.Height / 2) - (sy - wm.Size.Height*wsy)))
// 	// 			// } else {
// 	// 			if dy < 0 {
// 	// 				m.matrix.Translate(0, 0)
// 	// 			} else {
// 	// 				m.matrix.Translate(0, -dy)
// 	// 			}
// 	// 		}
// 	// 	}
// 	// }

// 	// co := objects.UseObjects().Camera()

// 	// cZoomedX, cZoomedY := co.GetScaledPosX(), co.GetScaledPosY()
// 	// pZoomedOffsetX := m.Followed.GetScaledOffsetX()

// 	// m.matrix.Translate(-cZoomedX, -cZoomedY)

// 	// if m.Followed.TranslationMovementYBlocked {
// 	// 	if cZoomedY <= 0 && m.Followed.IsDirectionUP() {
// 	// 		m.Followed.SetTranslationYMovementBlocked(false)
// 	// 	}
// 	// 	if cZoomedY >= wm.Size.Height*zoomedMapScaleY && m.Followed.IsDirectionDOWN() {
// 	// 		m.Followed.SetTranslationYMovementBlocked(false)
// 	// 	}
// 	// }

// 	// if m.Followed.TranslationMovementXBlocked {

// 	// 	if cZoomedX + pZoomedOffsetX*2 >= wm.Size.Width*zoomedMapScaleX && m.Followed.IsDirectionRIGHT() {
// 	// 		// fmt.Println(cZoomedX + pZoomedOffsetX*2 - wm.Size.Width*zoomedMapScaleX, "HERE")
// 	// 		m.Followed.SetTranslationXMovementBlocked(false)
// 	// 	}
// 	// 	if cZoomedX <= 0 && m.Followed.IsDirectionLEFT() {
// 	// 		m.Followed.SetTranslationXMovementBlocked(false)
// 	// 	}
// 	// }

// 	// fmt.Println(cZoomedX, pZoomedOffsetX*2, wm.Size.Width*zoomedMapScaleX)

// 	// ax, ay := w.GetZoomedAttachedPos()
// 	// co := objects.UseObjects().Camera()
// 	// fmt.Println(wm.Size.Width, co.RawPos)

// 	// if !m.Followed.TranslationMovementXBlocked && !m.Followed.TranslationMovementYBlocked {
// 	// m.matrix.Translate(ax, ay)
// 	// }

// 	// if !m.Followed.TranslationMovementXBlocked && m.Followed.TranslationMovementYBlocked {
// 	// 	m.matrix.Translate(ax, 0)
// 	// }

// 	// if m.Followed.TranslationMovementXBlocked && !m.Followed.TranslationMovementYBlocked {
// 	// 	m.matrix.Translate(0, ay)
// 	// }
// }

// type Camera struct {
// 	Hero
// 	Map
// }

// func (c *Camera) Follow(p *objects.PC) {
// 	c.Hero.Followed = p
// 	c.Map.Followed = p
// }

// //Updates camera properties
// func (c *Camera) UpdateMatrices() {
// 	c.Map.matrix.Reset()
// 	c.Hero.followedMatrix.Reset()

// 	c.Hero.UpdateMatrix()
// 	c.Map.UpdateMatrix()
// }

// //Uses or creates a new instance of camera
// func UseCamera() *Camera {
// 	if instance == nil {
// 		instance = new(Camera)
// 	}
// 	return instance
// }

package camera

import (
	"fmt"
	"math"
	"time"

	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
	"github.com/YarikRevich/hide-seek-client/internal/core/world"
	"github.com/hajimehoshi/ebiten/v2"
)

// Camera can look at positions, zoom and rotate.
type Camera struct {
	X, Y, Rot, Scale float64
	Width, Height    int
}

// NewCamera returns a new Camera
func NewCamera(width, height int, x, y, rotation, zoom float64) *Camera {
	return &Camera{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
		Rot:    rotation,
		Scale:  zoom,
	}
}

func (c *Camera) RunShakingLimitedAnimation(duration, delay time.Duration, maxRot, rotGap float64, cancel <-chan int) {
	shakingDuration := time.After(duration)
	shakingTicker := time.NewTicker(delay)
	if delay <= 0 {
		shakingTicker.Stop()
	}
	stubTicker := time.NewTicker(time.Millisecond * 400)
	shakingDirection := 1
	endingIter := false
	go func() {
		for {
			select {
			case <-shakingTicker.C:
				if endingIter && math.Floor(c.Rot*100)/100 == 0 {
					c.Rot = 0
					return
				}
				switch shakingDirection {
				case 1:

					if c.Rot <= maxRot {
						c.Rot += rotGap
					} else {
						c.Rot = maxRot
						shakingDirection = 0
					}
				case 0:
					if c.Rot >= -maxRot {
						c.Rot -= rotGap
					} else {
						shakingDirection = 1
					}
				}
			case <-shakingDuration:
				endingIter = true
			case <-cancel:
				endingIter = true
			case <-stubTicker.C:
			}
		}
	}()
}

// // SetPosition looks at a position
// func (c *Camera) SetPosition(x, y float64) *Camera {
// 	c.X = x
// 	c.Y = y
// 	return c
// }

// MovePosition moves the Camera by x and y.
// Use SetPosition if you want to set the position
func (c *Camera) MovePosition(x, y float64) *Camera {
	c.X += x
	c.Y += y
	return c
}

func (c *Camera) SetZeroScreenCameraPositionX() {
	// prevCPos := c.GetCameraTranslation()
	// prevX := c.X
	// fmt.Println(prevX, prevCPos.X)
	c.SetPositionX(0)
	cPos := c.GetCameraTranslation()
	fmt.Println(cPos.X, "ZERO")
	c.SetPositionX(c.GetWorldCoordX(cPos.X))
	// fmt.Println(cPos, prevCPos, prevX)
	// fmt.Println((prevX * cPos.X) / prevCPos.X)
	// c.SetPositionX(c.X / cPos.X)
	// c.SetPositionX((prevX * -cPos.X) / -prevCPos.X)
}

func (c *Camera) SetPositionX(x float64) *Camera {
	c.X = x
	return c
}

func (c *Camera) SetPositionY(y float64) *Camera {
	c.Y = y
	return c
}

// Rotate rotates by phi
func (c *Camera) Rotate(phi float64) *Camera {
	c.Rot += phi
	return c
}

// SetRotation sets the rotation to rot
func (c *Camera) SetRotation(rot float64) *Camera {
	c.Rot = rot
	return c
}

// Zoom *= the current zoom
func (c *Camera) Zoom(mul float64) *Camera {
	previousScale := c.Scale
	c.Scale *= mul
	// s := screen.UseScreen().GetAxis()
	appliedCPos := c.GetCameraTranslation()
	appliedCPos.Y -= screen.UseScreen().GetHUDOffset()

	wM := world.UseWorld().GetWorldMap()
	wMSize := wM.GetSize()

	// fmt.Println(appliedX, appliedY, c.X, c.Y)
	// v2, _ := c.GetScreenCoordsTranslation(0, c.Y)
	// appliedY -= screen.UseScreen().GetHUDOffset()

	// appliedWorldX, appliedWorldY := c.GetWorldCoordsTranslation(c.X, c.Y)
	// appliedWorldY -= screen.UseScreen().GetHUDOffset()

	// fmt.Println(v1-c.X, vq-c.Y)
	// fmt.Println(appliedX, s.X/c.Scale)s
	if (!math.Signbit(appliedCPos.X) || !math.Signbit(appliedCPos.Y)) || ((wMSize.X*c.Scale < wMSize.X) || (wMSize.Y*c.Scale < wMSize.Y)) {
		fmt.Println(previousScale)
		c.Scale = previousScale
	}
	c.Resize(c.Width, c.Height)
	return c
}

// // SetZoom sets the zoom
// func (c *Camera) SetZoom(zoom float64) *Camera {
// 	c.Scale = zoom
// 	if c.Scale <= 0.01 {
// 		c.Scale = 0.01
// 	}
// 	c.Resize(c.Width, c.Height)
// 	return c
// }

// Resize resizes the camera Surface
func (c *Camera) Resize(w, h int) *Camera {
	c.Width = w
	c.Height = h
	return c
}

// // GetTranslation returns the coordinates based on the given x,y offset and the
// // camera's position
// func (c *Camera) GetTranslation(x, y float64) *ebiten.DrawImageOptions {
// 	op := &ebiten.DrawImageOptions{}
// 	op.GeoM.Translate(float64(c.Width)/2, float64(c.Height)/2)
// 	op.GeoM.Translate(-c.X+x, -c.Y+y)
// 	op.GeoM.Scale(c.Scale, c.Scale)
// 	return op
// }

// Blit draws the camera's surface to the screen and applies zoom
func (c *Camera) GetCameraOptions() *ebiten.DrawImageOptions {
	op := &ebiten.DrawImageOptions{}

	s := screen.UseScreen().GetAxis()

	op.GeoM.Rotate(c.Rot)
	op.GeoM.Translate(-s.X, -s.Y)
	op.GeoM.Translate(-c.X, -c.Y)

	op.GeoM.Translate((c.X - s.X), (c.Y - s.Y))
	op.GeoM.Translate(-s.X/2, -(s.Y/2 - screen.UseScreen().GetHUDOffset()/2))
	op.GeoM.Translate(-(c.X - s.X), -(c.Y - s.Y))
	op.GeoM.Scale(c.Scale, c.Scale)
	op.GeoM.Translate(s.X, s.Y)
	op.GeoM.Translate(s.X*c.Scale, s.Y*c.Scale)

	return op
}

func (c *Camera) GetCameraTranslationXM(x float64) float64 {
	return c.getCameraTranslation(x, 0).X
}

func (c *Camera) GetCameraTranslationYM(y float64) float64 {
	return c.getCameraTranslation(y, 0).Y
}

func (c *Camera) GetCameraTranslation() types.Vec2 {
	return c.getCameraTranslation(c.X, c.Y)
}

func (c *Camera) getCameraTranslation(x, y float64) types.Vec2 {
	var rX, rY float64

	s := screen.UseScreen().GetAxis()

	rX += -s.X
	rY += -s.Y

	rX += -x
	rY += -y

	rX += x - s.X
	rY += y - s.Y

	rX += -s.X / 2
	rY += -(s.Y/2 - screen.UseScreen().GetHUDOffset()/2)

	rX += -(x - s.X)
	rY += -(y - s.Y)

	rX *= c.Scale
	rY *= c.Scale

	rX += s.X
	rY += s.Y

	rX += s.X * c.Scale
	rY += s.Y * c.Scale

	return types.Vec2{X: rX, Y: rY}
}

func (c *Camera) GetScreenCoordsTranslation(x, y float64) types.Vec2 {
	var rX, rY float64

	s := screen.UseScreen().GetAxis()

	rX += -s.X
	rY += -s.Y

	rX += x
	rY += y

	rX += -(x - s.X)
	rY += -(y - s.Y)

	rX += -s.X / 2
	rY += -(s.Y/2 - screen.UseScreen().GetHUDOffset()/2)

	rX += (x - s.X)
	rY += (y - s.Y)

	rX *= c.Scale
	rY *= c.Scale

	rX += s.X
	rY += s.Y

	rX += s.X * c.Scale
	rY += s.Y * c.Scale

	return types.Vec2{X: rX, Y: rY}
}

func (c *Camera) GetWorldCoordX(wc float64) float64 {
	cPos := c.GetCameraTranslation()
	fmt.Println((wc*c.X)/-cPos.X, "GET WORLD COORD X")
	return (wc * c.X) / -cPos.X
}
func (c *Camera) GetWorldCoordY(wc float64) float64 {
	cPos := c.GetCameraTranslation()
	return (wc * c.Y) / -cPos.Y
}

// // GetCursorCoords converts cursor/screen coords into world coords
// func (c *Camera) GetCursorCoords() (float64, float64) {
// 	cx, cy := ebiten.CursorPosition()
// 	return c.GetWorldCoords(float64(cx), float64(cy))
// }

var Cam = NewCamera(1113, 670, 0, 0, 0, 2)
