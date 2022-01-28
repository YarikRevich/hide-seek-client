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

// //Updates camera properties
// func (c *Camera) UpdateMatrices() {
// 	c.Map.matrix.Reset()
// 	c.Hero.followedMatrix.Reset()

// 	c.Hero.UpdateMatrix()
// 	c.Map.UpdateMatrix()
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

type Animation struct{}

func (a *Animation) StartAnimation(duration, delay time.Duration, maxRot, rotGap float64, cancel <-chan int) {

}

// Camera can look at positions, zoom and rotate.
type Camera struct {
	X, Y, Rot, Scale, MaxScale, MinScale float64
	Width, Height                        int
	Animation
}

// NewCamera returns a new Camera
func NewCamera(width, height int, x, y, rotation, zoom float64) *Camera {
	return &Camera{
		X:        x,
		Y:        y,
		Width:    width,
		Height:   height,
		Rot:      rotation,
		Scale:    2.5,
		MaxScale: 3.0,
		MinScale: 2.4,
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

// MovePosition moves the Camera by x and y.
// Use SetPosition if you want to set the position
func (c *Camera) MovePosition(x, y float64) *Camera {
	c.X += x
	c.Y += y
	return c
}

func (c *Camera) SetPositionX(x float64) *Camera {
	c.X = x
	return c
}

func (c *Camera) SetPositionY(y float64) *Camera {
	c.Y = y
	return c
}

func (c *Camera) SetZeroPositionX() {
	c.SetPositionX(c.getCameraTranslationX(0) / c.Scale)
}

func (c *Camera) SetZeroPositionY() {
	sHUD := screen.UseScreen().GetHUDOffset()
	c.SetPositionY(c.getCameraTranslationY(0)/c.Scale - sHUD/c.Scale)
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
	fmt.Println(c.MinScale < mul, mul < c.MaxScale, c.MinScale, mul, c.MaxScale)
	if c.MinScale < c.Scale*mul && c.Scale*mul < c.MaxScale {

		previousScale := c.Scale
		c.Scale *= mul
		appliedCPos := c.GetCameraTranslation()
		appliedCPos.Y -= screen.UseScreen().GetHUDOffset()

		wM := world.UseWorld().GetWorldMap()
		wMSize := wM.GetSize()

		sAxis := screen.UseScreen().GetAxis()
		if (!math.Signbit(appliedCPos.X) || !math.Signbit(appliedCPos.Y)) ||
			((wMSize.X*c.Scale < wMSize.X) || (wMSize.Y*c.Scale < wMSize.Y)) ||
			((math.Abs(appliedCPos.X) > wMSize.X*c.Scale-sAxis.X*2) || math.Abs(appliedCPos.Y) > wMSize.Y*c.Scale-sAxis.Y*2) {
			c.Scale = previousScale
		}
		fmt.Println(c.Scale)
		c.Resize(c.Width, c.Height)
	}
	return c
}

// Resize resizes the camera Surface
func (c *Camera) Resize(w, h int) *Camera {
	c.Width = w
	c.Height = h
	return c
}

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

func (c *Camera) getCameraTranslationX(x float64) float64 {
	return c.getCameraTranslation(x, 0).X
}

func (c *Camera) getCameraTranslationY(y float64) float64 {
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

var Cam = NewCamera(1113, 670, 0, 0, 0, 2)
