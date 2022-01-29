package camera

import (
	"fmt"
	"math"
	"time"

	"github.com/YarikRevich/hide-seek-client/internal/core/objects"
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
	"github.com/YarikRevich/hide-seek-client/internal/core/world"
	"github.com/hajimehoshi/ebiten/v2"
)

type AnimationOpts struct {
	Type int

	//General duration of animation
	Duration time.Duration

	//Delay between animation loop
	Delay time.Duration

	//Max rotation for object during animation
	MaxRotation float64

	//Rotation increases by gap
	RotationGap float64

	//Direction in which rotation will be committed
	RotationDirection int

	//Chan used for animation ending
	Cancel <-chan int
}

type AnimationExecutionOpts struct {
	opts AnimationOpts

	animationDuration <-chan time.Time

	animationDelay *time.Ticker

	stub *time.Ticker
}

const (
	ForwardRotationDirection = iota
	BackRotationDirection

	LavaZoneAnimation = iota
)

var instance *Camera

// Camera can look at positions, zoom and rotate.
type Camera struct {
	X, Y, Rot, Scale, MaxScale, MinScale float64

	//Means the objects camera follows
	followed *objects.Base
}

func (c *Camera) SetFollowed(b *objects.Base) {
	c.followed = b
}

func (c *Camera) GetFollowed() *objects.Base {
	return c.followed
}

//Defines camera shaking animation
func (c *Camera) lavaZoneAnimation(animationExecutionOpts AnimationExecutionOpts) {
	var endingIteration bool
	go func() {
		for {
			select {
			case <-animationExecutionOpts.animationDelay.C:
				if endingIteration && math.Floor(c.Rot*100)/100 == 0 {
					c.Rot = 0
					return
				}
				switch animationExecutionOpts.opts.RotationDirection {
				case ForwardRotationDirection:
					if c.Rot <= animationExecutionOpts.opts.MaxRotation {
						c.Rot += animationExecutionOpts.opts.RotationGap
					} else {
						c.Rot = animationExecutionOpts.opts.MaxRotation
						animationExecutionOpts.opts.RotationDirection = BackRotationDirection
					}
				case BackRotationDirection:
					if c.Rot >= -animationExecutionOpts.opts.MaxRotation {
						c.Rot -= animationExecutionOpts.opts.RotationGap
					} else {
						animationExecutionOpts.opts.RotationDirection = ForwardRotationDirection
					}
				}
			case <-animationExecutionOpts.animationDuration:
				endingIteration = true
			case <-animationExecutionOpts.opts.Cancel:
				endingIteration = true
			case <-animationExecutionOpts.stub.C:
			}
		}
	}()
}

//Animation execution pipeline
func (c *Camera) StartAnimation(opts AnimationOpts) {
	animationDuration := time.After(opts.Duration)
	animationDelay := time.NewTicker(opts.Delay)
	stub := time.NewTicker(time.Millisecond * 200)

	if opts.Delay <= 0 {
		animationDuration = nil
	}

	animationExecutionOpts := AnimationExecutionOpts{
		opts:              opts,
		animationDuration: animationDuration,
		animationDelay:    animationDelay,
		stub:              stub,
	}

	switch opts.Type {
	case LavaZoneAnimation:
		c.lavaZoneAnimation(animationExecutionOpts)
	}
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
	}
	return c
}

//Applies scale and further translation for draw
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

func (c *Camera) GetCameraTranslation() types.Vec2 {
	return c.getCameraTranslation(c.X, c.Y)
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
	return (wc * c.X) / -cPos.X
}

func (c *Camera) GetWorldCoordY(wc float64) float64 {
	cPos := c.GetCameraTranslation()
	return (wc * c.Y) / -cPos.Y
}

func (c *Camera) IsOuttaCoordX(x float64) bool {
	cPos := c.GetCameraTranslation()
	return -cPos.X > x
}

func (c *Camera) IsOuttaCoordY(y float64) bool {
	cPos := c.GetCameraTranslation()
	return -cPos.Y > y
}

func (c *Camera) IsLowerZeroCoordX() bool {
	cPos := c.GetCameraTranslation()
	return -cPos.X <= 0
}

func (c *Camera) IsLowerZeroCoordY() bool {
	cPos := c.GetCameraTranslation()
	sHUD := screen.UseScreen().GetHUDOffset()
	return -cPos.Y <= -sHUD
}

func UseCamera() *Camera {
	if instance == nil {
		instance = &Camera{
			Scale:    2.5,
			MaxScale: 3.2,
			MinScale: 2.3,
		}
	}
	return instance
}
