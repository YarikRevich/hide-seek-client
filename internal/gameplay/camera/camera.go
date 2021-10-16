package camera

import (
	"fmt"
	"image"
	"math"

	// "github.com/YarikRevich/HideSeek-Client/internal/gameplay/pc"
	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/pc"
	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/world"
	"github.com/YarikRevich/HideSeek-Client/internal/screen"
)

var instance *Camera

type Camera struct {
	MapBoarders struct {
		Min, Max image.Point
	}
	CamBoarders struct {
		Min, Max image.Point
	}
	Zoom     float64
	Position struct {
		X, Y float64
	}
}

//Checks if pc inside the camera view
func (c *Camera) InCameraView(x, y float64) bool {
	return (x <= float64(c.CamBoarders.Max.X) && x >= float64(c.CamBoarders.Min.X)) &&
		(y <= float64(c.CamBoarders.Max.Y) && y >= float64(c.CamBoarders.Min.Y))
}

//Checks if camera view is outta passed coords
func (c *Camera) isCameraOuttaCoords(x, y float64) bool {
	return false
}

//Checks if pc crossed min x
func (c *Camera) isCrossedMinX(x float64) bool {
	return x <= float64(c.CamBoarders.Min.X)
}

//Checks if pc crossed min Y
func (c *Camera) isCrossedMinY(y float64) bool {
	return y <= float64(c.CamBoarders.Min.Y)
}

//Checks if pc crossed max x
func (c *Camera) isCrossedMaxX(x float64) bool {
	return x >= float64(c.CamBoarders.Max.X)
}

//Checks if pc crossed max Y
func (c *Camera) isCrossedMaxY(y float64) bool {
	return y >= float64(c.CamBoarders.Max.Y)
}

//Inits camera view due to the pc spawn position
func (c *Camera) InitCamera() {
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

	fmt.Println(c.Position)
}

//Moves camera if position of the pc gets outta boarders of the cam
func (c *Camera) MoveIfBoarderCrossed() {
	p := pc.UsePC()

	const maxCoefficient = 20
	var coefficient float64

	if c.isCrossedMinX(p.X) {
		coefficient = maxCoefficient
		if d := math.Abs(float64(c.CamBoarders.Min.X - c.MapBoarders.Min.X)); d > coefficient {
			coefficient = d
		}
	}

	if c.isCrossedMaxX(p.X) {
		coefficient = maxCoefficient
		if d := math.Abs(float64(c.CamBoarders.Max.X - c.MapBoarders.Max.X)); d > coefficient {
			coefficient = d
		}
	}

	if c.isCrossedMinY(p.Y) {
		coefficient = maxCoefficient
		if d := math.Abs(float64(c.CamBoarders.Min.Y - c.MapBoarders.Min.Y)); d > coefficient {
			coefficient = d
		}
	}

	if c.isCrossedMaxY(p.Y) {
		coefficient = maxCoefficient
		if d := math.Abs(float64(c.CamBoarders.Max.Y - c.MapBoarders.Max.Y)); d > coefficient {
			coefficient = d
		}
	}

	// if c.isCrossedX(p.X, p.Y) {

	// }
	// if p.X >= float64(c.CamBoarders.Max.X) || p.X <= float64(c.CamBoarders.Min.X){
	// }

	// if p.Y >= float64(c.CamBoarders.Max.Y) || p.Y <= float64(c.CamBoarders.Min.Y){

	// }
}

func (c *Camera) GetCameraPosition()(float64, float64){
	p := pc.UsePC()
	w := world.UseWorld()

			
		imageW, imageH := w.Location.Image.Size()
		screenW, screenH := screen.Size()
	cx, cy := -p.Buffs.SpeedX * p.X, -p.Buffs.SpeedY * p.Y

	// fmt.Println(p.IsXChanged(), p.IsYChanged(), cx, cy)
	fmt.Println(cy, (w.Metadata.Size.Height * w.Metadata.Scale.CoefficiantY), !p.IsXChanged())

	if cy > (w.Metadata.Size.Height * w.Metadata.Scale.CoefficiantY) && !p.IsXChanged(){
		return 0, 0
	}

	if cx > (w.Metadata.Size.Width * w.Metadata.Scale.CoefficiantX) && !p.IsYChanged(){
		return 0, 0
	}


	if cx > (w.Metadata.Size.Width * w.Metadata.Scale.CoefficiantX) && cy < (w.Metadata.Size.Height * w.Metadata.Scale.CoefficiantY){
		return 0, cy
	}

	if cy > (w.Metadata.Size.Height * w.Metadata.Scale.CoefficiantY) && cx < (w.Metadata.Size.Width * w.Metadata.Scale.CoefficiantX){
		return cx, 0
	}

	

	return cx, cy
}

//Uses or creates a new instance of camera
func UseCamera() *Camera {
	if instance == nil {
		w := world.UseWorld()

		instance = new(Camera)

		instance.Zoom = 40
		instance.MapBoarders.Min.X = 0
		instance.MapBoarders.Min.Y = int(w.Metadata.Size.Height)

		instance.MapBoarders.Max.X = int(w.Metadata.Size.Width)
		instance.MapBoarders.Max.Y = int(w.Metadata.Size.Height)
	}
	return instance
}
