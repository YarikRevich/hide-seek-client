package camera

import (
	"fmt"
	"math"

	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	screenhistory "github.com/YarikRevich/HideSeek-Client/internal/core/screen"
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
	mapAxis struct {
		X, Y float64
	}

	// sleepZones []struct {
	// 	Min, Max image.Point
	// }
	zoom float64
	// size struct {
	// 	Width, Height float64
	// }

	//Map scale which updates deps between
	//standard coefficient map background scales
	mapScale struct {
		X, Y float64
	}

	lastMapScale struct {
		X, Y float64
	}

	//Hero scale which updates deps between
	//standard coefficient pc skin scales
	heroScale struct {
		X, Y float64
	}

	maxHeroScale struct {
		X, Y float64
	}

	connectedHeroPos struct {
		X, Y float64
	}

	lastConnectedHeroPos struct {
		X, Y float64
	}

	scaledConnectedHeroPos struct {
		X, Y float64
	}

	//States if pc crossed x axis
	isHeroMovementBlockedX bool

	//States if pc crossed y axis
	isHeroMovementBlockedY bool

	scaledMapTranslation struct {
		X, Y float64
	}

	lastScaledMapTranslation struct {
		X, Y float64
	}

	scaledHeroTranslation struct {
		X, Y float64
	}

	lastHeroScale struct {
		X, Y float64
	}

	lastHeroTranslation struct {
		X, Y float64
	}
}

func (c *Camera) isHeroTranslationBlocked() bool {
	return c.isHeroMovementBlockedX || c.isHeroMovementBlockedY
}

//Checks if pc has crossed the X axis
func (c *Camera) isCrossedAxisX() bool {
	return (int(c.scaledHeroTranslation.X) - 5) <= int(c.mapAxis.X) && int(c.mapAxis.X) <= (int(c.scaledHeroTranslation.X) + 5)
}

//Checks if pc has crossed the Y axis
func (c *Camera) isCrossedAxisY() bool {
	return (int(c.scaledHeroTranslation.Y) - 5) <= int(c.mapAxis.Y) && int(c.mapAxis.Y) <= (int(c.scaledHeroTranslation.Y) + 5)
}

//Updates scale coeffients for map matrix
func (c *Camera) updateMapScale() {
	w := objects.UseObjects().World()

	
	sx, sy := w.GetMapScale(screenhistory.GetScreen().Size())

	c.mapScale.X = ((sx + w.Metadata.Scale.CoefficiantX) / 100 * c.zoom) * 3
	c.mapScale.Y = ((sy + w.Metadata.Scale.CoefficiantY) / 100 * c.zoom) * 3
}

//Updates scale coeffients for hero matrix
func (c *Camera) updateHeroScale() {
	p := objects.UseObjects().PC()
	c.heroScale.X = (p.Metadata.Scale.CoefficiantX / 100 * c.zoom)
	c.heroScale.Y = (p.Metadata.Scale.CoefficiantY / 100 * c.zoom)
}

//Clears all the metrics for map and hero matrices
func (c *Camera) clearMatrices() {
	c.MapMatrix = ebiten.GeoM{}
	c.HeroMatrix = ebiten.GeoM{}
}

//Updates general metrics for map matrix
func (c *Camera) updateMapMatrix() {
	c.MapMatrix.Scale(float64(c.mapScale.X), float64(c.mapScale.Y))

	// w := objects.UseObjects().World()

	// fmt.Println()
	// if c.isHeroTranslationBlocked() {
	// 	if c.isHeroMovementBlockedX {
	// 		if (c.scaledHeroTranslation.X+c.scaledConnectedHeroPos.X <= w.Metadata.Size.Width*c.mapScale.X) &&
	// 			(c.scaledHeroTranslation.X-c.scaledConnectedHeroPos.X) >= 0 {
	// 			c.MapMatrix.Translate(-(c.scaledHeroTranslation.X - c.scaledConnectedHeroPos.X), 0)
	// 		} else {
	// 			c.isHeroMovementBlockedX = false
	// 			c.scaledMapTranslation.X = (c.scaledHeroTranslation.X - c.scaledConnectedHeroPos.X)
	// 			c.scaledMapTranslation.Y = 0

	// 		}
	// 	}

	// 	if c.isHeroMovementBlockedY {
	// 		if (c.scaledHeroTranslation.Y+c.scaledConnectedHeroPos.Y <= w.Metadata.Size.Height*c.mapScale.Y) &&
	// 			(c.scaledHeroTranslation.Y-c.scaledConnectedHeroPos.Y) >= 0 {
	// 			c.MapMatrix.Translate(0, -(c.scaledHeroTranslation.Y - c.scaledConnectedHeroPos.Y))
	// 		} else {
	// 			c.isHeroMovementBlockedY = false
	// 			c.scaledMapTranslation.X = 0
	// 			c.scaledMapTranslation.Y = (c.scaledHeroTranslation.Y - c.scaledConnectedHeroPos.Y)

	// 		}
	// 	}
	// }

	w := objects.UseObjects().World()
	p := objects.UseObjects().PC()
	fmt.Println("JESTEM TU", c.scaledHeroTranslation, p.RawPos, c.mapScale, )
	fmt.Println(w.GetMapScale(screenhistory.GetScreen().Size()))
	c.MapMatrix.Translate(-c.scaledHeroTranslation.X, -c.scaledHeroTranslation.Y)
}

func (c *Camera) updateScaledHeroTranslation() {
	c.scaledHeroTranslation.X = c.lastHeroTranslation.X * c.heroScale.X / c.lastHeroScale.X
	c.scaledHeroTranslation.Y = c.lastHeroTranslation.Y * c.heroScale.Y / c.lastHeroScale.Y
}

func (c *Camera) updateLastHeroTranslation() {
	p := objects.UseObjects().PC()

	if math.IsNaN(c.lastHeroTranslation.X) || math.IsNaN(c.lastHeroTranslation.Y) {
		c.lastHeroTranslation.X = p.RawPos.X
		c.lastHeroTranslation.Y = p.RawPos.Y
	} else {
		c.lastHeroTranslation.X = c.scaledHeroTranslation.X
		c.lastHeroTranslation.Y = c.scaledHeroTranslation.Y
	}

	if p.IsXChanged() || p.IsYChanged() || p.IsAnimatied() {
		c.lastHeroTranslation.X = p.RawPos.X * c.heroScale.X / c.maxHeroScale.X
		c.lastHeroTranslation.Y = p.RawPos.Y * c.heroScale.Y / c.maxHeroScale.Y
	}

}

//Saves last hero scale
func (c *Camera) updateLastHeroScale() {
	c.lastHeroScale = c.heroScale
}

//Saves max hero scale which is used for
//scaled translation calculation after pc
//had moved
func (c *Camera) saveMaxHeroScale() {
	p := objects.UseObjects().PC()
	c.maxHeroScale.X = (p.Metadata.Scale.CoefficiantX / 100 * 55) * 3
	c.maxHeroScale.Y = (p.Metadata.Scale.CoefficiantY / 100 * 55) * 3
}

func (c *Camera) updateLastConnectedPos() {
	if c.lastConnectedHeroPos.X == 0 || c.lastConnectedHeroPos.Y == 0 {
		c.lastConnectedHeroPos = c.connectedHeroPos
	} else {
		c.lastConnectedHeroPos = c.scaledConnectedHeroPos
	}
}

func (c *Camera) updateScaledMapTranslation() {
	c.scaledMapTranslation.X = c.lastScaledMapTranslation.X * c.mapScale.X / c.lastMapScale.X
	c.scaledMapTranslation.Y = c.lastScaledMapTranslation.Y * c.mapScale.Y / c.lastMapScale.Y
}

func (c *Camera) updateLastScaledMapTranslation() {
	if c.lastScaledMapTranslation.X == 0 && c.lastScaledMapTranslation.Y == 0 {
		p := objects.UseObjects().PC()
		c.scaledMapTranslation.X = p.RawPos.X * c.heroScale.X / c.maxHeroScale.X
		c.scaledMapTranslation.Y = p.RawPos.Y * c.heroScale.Y / c.maxHeroScale.Y
	} else {
		c.lastScaledMapTranslation = c.scaledMapTranslation
	}
}

func (c *Camera) updateScaledConnectedPos() {
	c.scaledConnectedHeroPos.X = c.lastConnectedHeroPos.X * c.heroScale.X / c.lastHeroScale.X
	c.scaledConnectedHeroPos.Y = c.lastConnectedHeroPos.Y * c.heroScale.Y / c.lastHeroScale.Y
}

//Updates general metrics for hero matrix
func (c *Camera) updateHeroMatrix() {
	p := objects.UseObjects().PC()
	c.HeroMatrix.Scale(p.GetMovementRotation(), 1)
	c.HeroMatrix.Scale(c.heroScale.X, c.heroScale.Y)

	if !c.isHeroMovementBlockedX && c.isCrossedAxisX() {
		c.connectedHeroPos.X = c.scaledHeroTranslation.X
		c.isHeroMovementBlockedX = true
	}

	if !c.isHeroMovementBlockedY && c.isCrossedAxisY() {
		c.connectedHeroPos.Y = c.scaledHeroTranslation.Y
		c.isHeroMovementBlockedY = true
	}

	// if !c.isHeroTranslationBlocked() {
	// 	c.HeroMatrix.Translate(c.scaledHeroTranslation.X, c.scaledHeroTranslation.Y)
	// } else {
	// 	if c.isHeroMovementBlockedX && c.isHeroMovementBlockedY {
	// 		c.HeroMatrix.Translate(c.scaledConnectedHeroPos.X, c.scaledConnectedHeroPos.Y)
	// 	}

	// 	if c.isHeroMovementBlockedX && !c.isHeroMovementBlockedY {
	// 		c.HeroMatrix.Translate(c.scaledConnectedHeroPos.X, c.scaledHeroTranslation.Y)
	// 	}

	// 	if !c.isHeroMovementBlockedX && c.isHeroMovementBlockedY {
	// 		c.HeroMatrix.Translate(c.scaledHeroTranslation.X, c.scaledConnectedHeroPos.Y)
	// 	}
	// }
}

func (c *Camera) updateMapAxis() {
	w := objects.UseObjects().World()
	a := (w.Metadata.RawSize.Width - (w.Metadata.RawSize.Width / (c.mapScale.X + w.Metadata.Scale.CoefficiantX))) / 2
	b := (w.Metadata.Size.Width - (w.Metadata.Size.Width / (c.mapScale.X + w.Metadata.Scale.CoefficiantX))) / 2
	c.mapAxis.X = (a + b) / 2

	a = (w.Metadata.RawSize.Height - (w.Metadata.RawSize.Height / (c.mapScale.Y + w.Metadata.Scale.CoefficiantY))) / 2
	b = (w.Metadata.Size.Height - (w.Metadata.Size.Height / (c.mapScale.Y + w.Metadata.Scale.CoefficiantY))) / 2
	c.mapAxis.Y = (a + b) / 2.2
}

//Updates the properties used for other
//camera calculations

func (c *Camera) updatePreDeps() {
	objects.UseObjects().PC().SaveLastPosition()

	c.updateMapAxis()

	c.updateScaledHeroTranslation()
	c.updateLastHeroTranslation()
	c.updateLastHeroScale()

	c.updateScaledMapTranslation()
	c.updateLastScaledMapTranslation()
}
func (c *Camera) updatePostDeps() {
	c.updateLastConnectedPos()
	c.updateScaledConnectedPos()
}

//Updates camera properties
func (c *Camera) UpdateCamera(screen *ebiten.Image) {
	c.updateMapScale()
	c.updateHeroScale()

	c.clearMatrices()

	c.updatePreDeps()

	c.updateMapMatrix()
	c.updateHeroMatrix()

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
	if c.zoom < 55 {
		c.zoom++
	}
}

//Decrements zoom property
func (c *Camera) ZoomOut() {
	if c.zoom > 15 {
		c.zoom--
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
			zoom: 50}
		instance.saveMaxHeroScale()
	}
	return instance
}
