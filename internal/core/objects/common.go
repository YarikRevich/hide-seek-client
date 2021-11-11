package objects

import (
	"image"
	"path/filepath"
	"unsafe"

	// "github.com/YarikRevich/HideSeek-Client/internal/core/keycodes"
	// imagecollecion "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/image_loader/collection"
	// metadatacollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"
	// metadatamodels "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/models"
	"github.com/YarikRevich/HideSeek-Client/internal/core/keycodes"
	"github.com/YarikRevich/HideSeek-Client/internal/core/networking/api"
	"github.com/YarikRevich/HideSeek-Client/internal/core/sources"
	"github.com/YarikRevich/caching/pkg/zeroshifter"
	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
)

type Role int

const (
	PCRole Role = iota
	ElementRole
	WeaponRole
	AmmoRole
)

type Physics struct {
	Jump []keycodes.Direction
}

type Animation struct {
	PositionBeforeAnimation image.Point
	FrameCount              uint32
	FrameDelayCounter       uint32
	CurrentFrameMatrix      []float64
}

type Skin struct {
	Name string
	Path string

	// Size struct {
	// 	Width, Height float64
	// } `json:"-"`

	// Buffs struct {
	// 	Speed struct {
	// 		X, Y float64
	// 	}
	// }

	// Margins struct {
	// 	LeftMargin, TopMargin float64
	// }

	// Scale struct {
	// 	CoefficiantX, CoefficiantY float64
	// }
}

/*The object structure which describes
each object on the map
*/
type Object struct {
	Animation
	Skin
	Physics

	// Names parentid the object referes to
	WorldID, ParentID, ID uuid.UUID

	RawPos struct {
		X, Y float64
	}

	Direction keycodes.Direction

	Spawn image.Point

	Role Role

	//Only client fields
	PositionHistory zeroshifter.IZeroShifter `json:"-"`
}

func (o *Object) IsByObject(ob Object) bool {
	return o.ParentID == ob.ID
}

func (o *Object) SaveLastPosition() {
	o.PositionHistory.Add(image.Point{X: int(o.RawPos.X), Y: int(o.RawPos.Y)})
}

func (o *Object) savePositionBeforeAnimation() {
	o.Animation.PositionBeforeAnimation = image.Point{X: int(o.RawPos.X), Y: int(o.RawPos.Y)}
}

func (o *Object) SetX(x float64) {
	if x < o.RawPos.X {
		o.Direction = keycodes.LEFT
	} else if x > o.RawPos.X {
		o.Direction = keycodes.RIGHT
	}

	o.RawPos.X = x
	o.savePositionBeforeAnimation()
}

func (o *Object) SetY(y float64) {
	if y < o.RawPos.Y {
		o.Direction = keycodes.UP
	} else if y > o.RawPos.Y {
		o.Direction = keycodes.DOWN
	}

	o.RawPos.Y = y
	o.savePositionBeforeAnimation()
}

//Checks if x pos has been changed
//in comparison to the last x poses
func (o *Object) IsXChanged() bool {
	var prevX int
	for _, v := range o.PositionHistory.Get() {
		pos := v.(image.Point)
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

//Checks if y pos has been changed
//in comparison to the last y poses
func (o *Object) IsYChanged() bool {
	var prevY int
	for _, v := range o.PositionHistory.Get() {
		pos := v.(image.Point)

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

//Sets spawn point for the object
func (o *Object) SetSpawn(spawns []image.Point) {
	o.RawPos.X = 500
	o.RawPos.Y = 500
}

//Checks if pc executes animation
func (o *Object) IsAnimatied() bool {
	return len(o.Physics.Jump) != 0
}

//Returns last saved position before animation was executed
func (o *Object) GetPositionBeforeAnimation() (float64, float64) {
	return float64(o.Animation.PositionBeforeAnimation.X),
		float64(o.Animation.PositionBeforeAnimation.Y)
}

func (o *Object) SetSkin(path string) {
	o.Path = path
	_, file := filepath.Split(path)
	o.Name = file
}

func (o *Object) GetImage() *ebiten.Image {
	return ebiten.NewImageFromImage(sources.UseSources().Images().GetImage(o.Path))
}

//Returns image where animation properties applied to
func (o *Object) GetAnimatedImage() *ebiten.Image{
	i := o.GetImage()
	m := o.GetMetadata().Modified

	sx, sy := int((m.Animation.FrameX+float64(o.Animation.FrameCount))*m.Animation.FrameWidth), int(m.Animation.FrameY)
	return i.SubImage(image.Rect(sx, sy, sx+int(m.Animation.FrameWidth), sy+int(m.Animation.FrameHeight))).(*ebiten.Image)
}

func (o *Object) GetMetadata() *sources.ModelCombination {
	return sources.UseSources().Metadata().GetMetadata(o.Path)
}

func (o *Object) ToAPIMessage() *api.Object{
	return &api.Object{
		Animation: &api.Animation{
			PositionBeforeAnimation: &api.Position{
				X: float64(o.Animation.PositionBeforeAnimation.X),
				Y: float64(o.Animation.PositionBeforeAnimation.Y),
			},
			FrameCount: o.FrameCount,
			FrameDelayCounter: o.FrameDelayCounter,
			CurrentFrameMatrix: o.CurrentFrameMatrix,
		},
		Skin: &api.Skin{
			Name: o.Name,
			Path: o.Path,
		},
		Physics: &api.Physics{
			Jump: *(*[]int64)(unsafe.Pointer(&o.Jump)),
		},
		WorldId: o.WorldID.String(),
		ParentId: o.ParentID.String(),
		Id: o.ID.String(),
		RawPos: &api.Position{
			X: o.RawPos.X,
			Y: o.RawPos.Y,
		},
		Spawn: &api.Position{
			X: float64(o.Spawn.X),
			Y: float64(o.Spawn.Y),
		},
		Direction: int64(o.Direction),
		Role: int64(o.Role),
	}
}

func (o *Object) FromAPIMessage(m *api.Object){
	o.Animation.PositionBeforeAnimation.X = int(m.Animation.PositionBeforeAnimation.X)
	o.Animation.PositionBeforeAnimation.Y = int(m.Animation.PositionBeforeAnimation.Y)
	o.Animation.FrameCount = m.Animation.FrameCount
	o.Animation.FrameDelayCounter = m.Animation.FrameDelayCounter
	o.Animation.CurrentFrameMatrix = m.Animation.CurrentFrameMatrix

	o.Skin.Name = m.Skin.Name
	o.Skin.Path = m.Skin.Path

	o.Physics.Jump = *(*[]keycodes.Direction)(unsafe.Pointer(&m.Physics.Jump))

	o.WorldID = uuid.MustParse(m.WorldId)
	o.ParentID = uuid.MustParse(m.ParentId)
	o.ID = uuid.MustParse(m.Id)

	o.RawPos.X = m.RawPos.X
	o.RawPos.Y = m.RawPos.Y
	o.Spawn.X = int(m.Spawn.X)
	o.Spawn.Y = int(m.Spawn.Y)
	o.Direction = keycodes.Direction(m.Direction)
	o.Role = Role(m.Role)
}

func NewObject() *Object {
	return new(Object)
}
