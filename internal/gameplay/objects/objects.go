package objects

import (
	"image"

	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/models"
	"github.com/YarikRevich/caching/pkg/zeroshifter"
	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
)

type Role int

const (
	PC Role = iota
	Weapon
	Ammo
	Element
)

type Animation struct {
	PositionBeforeAnimation image.Point
	FrameCount              uint32
	FrameDelay              uint32
	FrameDelayCounter       uint32
	CurrentFrameMatrix      []float64
}

type Skin struct {
	Animation Animation
	Image    *ebiten.Image `json:"-"`
	Metadata *models.Metadata
}

/*The object structure which describes
each object on the map
*/

type Object struct {
	Skin

	//Names parentid the object referes to
	ParentID uuid.UUID

	ID uuid.UUID

	RawPos struct {
		X, Y float64
	}

	PositionHistory zeroshifter.IZeroShifter `json:"-"`

	Role Role

}

func (o *Object) savePositionHistory() {
	o.PositionHistory.Add(struct{ X, Y float64 }{X: o.RawPos.X, Y: o.RawPos.Y})
}

func (o *Object) savePositionBeforeAnimation() {
	o.Equipment.Skin.Animation.PositionBeforeAnimation = image.Point{X: int(o.RawPos.X), Y: int(o.RawPos.Y)}
}


func (o *Object) SetX(x float64) {
	o.RawPos.X = x
	o.savePositionBeforeAnimation()
	
}

func (o *Object) SetY(y float64) {
	o.RawPos.Y = y
	o.savePositionBeforeAnimation()
}

func (o *Object) UpdatePositionChanges(){
	o.savePositionHistory()
}

func (o *Object) IsXChanged() bool {
	var prevX float64
	for _, v := range o.PositionHistory.Get() {
		pos := v.(struct{ X, Y float64 })
		if prevX == 0 {
			prevX = pos.X
			continue
		}
		if prevX == pos.X {
			return false
		}
	}
	return true
}

func (o *Object) IsYChanged() bool {
	var prevY float64
	for _, v := range o.PositionHistory.Get() {
		pos := v.(struct{ X, Y float64 })

		if prevY == 0 {
			prevY = pos.Y
			continue
		}
		if prevY == pos.Y {
			return false
		}
	}
	return true
}