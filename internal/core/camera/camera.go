package camera

// import (
// 	"fmt"
// 	"math"
// 	"time"

// 	"github.com/YarikRevich/hide-seek-client/internal/core/objects"
// 	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
// 	"github.com/YarikRevich/hide-seek-client/internal/core/types"
// 	"github.com/YarikRevich/hide-seek-client/internal/core/world"
// 	"github.com/hajimehoshi/ebiten/v2"
// )

// type AnimationOpts struct {
// 	Type int

// 	//General duration of animation
// 	Duration time.Duration

// 	//Delay between animation loop
// 	Delay time.Duration

// 	//Max rotation for object during animation
// 	MaxRotation float64

// 	//Rotation increases by gap
// 	RotationGap float64

// 	//Direction in which rotation will be committed
// 	RotationDirection int

// 	//Chan used for animation ending
// 	Cancel <-chan int
// }

// type AnimationExecutionOpts struct {
// 	opts AnimationOpts

// 	animationDuration <-chan time.Time

// 	animationDelay *time.Ticker

// 	stub *time.Ticker
// }

// const (
// 	ForwardRotationDirection = iota
// 	BackRotationDirection

// 	LavaZoneAnimation = iota
// )

// var instance *Camera

// // Camera can look at positions, zoom and rotate.
// type Camera struct {
// 	Rot                       float64
// 	Scale, MaxScale, MinScale types.Vec2

// 	//Means the objects camera follows
// 	Followed *objects.Base
// }

// func (c *Camera) SetFollowed(b *objects.Base) {
// 	c.Followed = b
// }

// //Defines camera shaking animation
// func (c *Camera) lavaZoneAnimation(animationExecutionOpts AnimationExecutionOpts) {
// 	var endingIteration bool
// 	go func() {
// 		for {
// 			select {
// 			case <-animationExecutionOpts.animationDelay.C:
// 				if endingIteration && math.Floor(c.Rot*100)/100 == 0 {
// 					c.Rot = 0
// 					return
// 				}
// 				switch animationExecutionOpts.opts.RotationDirection {
// 				case ForwardRotationDirection:
// 					if c.Rot <= animationExecutionOpts.opts.MaxRotation {
// 						c.Rot += animationExecutionOpts.opts.RotationGap
// 					} else {
// 						c.Rot = animationExecutionOpts.opts.MaxRotation
// 						animationExecutionOpts.opts.RotationDirection = BackRotationDirection
// 					}
// 				case BackRotationDirection:
// 					if c.Rot >= -animationExecutionOpts.opts.MaxRotation {
// 						c.Rot -= animationExecutionOpts.opts.RotationGap
// 					} else {
// 						animationExecutionOpts.opts.RotationDirection = ForwardRotationDirection
// 					}
// 				}
// 			case <-animationExecutionOpts.animationDuration:
// 				endingIteration = true
// 			case <-animationExecutionOpts.opts.Cancel:
// 				endingIteration = true
// 			case <-animationExecutionOpts.stub.C:
// 			}
// 		}
// 	}()
// }

// //Animation execution pipeline
// func (c *Camera) StartAnimation(opts AnimationOpts) {
// 	animationDuration := time.After(opts.Duration)
// 	animationDelay := time.NewTicker(opts.Delay)
// 	stub := time.NewTicker(time.Millisecond * 200)

// 	if opts.Delay <= 0 {
// 		animationDuration = nil
// 	}

// 	animationExecutionOpts := AnimationExecutionOpts{
// 		opts:              opts,
// 		animationDuration: animationDuration,
// 		animationDelay:    animationDelay,
// 		stub:              stub,
// 	}

// 	switch opts.Type {
// 	case LavaZoneAnimation:
// 		c.lavaZoneAnimation(animationExecutionOpts)
// 	}
// }

// // // MovePosition moves the Camera by x and y.
// // // Use SetPosition if you want to set the position
// // func (c *Camera) MovePosition(x, y float64) *Camera {
// // 	// cPos.X += x
// // 	// cPos.Y += y
// // 	return c
// // }

// // func (c *Camera) SetPositionX(x float64) *Camera {
// // 	// cPos.X = x
// // 	return c
// // }

// // func (c *Camera) SetPositionY(y float64) *Camera {
// // 	cPos.Y = y
// // 	return c
// // }

// // func (c *Camera) SetZeroPositionX() {
// // 	c.SetPositionX(c.getCameraTranslationX(0) / c.Scale.X)
// // }

// // func (c *Camera) SetZeroPositionY() {
// // 	sHUD := screen.UseScreen().GetHUDOffset()
// // 	c.SetPositionY(c.getCameraTranslationY(0)/c.Scale.X - sHUD/c.Scale.X)
// // }

// // func (c *Camera) GetScale() types.Vec2 {
// // 	s := screen.UseScreen().GetSize()
// // 	// c.Scale
// // 	return types.Vec2{X: c.Scale.X * s.X / 100, Y: c.Scale.Y * s.Y / 100}
// // }

// // Rotate rotates by phi
// func (c *Camera) Rotate(phi float64) *Camera {
// 	c.Rot += phi
// 	return c
// }

// // SetRotation sets the rotation to rot
// func (c *Camera) SetRotation(rot float64) *Camera {
// 	c.Rot = rot
// 	return c
// }

// func (c *Camera) Zoom(mul float64) *Camera {
// 	if c.MinScale.X < c.Scale.X*mul && c.Scale.X*mul < c.MaxScale.X {

// 		previousScale := c.Scale
// 		c.Scale.X *= mul
// 		c.Scale.Y *= mul
// 		appliedCPos := c.GetCameraTranslation()
// 		appliedCPos.Y -= screen.UseScreen().GetHUDOffset()

// 		wM := world.UseWorld().GetWorldMap()
// 		wMSize := wM.GetSize()

// 		sAxis := screen.UseScreen().GetAxis()
// 		if (!math.Signbit(appliedCPos.X) || !math.Signbit(appliedCPos.Y)) ||
// 			((wMSize.X*c.Scale.X < wMSize.X) || (wMSize.Y*c.Scale.Y < wMSize.Y)) ||
// 			((math.Abs(appliedCPos.X) > wMSize.X*c.Scale.X-sAxis.X*2) || math.Abs(appliedCPos.Y) > wMSize.Y*c.Scale.X-sAxis.Y*2) {
// 			c.Scale = previousScale
// 		}
// 	}
// 	return c
// }

// //Applies scale and further translation for draw
// func (c *Camera) GetCameraOptions() *ebiten.DrawImageOptions {
// 	op := &ebiten.DrawImageOptions{}
// 	sAxis := screen.UseScreen().GetAxis()
// 	scale := c.Followed.GetScale()
// 	cPos := c.Followed.GetPosForCamera()

// 	op.GeoM.Rotate(c.Rot)
// 	op.GeoM.Translate(-sAxis.X, -sAxis.Y)
// 	op.GeoM.Translate(-cPos.X, -cPos.Y)
// 	op.GeoM.Scale(scale.X, scale.Y)

// 	op.GeoM.Translate((cPos.X - sAxis.X), (cPos.Y - sAxis.Y))
// 	op.GeoM.Translate(-sAxis.X/2, -(sAxis.Y/2 - screen.UseScreen().GetHUDOffset()/2))
// 	op.GeoM.Translate(-(cPos.X - sAxis.X), -(cPos.Y - sAxis.Y))
// 	op.GeoM.Scale(c.Scale.X, c.Scale.Y)
// 	op.GeoM.Translate(sAxis.X, sAxis.Y)
// 	// op.GeoM.Translate(sAxis.X*c.Scale.X, sAxis.Y*c.Scale.Y)
// 	// op.GeoM.Scale(1/scale.X, 1/scale.Y)
// 	fmt.Println(op.GeoM.String())
// 	return op
// }

// // func (c *Camera) getCameraTranslationX(x float64) float64 {
// // 	return c.getCameraTranslation(x, 0).X
// // }

// // func (c *Camera) getCameraTranslationY(y float64) float64 {
// // 	return c.getCameraTranslation(y, 0).Y
// // }

// func (c *Camera) getCameraTranslation(x, y float64) types.Vec2 {
// 	var rX, rY float64

// 	s := screen.UseScreen().GetAxis()
// 	// size := screen.UseScreen().GetSize()

// 	rX += -s.X
// 	rY += -s.Y

// 	rX += -x
// 	rY += -y

// 	rX += x - s.X
// 	rY += y - s.Y

// 	rX += -s.X / 2
// 	rY += -(s.Y/2 - screen.UseScreen().GetHUDOffset()/2)

// 	rX += -(x - s.X)
// 	rY += -(y - s.Y)

// 	rX *= c.Scale.X
// 	rY *= c.Scale.Y

// 	rX += s.X
// 	rY += s.Y

// 	rX += s.X * c.Scale.X
// 	rY += s.Y * c.Scale.Y

// 	return types.Vec2{X: rX, Y: rY}
// }

// func (c *Camera) GetCameraTranslation() types.Vec2 {
// 	cPos := c.Followed.GetPosForCamera()
// 	return c.getCameraTranslation(cPos.X, cPos.Y)
// }

// func (c *Camera) GetScreenCoordsTranslation(x, y float64) types.Vec2 {
// 	var rX, rY float64

// 	s := screen.UseScreen().GetAxis()

// 	rX += -s.X
// 	rY += -s.Y

// 	rX += x
// 	rY += y

// 	rX += -(x - s.X)
// 	rY += -(y - s.Y)

// 	rX += -s.X / 2
// 	rY += -(s.Y/2 - screen.UseScreen().GetHUDOffset()/2)

// 	rX += (x - s.X)
// 	rY += (y - s.Y)

// 	rX *= c.Scale.X
// 	rY *= c.Scale.Y

// 	rX += s.X
// 	rY += s.Y

// 	rX += s.X * c.Scale.X
// 	rY += s.Y * c.Scale.Y

// 	return types.Vec2{X: rX, Y: rY}
// }

// func (c *Camera) GetWorldCoordX(wc float64) float64 {
// 	cPos := c.GetCameraTranslation()
// 	return (wc * cPos.X) / -cPos.X
// }

// func (c *Camera) GetWorldCoordY(wc float64) float64 {
// 	cPos := c.GetCameraTranslation()
// 	return (wc * cPos.Y) / -cPos.Y
// }

// func (c *Camera) IsOuttaCoordX(x float64) bool {
// 	cPos := c.GetCameraTranslation()
// 	return -cPos.X > x
// }

// func (c *Camera) IsOuttaCoordY(y float64) bool {
// 	cPos := c.GetCameraTranslation()
// 	return -cPos.Y > y
// }

// func (c *Camera) IsLowerZeroCoordX() bool {
// 	cPos := c.GetCameraTranslation()
// 	return -cPos.X <= 0
// }

// func (c *Camera) IsLowerZeroCoordY() bool {
// 	cPos := c.GetCameraTranslation()
// 	sHUD := screen.UseScreen().GetHUDOffset()
// 	return -cPos.Y <= -sHUD
// }

// func UseCamera() *Camera {
// 	if instance == nil {
// 		instance = &Camera{
// 			Followed: &world.UseWorld().GetPC().Base,
// 			Scale:    types.Vec2{X: 2.5, Y: 2.5},
// 			MaxScale: types.Vec2{X: 3.2, Y: 3.2},
// 			MinScale: types.Vec2{X: 2.3, Y: 2.3},
// 		}
// 	}
// 	return instance
// }

import (
	"math"

	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
	"github.com/kvartborg/vector"
)

type Camera struct {
	Position           types.Vec3
	Zoom, Angle, Pitch float64

	Rotation types.Matrix4
}

func (c *Camera) ZoomIn(v float64) {
	c.Zoom += v
}

func (c *Camera) ZoomOut(v float64) {
	c.Zoom -= v
}

func (c *Camera) MoveAngle(v float64) {
	c.Angle += v
}

func (c *Camera) MovePitch(v float64) {
	c.Pitch += v
}

func (c *Camera) MovePositionX(v float64) {
	c.Position.X += v
}

func (c *Camera) MovePositionY(v float64) {
	c.Position.Y += v
}

func (c *Camera) MovePositionZ(v float64) {
	c.Position.Z += v
}

func (ca *Camera) Rotate(x, y, z, angle float64) {
	mat := types.Matrix4{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
	vector := vector.Vector{x, y, z}.Unit()
	s := math.Sin(angle)
	c := math.Cos(angle)
	m := 1 - c

	mat[0][0] = m*vector[0]*vector[0] + c
	mat[0][1] = m*vector[0]*vector[1] + vector[2]*s
	mat[0][2] = m*vector[2]*vector[0] - vector[1]*s

	mat[1][0] = m*vector[0]*vector[1] - vector[2]*s
	mat[1][1] = m*vector[1]*vector[1] + c
	mat[1][2] = m*vector[1]*vector[2] + vector[0]*s

	mat[2][0] = m*vector[2]*vector[0] + vector[1]*s
	mat[2][1] = m*vector[1]*vector[2] - vector[0]*s
	mat[2][2] = m*vector[2]*vector[2] + c

	ca.Rotation = mat.GetMultiplied(ca.Rotation)
}

func (c *Camera) GetProjection(sm *screen.ScreenManager) types.Matrix4 {
	w, h := sm.Image.Size()
	asr := float64(h) / float64(w)

	return types.Matrix4{
		{2 / (1*c.Zoom - (-1 * c.Zoom)), 0, 0, 0},
		{0, 2 / (asr*c.Zoom - (-asr * c.Zoom)), 0, 0},
		{0, 0, -2, 0},
		{0, 0, 0, 1},
	}
}

func (camera *Camera) GetView() types.Matrix4 {
	var mat types.Matrix4
	mat[3][0] = camera.Position.X
	mat[3][1] = camera.Position.Y
	mat[3][2] = camera.Position.Z

	return mat.GetMultiplied(camera.Rotation.GetTransposed())
}

func (camera *Camera) Render() {

}
