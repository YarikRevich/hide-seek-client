package objects

import (
	"image"
	"path/filepath"
	"unsafe"

	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
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
	FrameCount              uint64
	FrameDelayCounter       uint64
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

	RawPos, RawPosForCamera, AttachedPos struct {
		X, Y float64
	}

	Direction, SubDirection keycodes.Direction

	Spawn image.Point

	Role Role

	TranslationMovementXBlocked, TranslationMovementYBlocked bool

	//Only client fields
	PositionHistory zeroshifter.IZeroShifter
}

func (o *Object) UpdateDirection() {
	k := events.UseEvents().Keyboard()
	g := events.UseEvents().Gamepad()
	if g.IsGamepadConnected() {
		if g.AreGamepadButtonsCombinedInOrder(keycodes.GamepadDOWNButton, keycodes.GamepadRIGHTButton) {
			o.Direction = keycodes.DOWN
			o.SubDirection = keycodes.RIGHT
		} else if g.AreGamepadButtonsCombinedInOrder(keycodes.GamepadRIGHTButton, keycodes.GamepadDOWNButton) {
			o.Direction = keycodes.RIGHT
			o.SubDirection = keycodes.DOWN
		} else if g.AreGamepadButtonsCombinedInOrder(keycodes.GamepadDOWNButton, keycodes.GamepadLEFTButton) {
			o.Direction = keycodes.DOWN
			o.SubDirection = keycodes.LEFT
		} else if g.AreGamepadButtonsCombinedInOrder(keycodes.GamepadLEFTButton, keycodes.GamepadDOWNButton) {
			o.Direction = keycodes.LEFT
			o.SubDirection = keycodes.DOWN
		} else if g.AreGamepadButtonsCombinedInOrder(keycodes.GamepadUPButton, keycodes.GamepadRIGHTButton) {
			o.Direction = keycodes.UP
			o.SubDirection = keycodes.RIGHT
		} else if g.AreGamepadButtonsCombinedInOrder(keycodes.GamepadRIGHTButton, keycodes.GamepadUPButton) {
			o.Direction = keycodes.RIGHT
			o.SubDirection = keycodes.UP
		} else if g.AreGamepadButtonsCombinedInOrder(keycodes.GamepadUPButton, keycodes.GamepadLEFTButton) {
			o.Direction = keycodes.UP
			o.SubDirection = keycodes.LEFT
		} else if g.AreGamepadButtonsCombinedInOrder(keycodes.GamepadLEFTButton, keycodes.GamepadUPButton) {
			o.Direction = keycodes.LEFT
			o.SubDirection = keycodes.UP
		} else if g.IsGamepadButtonPressed(keycodes.GamepadUPButton) {
			o.Direction = keycodes.UP
			o.SubDirection = keycodes.NONE
		} else if g.IsGamepadButtonPressed(keycodes.GamepadLEFTButton) {
			o.Direction = keycodes.LEFT
			o.SubDirection = keycodes.NONE
		} else if g.IsGamepadButtonPressed(keycodes.GamepadRIGHTButton) {
			o.Direction = keycodes.RIGHT
			o.SubDirection = keycodes.NONE
		} else if g.IsGamepadButtonPressed(keycodes.GamepadDOWNButton) {
			o.Direction = keycodes.DOWN
			o.SubDirection = keycodes.NONE
		}
	} else {
		if k.AreKeysCombinedInOrder(ebiten.KeyS, ebiten.KeyD) || k.AreKeysCombinedInOrder(ebiten.KeyArrowDown, ebiten.KeyArrowRight) {
			o.Direction = keycodes.DOWN
			o.SubDirection = keycodes.RIGHT
		} else if k.AreKeysCombinedInOrder(ebiten.KeyD, ebiten.KeyS) || k.AreKeysCombinedInOrder(ebiten.KeyArrowRight, ebiten.KeyArrowDown) {
			o.Direction = keycodes.RIGHT
			o.SubDirection = keycodes.DOWN
		} else if k.AreKeysCombinedInOrder(ebiten.KeyS, ebiten.KeyA) || k.AreKeysCombinedInOrder(ebiten.KeyArrowDown, ebiten.KeyArrowLeft) {
			o.Direction = keycodes.DOWN
			o.SubDirection = keycodes.LEFT
		} else if k.AreKeysCombinedInOrder(ebiten.KeyA, ebiten.KeyS) || k.AreKeysCombinedInOrder(ebiten.KeyArrowLeft, ebiten.KeyArrowDown) {
			o.Direction = keycodes.LEFT
			o.SubDirection = keycodes.DOWN
		} else if k.AreKeysCombinedInOrder(ebiten.KeyW, ebiten.KeyD) || k.AreKeysCombinedInOrder(ebiten.KeyArrowUp, ebiten.KeyArrowRight) {
			o.Direction = keycodes.UP
			o.SubDirection = keycodes.RIGHT
		} else if k.AreKeysCombinedInOrder(ebiten.KeyD, ebiten.KeyW) || k.AreKeysCombinedInOrder(ebiten.KeyArrowRight, ebiten.KeyArrowUp) {
			o.Direction = keycodes.RIGHT
			o.SubDirection = keycodes.UP
		} else if k.AreKeysCombinedInOrder(ebiten.KeyW, ebiten.KeyA) || k.AreKeysCombinedInOrder(ebiten.KeyArrowUp, ebiten.KeyArrowLeft) {
			o.Direction = keycodes.UP
			o.SubDirection = keycodes.LEFT
		} else if k.AreKeysCombinedInOrder(ebiten.KeyA, ebiten.KeyW) || k.AreKeysCombinedInOrder(ebiten.KeyArrowLeft, ebiten.KeyArrowUp) {
			o.Direction = keycodes.LEFT
			o.SubDirection = keycodes.UP
		} else if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
			o.Direction = keycodes.UP
			o.SubDirection = keycodes.NONE
		} else if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
			o.Direction = keycodes.LEFT
			o.SubDirection = keycodes.NONE
		} else if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
			o.Direction = keycodes.RIGHT
			o.SubDirection = keycodes.NONE
		} else if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
			o.Direction = keycodes.DOWN
			o.SubDirection = keycodes.NONE
		}
	}
}

func (o *Object) IsDirectionUP() bool {
	return o.Direction == keycodes.UP || o.SubDirection == keycodes.UP
}
func (o *Object) IsDirectionLEFT() bool {
	return o.Direction == keycodes.LEFT || o.SubDirection == keycodes.LEFT
}
func (o *Object) IsDirectionRIGHT() bool {
	return o.Direction == keycodes.RIGHT || o.SubDirection == keycodes.RIGHT
}
func (o *Object) IsDirectionDOWN() bool {
	return o.Direction == keycodes.DOWN || o.SubDirection == keycodes.DOWN
}

func (o *Object) IsBy(ob Object) bool {
	return o.ParentID == ob.ID
}

func (o *Object) SaveLastPosition() {
	o.PositionHistory.Add(image.Point{X: int(o.RawPos.X), Y: int(o.RawPos.Y)})
}

func (o *Object) savePositionBeforeAnimation() {
	o.Animation.PositionBeforeAnimation = image.Point{X: int(o.RawPos.X), Y: int(o.RawPos.Y)}
}

func (o *Object) SetRawX(x float64) {
	o.RawPos.X = x
	o.savePositionBeforeAnimation()
}

func (o *Object) SetRawY(y float64) {
	o.RawPos.Y = y
	o.savePositionBeforeAnimation()
}

func (o *Object) SetRawPosForCameraY(y float64) {
	o.RawPosForCamera.Y = y
}
func (o *Object) SetRawPosForCameraX(x float64) {
	o.RawPosForCamera.X = x
}

func (o *Object) SetAttachedPosX(x float64) {
	o.AttachedPos.X = x
}

func (o *Object) SetAttachedPosY(y float64) {
	o.AttachedPos.Y = y
}

func (o *Object) SetZoomedAttachedPosX(x float64) {
	w := UseObjects().World()
	mapScaleX, _ := w.GetZoomedMapScale()
	o.AttachedPos.X = x / mapScaleX
}

func (o *Object) SetZoomedAttachedPosY(y float64) {
	w := UseObjects().World()
	_, mapScaleY := w.GetZoomedMapScale()
	o.AttachedPos.Y = y / mapScaleY
}

func (o *Object) GetRawX() float64 {
	return o.RawPos.X
}

func (o *Object) GetRawY() float64 {
	return o.RawPos.Y
}

// func (o *Object) GetZoomedRawX() float64 {
// 	return o.RawPos.X * 
// }

// func (o *Object) GetZoomedRawY() float64 {
// 	return o.RawPos.Y
// }

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

//Sets skin for the object
func (o *Object) SetSkin(path string) {
	o.Path = path
	_, file := filepath.Split(path)
	o.Name = file
}

//Returns images for the skin selected
func (o *Object) GetImage() *ebiten.Image {
	return sources.UseSources().Images().GetImage(o.Path)
}

func (o *Object) GetCopyOfImage() *ebiten.Image {
	return ebiten.NewImageFromImage(sources.UseSources().Images().GetImage(o.Path))
}

//Returns image where animation properties applied to
func (o *Object) GetAnimatedImage() *ebiten.Image {
	i := o.GetImage()
	m := o.GetMetadata().Modified

	sx, sy := int((m.Animation.FrameX+float64(o.Animation.FrameCount))*m.Animation.FrameWidth), int(m.Animation.FrameY)
	return i.SubImage(image.Rect(sx, sy, sx+int(m.Animation.FrameWidth), sy+int(m.Animation.FrameHeight))).(*ebiten.Image)
}

//Returns metadata for the skin selected
func (o *Object) GetMetadata() *sources.ModelCombination {
	return sources.UseSources().Metadata().GetMetadata(o.Path)
}

//API//

func (o *Object) ToAPIMessage() *api.Object {
	return &api.Object{
		Animation: &api.Animation{
			PositionBeforeAnimation: &api.Position{
				X: float64(o.Animation.PositionBeforeAnimation.X),
				Y: float64(o.Animation.PositionBeforeAnimation.Y),
			},
			FrameCount:         uint64(o.FrameCount),
			FrameDelayCounter:  uint64(o.FrameDelayCounter),
			CurrentFrameMatrix: o.CurrentFrameMatrix,
		},
		Skin: &api.Skin{
			Name: o.Name,
			Path: o.Path,
		},
		Physics: &api.Physics{
			Jump: *(*[]int64)(unsafe.Pointer(&o.Jump)),
		},
		WorldId:  o.WorldID.String(),
		ParentId: o.ParentID.String(),
		Id:       o.ID.String(),
		RawPos: &api.Position{
			X: o.RawPos.X,
			Y: o.RawPos.Y,
		},
		Spawn: &api.Position{
			X: float64(o.Spawn.X),
			Y: float64(o.Spawn.Y),
		},
		Direction: int64(o.Direction),
		Role:      int64(o.Role),
	}
}

func (o *Object) FromAPIMessage(m *api.Object) {
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

func (o *Object) GetZoomForSkin(zoom float64) (float64, float64) {
	m := o.GetMetadata().Modified
	return (m.Scale.CoefficiantX / 100 * zoom), (m.Scale.CoefficiantY / 100 * zoom)
}

func (o *Object) GetMaxZoomForSkin(maxZoom float64) (float64, float64) {
	m := o.GetMetadata().Modified
	return (m.Scale.CoefficiantX / 100 * maxZoom), (m.Scale.CoefficiantY / 100 * maxZoom)
}

func (o *Object) GetZoomedRawPos(mapScaleX, mapScaleY float64) (float64, float64) {
	return o.RawPos.X * mapScaleX, o.RawPos.Y * mapScaleY
}

func (o *Object) GetZoomedRawPosForCamera(mapScaleX, mapScaleY float64) (float64, float64) {
	// return o.RawPosForCamera.X * mapScaleX, o.RawPosForCamera.Y * mapScaleY
	return o.RawPosForCamera.X, o.RawPosForCamera.Y 
}

func (o *Object) GetZoomedAttachedPos(mapScaleX, mapScaleY float64) (float64, float64) {
	return o.AttachedPos.X * mapScaleX, o.AttachedPos.Y * mapScaleY
}

func (o *Object) SetTranslationXMovementBlocked(s bool) {
	o.TranslationMovementXBlocked = s
}

func (o *Object) SetTranslationYMovementBlocked(s bool) {
	o.TranslationMovementYBlocked = s
}

func (o *Object) IsTranslationMovementBlocked() bool {
	return o.TranslationMovementXBlocked || o.TranslationMovementYBlocked
}

func NewObject() *Object {
	return new(Object)
}
