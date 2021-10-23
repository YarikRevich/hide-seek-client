package camera

import (
	"image"
	"math"

	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/pc"
	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/world"
	"github.com/hajimehoshi/ebiten/v2"
)

var instance *Camera

type Camera struct {
	HeroMatrix, MapMatrix ebiten.GeoM

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

	ConnectedPos struct {
		X, Y float64
	}

	//States if pc crossed any axis
	IsHeroMovementBlocked bool

	ScaledTranslation struct {
		X, Y float64
	}

	LastHeroScale struct {
		X, Y float64
	}

	LastHeroTranslation struct {
		X, Y float64
	}
	// Transmition []struct {
	// 	X, Y float64
	// }
	LastZoom float64

	InitialZoomInBreakpoint float64
	LastZoomIn              float64

	InitialZoomOutBreakpoint float64
	LastZoomOut              float64
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

func (c *Camera) attachHero() {}

func (c *Camera) disconnectHero() {

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

func (c *Camera) IsCrossedAxisX() bool {
	p := pc.UsePC()
	w := world.UseWorld()

	return int(p.RawPos.X) == int(w.Metadata.Size.Width/c.MapScale.X)
}

func (c *Camera) IsCrossedAxisY() bool {
	p := pc.UsePC()
	w := world.UseWorld()

	return int(p.RawPos.Y) == int(w.Metadata.Size.Height/c.MapScale.Y)
}

//Updates scale coeffients
func (c *Camera) UpdateMapScale(screen *ebiten.Image) {
	w := world.UseWorld()

	sx, sy := w.GetMapScale(screen.Size())

	c.MapScale.X = (w.Metadata.RawSize.Width / 100 * c.Zoom) * sx / 100
	c.MapScale.Y = (w.Metadata.RawSize.Height / 100 * c.Zoom) / sy / 1.5 / 100
}

func (c *Camera) UpdateHeroScale() {
	p := pc.UsePC()
	c.HeroScale.X = (p.Metadata.Scale.CoefficiantX / 100 * c.Zoom) * 3
	c.HeroScale.Y = (p.Metadata.Scale.CoefficiantY / 100 * c.Zoom) * 3
}

func (c *Camera) ClearMatrices() {
	c.MapMatrix = ebiten.GeoM{}
	c.HeroMatrix = ebiten.GeoM{}
}

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
}

func (c *Camera) saveScaledTranslation() {
	c.ScaledTranslation.X = c.LastHeroTranslation.X * c.HeroScale.X / c.LastHeroScale.X
	c.ScaledTranslation.Y = c.LastHeroTranslation.Y * c.HeroScale.Y / c.LastHeroScale.Y

	// fmt.Println(c.ScaledTranslation, "SCALED TRANSLATION")
	// fmt.Println(c.LastHeroTranslation, "LAST HERO TRANSLATION")
	
	// fmt.Println(c.HeroScale, "HERO SCALE")

	// fmt.Println(c.LastHeroScale, "LAST SCALE")
}

func (c *Camera) saveLastHeroTranslation() {
	p := pc.UsePC()

	// fmt.Println(p.RawPos.X, c.HeroScale.X, c.LastHeroScale.X, "CHANGED")

	if math.IsNaN(c.LastHeroTranslation.X) || math.IsNaN(c.LastHeroTranslation.Y) {
		c.LastHeroTranslation.X = p.RawPos.X
		c.LastHeroTranslation.Y = p.RawPos.Y
	} else {
		c.LastHeroTranslation.X = c.ScaledTranslation.X
		c.LastHeroTranslation.Y = c.ScaledTranslation.Y
	}

	if p.IsXChanged() || p.IsYChanged() {
		c.LastHeroTranslation.X = p.RawPos.X * c.HeroScale.X / c.MaxHeroScale.X
		c.LastHeroTranslation.Y = p.RawPos.Y * c.HeroScale.Y / c.MaxHeroScale.Y
	}

}

func (c *Camera) saveLastHeroScale() {
	c.LastHeroScale.X = c.HeroScale.X
	c.LastHeroScale.Y = c.HeroScale.Y
}

func (c *Camera) saveMaxHeroScale(){
	p := pc.UsePC()
	c.MaxHeroScale.X = (p.Metadata.Scale.CoefficiantX / 100 * 35) * 3
	c.MaxHeroScale.Y = (p.Metadata.Scale.CoefficiantY / 100 * 35) * 3
}

func (c *Camera) UpdateHeroMatrix() {
	p := pc.UsePC()
	c.HeroMatrix.Scale(p.GetMovementRotation(), 1)
	c.HeroMatrix.Scale(c.HeroScale.X, c.HeroScale.Y)

	if !c.IsHeroMovementBlocked {
		c.HeroMatrix.Translate(c.ScaledTranslation.X, c.ScaledTranslation.Y)
		// c.HeroMatrix.Translate(p.RawPos.X, p.RawPos.Y)
		// if c.IsCrossedAxisX() {
		// 	c.ConnectedPos.X = p.RawPos.X
		// 	c.IsHeroMovementBlocked = true
		// }

		// if c.IsCrossedAxisY() {
		// 	c.ConnectedPos.Y = p.RawPos.Y
		// 	c.IsHeroMovementBlocked = true
		// }
	} else {
		c.HeroMatrix.Translate(c.ConnectedPos.X, p.RawPos.Y)
	}
}

func (c *Camera) UpdateHistory() {
	pc.UsePC().UpdatePositionChanges()

	c.saveScaledTranslation()
	c.saveLastHeroScale()
	c.saveLastHeroTranslation()

}

//Updates camera properties
//-> Scale coefficients
//-> Sleep zones
func (c *Camera) UpdateCamera(screen *ebiten.Image) {
	c.UpdateMapScale(screen)
	c.UpdateHeroScale()
	// c.UpdateSleepZones()
	// c.UpdateCharachterTranslation()

	c.ClearMatrices()

	c.UpdateMapMatrix(screen)
	c.UpdateHeroMatrix()

	c.UpdateHistory()
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

func (c *Camera) ZoomIn() {
	if c.Zoom < 35 {
		// c.LastZoom = c.Zoom
		c.Zoom++
	}
}
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
			Zoom: 25}
		instance.saveMaxHeroScale()
	}
	return instance
}
