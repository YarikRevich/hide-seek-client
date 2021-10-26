package camera

import (
	"fmt"
	"image"
	"math"

	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/pc"
	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/world"
	"github.com/hajimehoshi/ebiten/v2"
)

var instance *Camera

type Camera struct {
	HeroMatrix, MapMatrix ebiten.GeoM

	//Boarders of the camera view user sees
	// MapViewBoarders struct {
	// 	Min, Max struct {
	// 		X, Y float64
	// 	}
	// }
	MapAxis struct {
		X, Y float64
	}

	SleepZones []struct {
		Min, Max image.Point
	}
	Zoom float64
	Size struct {
		Width, Height float64
	}

	//Map scale which updates deps between
	//standard coefficient map background scales
	MapScale struct {
		X, Y float64
	}
	//Hero scale which updates deps between
	//standard coefficient pc skin scales
	HeroScale struct {
		X, Y float64
	}

	MaxHeroScale struct {
		X, Y float64
	}

	ConnectedHeroPos struct {
		X, Y float64
	}

	LastConnectedHeroPos struct {
		X, Y float64
	}

	ScaledConnectedHeroPos struct {
		X, Y float64
	}

	//States if pc crossed x axis
	IsHeroMovementBlockedX bool

	//States if pc crossed y axis
	IsHeroMovementBlockedY bool

	ScaledMapTranslation struct {
		X, Y float64
	}

	ScaledHeroTranslation struct {
		X, Y float64
	}

	LastHeroScale struct {
		X, Y float64
	}

	LastHeroTranslation struct {
		X, Y float64
	}

	// LastZoom float64

	// InitialZoomInBreakpoint float64
	// LastZoomIn              float64

	// InitialZoomOutBreakpoint float64
	// LastZoomOut              float64
}

// //Checks if pc inside the camera view
// func (c *Camera) InCameraView(x, y float64) bool {
// 	return (x <= float64(c.CamBoarders.Max.X) && x >= float64(c.CamBoarders.Min.X)) &&
// 		(y <= float64(c.CamBoarders.Max.Y) && y >= float64(c.CamBoarders.Min.Y))
// }

// //Checks if camera view is outta passed coords
// func (c *Camera) isCameraOuttaCoords(x, y float64) bool {
// 	return false
// }

// //Checks if pc crossed min x
// func (c *Camera) isCrossedMinX(x float64) bool {
// 	return x <= float64(c.CamBoarders.Min.X)
// }

// //Checks if pc crossed min Y
// func (c *Camera) isCrossedMinY(y float64) bool {
// 	return y <= float64(c.CamBoarders.Min.Y)
// }

// //Checks if pc crossed max x
// func (c *Camera) isCrossedMaxX(x float64) bool {
// 	return x >= float64(c.CamBoarders.Max.X)
// }

// //Checks if pc crossed max Y
// func (c *Camera) isCrossedMaxY(y float64) bool {
// 	return y >= float64(c.CamBoarders.Max.Y)
// }

// p := pc.UsePC()
// w := world.UseWorld()

// wx := (w.Metadata.Size.Width * w.Metadata.Scale.CoefficiantX)
// wy := (w.Metadata.Size.Height * w.Metadata.Scale.CoefficiantY)

// cx := (w.Metadata.Size.Width * w.Metadata.Scale.CoefficiantX) / 100 * c.Zoom
// cy := (w.Metadata.Size.Height * w.Metadata.Scale.CoefficiantY) / 100 * c.Zoom
// // cc := math.Sqrt(math.Pow(cx, 2)+math.Pow(cy, 2)) / 2

// c.Position.X = p.X + (cx / 2)
// for (p.X != c.Position.X+(cx/2) && p.Y != c.Position.Y+(cy/2)) || c.isCameraOuttaCoords(wx, wy){
// 	if p.X > c.Position.X {
// 		c.Position.X++
// 	}else{
// 		c.Position.X--
// 	}

// 	if p.Y > c.Position.Y {
// 		c.Position.Y++
// 	}else{
// 		c.Position.Y--
// 	}
// }
func (c *Camera) isHeroTranslationBlocked() bool {
	return c.IsHeroMovementBlockedX || c.IsHeroMovementBlockedY
}

//Checks if pc at camera sleep zone
func (c *Camera) IsSleepZone(x, y float64) bool {
	for _, v := range c.SleepZones {
		if (x <= float64(v.Max.X) && x >= float64(v.Min.X)) &&
			(y <= float64(v.Max.Y) && y >= float64(v.Min.Y)) {
			return true
		}
	}

	return false
}

// func (c *Camera) IsGoneOuttaSleepZone() bool {
// 	return c.IsSleepZone(c.LastTranslation.X, c.LastTranslation.Y) && !c.IsSleepZone(c.Translation.X, c.Translation.Y)
// }

//Updates the sleep zones of the camera
func (c *Camera) UpdateSleepZones() {
	// w := world.UseWorld()

	// c.SleepZones = []struct{ Min, Max image.Point }{
	// 	{image.Point{X: 0,
	// 		Y: 0},
	// 		image.Point{X: w.Metadata.Size.Width / c.Scale.X / 2,
	// 			Y: int(w.Metadata.Size.Height) / c.Scale.Y / 2}},
	// 	{image.Point{X: int(w.Metadata.Size.Width-w.Metadata.Size.Width) / c.Scale.X / 2,
	// 		Y: 0},
	// 		image.Point{X: int(w.Metadata.Size.Width),
	// 			Y: int(w.Metadata.Size.Height) / c.Scale.Y / 2}},
	// 	{image.Point{X: 0,
	// 		Y: int(w.Metadata.Size.Height)},
	// 		image.Point{X: int(w.Metadata.Size.Width) / c.Scale.X / 2,
	// 			Y: int(w.Metadata.Size.Height)}},
	// 	{image.Point{X: int(w.Metadata.Size.Width-w.Metadata.Size.Width) / c.Scale.X / 2,
	// 		Y: int(w.Metadata.Size.Height)},
	// 		image.Point{X: int(w.Metadata.Size.Width),
	// 			Y: int(w.Metadata.Size.Height),
	// 		}},
	// }
}

//Checks if pc has crossed the X axis
func (c *Camera) IsCrossedAxisX() bool {
	return int(c.ScaledHeroTranslation.X) >= int(c.MapAxis.X)
}

//Checks if pc has crossed the Y axis
func (c *Camera) IsCrossedAxisY() bool {
	return int(c.ScaledHeroTranslation.Y) >= int(c.MapAxis.Y)
}

//Updates scale coeffients for map matrix
func (c *Camera) UpdateMapScale(screen *ebiten.Image) {
	w := world.UseWorld()

	sx, sy := w.GetMapScale(screen.Size())

	c.MapScale.X = ((sx + w.Metadata.Scale.CoefficiantX) / 100 * c.Zoom) * 3
	c.MapScale.Y = ((sy + w.Metadata.Scale.CoefficiantY) / 100 * c.Zoom) * 3
}

//Updates scale coeffients for hero matrix
func (c *Camera) UpdateHeroScale() {
	p := pc.UsePC()
	c.HeroScale.X = (p.Metadata.Scale.CoefficiantX / 100 * c.Zoom)
	c.HeroScale.Y = (p.Metadata.Scale.CoefficiantY / 100 * c.Zoom)
}

//Clears all the metrics for map and hero matrices
func (c *Camera) ClearMatrices() {
	c.MapMatrix = ebiten.GeoM{}
	c.HeroMatrix = ebiten.GeoM{}
}

//Updates general metrics for map matrix
func (c *Camera) UpdateMapMatrix(screen *ebiten.Image) {
	// p := pc.UsePC()
	// w := world.UseWorld()
	// if p.IsAnimatied() {
	// 	x, y := p.GetPositionBeforeAnimation()
	// 	c.MapMatrix.Translate(-x, -y)
	// } else {
	// 	c.MapMatrix.Translate(-p.X, -p.Y)
	// }

	// fmt.Println(c.Scale)

	c.MapMatrix.Scale(float64(c.MapScale.X), float64(c.MapScale.Y))

	w := world.UseWorld()
	if c.isHeroTranslationBlocked() {
		fmt.Println()
		if (c.ScaledHeroTranslation.X+c.ScaledConnectedHeroPos.X < w.Metadata.Size.Width*c.MapScale.X) &&
			(c.ScaledHeroTranslation.X-c.ScaledConnectedHeroPos.X) > 0 {
			c.MapMatrix.Translate(-(c.ScaledHeroTranslation.X - c.ScaledConnectedHeroPos.X), -(c.ScaledHeroTranslation.Y - c.ScaledConnectedHeroPos.Y))
		} else {
			c.IsHeroMovementBlockedX = false
		}

		// if (c.ScaledHeroTranslation.Y+c.ScaledConnectedHeroPos.Y < w.Metadata.Size.Height*c.MapScale.Y) &&
		// 	(c.ScaledHeroTranslation.Y-c.ScaledConnectedHeroPos.Y) > 0 {
		// 	c.MapMatrix.Translate(0, -(c.ScaledHeroTranslation.Y - c.ScaledConnectedHeroPos.Y))
		// } else {
		// 	c.IsHeroMovementBlockedY = false
		// }

		// if (c.ScaledHeroTranslation.Y+c.ScaledConnectedHeroPos.Y > w.Metadata.Size.Height*c.MapScale.Y) &&
		// 	(c.ScaledHeroTranslation.Y-c.ScaledConnectedHeroPos.Y) < 0 {
		// 	c.IsHeroMovementBlockedY = false
		// 	return
		// }

		// fmt.Println(c.IsHeroMovementBlockedX, c.IsHeroMovementBlockedY)

		// fmt.Println(-(c.ScaledHeroTranslation.Y - c.ScaledConnectedHeroPos.Y))

	}
}

func (c *Camera) updateScaledHeroTranslation() {
	c.ScaledHeroTranslation.X = c.LastHeroTranslation.X * c.HeroScale.X / c.LastHeroScale.X
	c.ScaledHeroTranslation.Y = c.LastHeroTranslation.Y * c.HeroScale.Y / c.LastHeroScale.Y
}

func (c *Camera) updateLastHeroTranslation() {
	p := pc.UsePC()

	if math.IsNaN(c.LastHeroTranslation.X) || math.IsNaN(c.LastHeroTranslation.Y) {
		c.LastHeroTranslation.X = p.RawPos.X
		c.LastHeroTranslation.Y = p.RawPos.Y
	} else {
		c.LastHeroTranslation.X = c.ScaledHeroTranslation.X
		c.LastHeroTranslation.Y = c.ScaledHeroTranslation.Y
	}

	if p.IsXChanged() || p.IsYChanged() || p.IsAnimatied() {
		c.LastHeroTranslation.X = p.RawPos.X * c.HeroScale.X / c.MaxHeroScale.X
		c.LastHeroTranslation.Y = p.RawPos.Y * c.HeroScale.Y / c.MaxHeroScale.Y
	}

}

//Saves last hero scale
func (c *Camera) updateLastHeroScale() {
	c.LastHeroScale.X = c.HeroScale.X
	c.LastHeroScale.Y = c.HeroScale.Y
}

//Saves max hero scale which is used for
//scaled translation calculation after pc
//had moved
func (c *Camera) saveMaxHeroScale() {
	p := pc.UsePC()
	c.MaxHeroScale.X = (p.Metadata.Scale.CoefficiantX / 100 * 55) * 3
	c.MaxHeroScale.Y = (p.Metadata.Scale.CoefficiantY / 100 * 55) * 3
}

func (c *Camera) updateLastConnectedPos() {
	// fmt.Println(c.LastConnectedPos, c.ConnectedPos)
	// && (!math.IsNaN(c.ConnectedHeroPos.X) && !math.IsNaN(c.ConnectedHeroPos.Y))
	if c.LastConnectedHeroPos.X == 0 || c.LastConnectedHeroPos.Y == 0 {
		c.LastConnectedHeroPos.X = c.ConnectedHeroPos.X
		c.LastConnectedHeroPos.Y = c.ConnectedHeroPos.Y
	} else {
		c.LastConnectedHeroPos.X = c.ScaledConnectedHeroPos.X
		c.LastConnectedHeroPos.Y = c.ScaledConnectedHeroPos.Y
	}
}

func (c *Camera) updateScaledConnectedPos() {
	c.ScaledConnectedHeroPos.X = c.LastConnectedHeroPos.X * c.HeroScale.X / c.LastHeroScale.X
	c.ScaledConnectedHeroPos.Y = c.LastConnectedHeroPos.Y * c.HeroScale.Y / c.LastHeroScale.Y
}

//Updates general metrics for hero matrix
func (c *Camera) UpdateHeroMatrix() {
	p := pc.UsePC()
	c.HeroMatrix.Scale(p.GetMovementRotation(), 1)
	c.HeroMatrix.Scale(c.HeroScale.X, c.HeroScale.Y)

	if !c.isHeroTranslationBlocked() {
		c.HeroMatrix.Translate(c.ScaledHeroTranslation.X, c.ScaledHeroTranslation.Y)
	} else {
		// if c.IsHeroMovementBlockedX && !c.IsHeroMovementBlockedY {
		c.HeroMatrix.Translate(c.ScaledConnectedHeroPos.X, c.ScaledConnectedHeroPos.Y)
		// } else if c.IsHeroMovementBlockedY && !c.IsHeroMovementBlockedX {
		// 	c.HeroMatrix.Translate(c.ScaledHeroTranslation.X, c.ScaledConnectedHeroPos.Y)
		// } else if c.IsHeroMovementBlockedX && c.IsHeroMovementBlockedY {
		// 	c.HeroMatrix.Translate(c.ScaledConnectedHeroPos.X, c.ScaledConnectedHeroPos.Y)
		// }
	}

	if !c.IsHeroMovementBlockedX {
		if c.IsCrossedAxisX() {
			c.ConnectedHeroPos.X = c.ScaledHeroTranslation.X
			c.IsHeroMovementBlockedX = true
		}
	}

	// if !c.IsHeroMovementBlockedY {
	// 	if c.IsCrossedAxisY() {
	// 		c.ConnectedHeroPos.Y = c.ScaledHeroTranslation.Y
	// 		c.IsHeroMovementBlockedY = true
	// 	}
	// }
}

func (c *Camera) updateMapAxis() {
	w := world.UseWorld()
	a := (w.Metadata.RawSize.Width - (w.Metadata.RawSize.Width / (c.MapScale.X + w.Metadata.Scale.CoefficiantX))) / 2
	b := (w.Metadata.Size.Width - (w.Metadata.Size.Width / (c.MapScale.X + w.Metadata.Scale.CoefficiantX))) / 2
	c.MapAxis.X = (a + b) / 2

	a = (w.Metadata.RawSize.Height - (w.Metadata.RawSize.Height / (c.MapScale.Y + w.Metadata.Scale.CoefficiantY))) / 2
	b = (w.Metadata.Size.Height - (w.Metadata.Size.Height / (c.MapScale.Y + w.Metadata.Scale.CoefficiantY))) / 2
	c.MapAxis.Y = (a + b) / 2.2
}

//Updates the properties used for other
//camera calculations

func (c *Camera) updatePreDeps() {
	pc.UsePC().UpdatePositionChanges()

	c.updateMapAxis()
	c.updateScaledHeroTranslation()

	c.updateLastHeroTranslation()
	c.updateLastHeroScale()
}
func (c *Camera) updatePostDeps() {

	c.updateLastConnectedPos()
	c.updateScaledConnectedPos()
}

//Updates camera properties
func (c *Camera) UpdateCamera(screen *ebiten.Image) {
	c.UpdateMapScale(screen)
	c.UpdateHeroScale()

	c.ClearMatrices()

	c.updatePreDeps()

	c.UpdateMapMatrix(screen)
	c.UpdateHeroMatrix()

	c.updatePostDeps()
}

// func (c *Camera) Disconnect
//Moves camera if position of the pc gets outta boarders of the cam
// func (c *Camera) MoveIfBoarderCrossed() {
// 	p := pc.UsePC()

// 	const maxCoefficient = 20
// 	var coefficient float64

// 	if c.isCrossedMinX(p.X) {
// 		coefficient = maxCoefficient
// 		if d := math.Abs(float64(c.CamBoarders.Min.X - c.MapBoarders.Min.X)); d > coefficient {
// 			coefficient = d
// 		}
// 	}

// 	if c.isCrossedMaxX(p.X) {
// 		coefficient = maxCoefficient
// 		if d := math.Abs(float64(c.CamBoarders.Max.X - c.MapBoarders.Max.X)); d > coefficient {
// 			coefficient = d
// 		}
// 	}

// 	if c.isCrossedMinY(p.Y) {
// 		coefficient = maxCoefficient
// 		if d := math.Abs(float64(c.CamBoarders.Min.Y - c.MapBoarders.Min.Y)); d > coefficient {
// 			coefficient = d
// 		}
// 	}

// 	if c.isCrossedMaxY(p.Y) {
// 		coefficient = maxCoefficient
// 		if d := math.Abs(float64(c.CamBoarders.Max.Y - c.MapBoarders.Max.Y)); d > coefficient {
// 			coefficient = d
// 		}
// 	}
// }

// func (c *Camera) SyncCameraPosition() {
// 	// w := world.UseWorld()

// 	// if c.Position.X == 0{
// 	// c.Position.X = w.Metadata.Size.Width / c.Scale.X / 2
// 	// }

// 	// fmt.Println(p.Y <= float64(c.MapBoarders.Max.Y), p.Y, float64(c.MapBoarders.Max.Y))

// 	fmt.Println(p.X, c.Position.X)
// 	// if p.X > c.Position.X {
// 	//w.Metadata.Size.Width / c.Scale.X / 2
// 	// c.Position.X++
// 	// }
// 	// c.Position.X += p.Buffs.SpeedX

// 	// fmt.Println()
// }

//Increments zoom property
func (c *Camera) ZoomIn() {
	if c.Zoom < 55 {
		c.Zoom++
	}
}

//Decrements zoom property
func (c *Camera) ZoomOut() {
	if c.Zoom > 15 {
		c.Zoom--
	}
}

// func (c *Camera) GetCameraViewSize(screenW, screenH int) (float64, float64) {
// 	w := world.UseWorld()
// 	imageW := w.Metadata.Size.Width * w.Metadata.Scale.CoefficiantX
// 	imageH := w.Metadata.Size.Height * w.Metadata.Scale.CoefficiantY

// 	sx, sy := w.RelativeMapSizeScale(screenW, screenH)
// 	return imageW * sx / c.Scale.X, imageH * sy / c.Scale.Y
// }

//Returns camera view scale
// func (c *Camera) GetCameraViewScale(screenW, screenH int) (float64, float64) {
// 	// w := world.UseWorld()
// 	// sx, sy := w.GetMapScale(screenW, screenH)
// 	return c.Scale.X, c.Scale.Y
// }

// //Returns translation for camera view
// //related by camera scale
// func (c *Camera) GetCameraViewTranslation(sx, sy float64) (float64, float64) {
// 	p := pc.UsePC()
// 	w := world.UseWorld()
// 	return -(p.X - (w.Metadata.RawSize.Width / c.Scale.X / 2)), -(p.Y - (w.Metadata.RawSize.Height / c.Scale.Y / 2))
// }

// func (c *Camera) GetCharacterTranslation(screenW, screenH int) (float64, float64) {
// if c.IsSleepZone(c.Translation.X, c.Translation.Y) {
// 	return p.X, p.Y
// }
// return 0, 0

// p := pc.UsePC()
// w := world.UseWorld()

// cvx, cvy := c.GetCameraViewScale(screenW, screenH)

// fmt.Println(p.X / c.Scale.X,
// 	w.Metadata.Size.Width / cvx / c.Scale.X / 2 + (p.Metadata.RawSize.Width*2))

// 	p := pc.UsePC()
// 	if c.IsSleepZone(p.X, p.Y){
// 		return p.X, 110
// 		// return p.X / c.Scale.X,  w.Metadata.Size.Height / cvy / c.Scale.Y / 2 - p.Metadata.RawSize.Height
// 	}

// 	// return p.X, p.Y
// 		return 245, 110
// 	// return w.Metadata.RawSize.Width / cvx / c.Scale.X / 2 + (p.Metadata.RawSize.Width*2),
// 	// w.Metadata.RawSize.Height / cvy / c.Scale.Y / 2 - p.Metadata.RawSize.Height
// }

// func (c *Camera) GetCameraTranslation(cvx, cvy float64) (float64, float64) {
// 	p := pc.UsePC()

// 	if c.IsGoneOuttaSleepZone() {
// 		c.Translation.X = p.X
// 		c.Translation.Y = p.Y
// 	}

// 	if !c.IsSleepZone(p.X, p.Y) {
// 		c.LastTranslation.X = c.Translation.X
// 		c.LastTranslation.Y = c.Translation.Y

// 		c.Translation.X = p.X
// 		c.Translation.Y = p.Y
// 	}

// 	// if p.X > 300{
// 	// }
// 	// if len(c.Transmition) != 0{
// 	// 	r := c.Transmition[len(c.Transmition)-1]
// 	// 	c.Transmition = c.Transmition[:len(c.Transmition)-1]
// 	// 	return r.X, r.Y
// 	// }

// 	return -c.Translation.X, -c.Translation.Y
// 	// return -c.Position.X, -c.Position.Y
// initialCameraPos := p.X + 20

// fmt.Println()

// imageW, imageH := w.Location.Image.Size()
// screenW, screenH := screen.Size()
// cx, cy := -p.Buffs.SpeedX * p.X, -p.Buffs.SpeedY * p.Y

// fmt.Println(p.IsXChanged(), p.IsYChanged(), cx, cy)
// fmt.Println(cy, (w.Metadata.Size.Height * w.Metadata.Scale.CoefficiantY), !p.IsXChanged())

// if cy > (w.Metadata.Size.Height * w.Metadata.Scale.CoefficiantY) && !p.IsXChanged(){
// 	return 0, 0
// }

// if cx > (w.Metadata.Size.Width * w.Metadata.Scale.CoefficiantX) && !p.IsYChanged(){
// 	return 0, 0
// }

// if cx > (w.Metadata.Size.Width * w.Metadata.Scale.CoefficiantX) && cy < (w.Metadata.Size.Height * w.Metadata.Scale.CoefficiantY){
// 	return 0, cy
// }

// if cy > (w.Metadata.Size.Height * w.Metadata.Scale.CoefficiantY) && cx < (w.Metadata.Size.Width * w.Metadata.Scale.CoefficiantX){
// 	return cx, 0
// }

// return cx, cy
// }

//Uses or creates a new instance of camera
func UseCamera() *Camera {
	if instance == nil {
		instance = &Camera{
			Zoom: 50}
		instance.saveMaxHeroScale()
	}
	return instance
}
