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
	AnimationStartPosition struct {
		X, Y float64
	}
	FrameCount         uint64
	FrameDelayCounter  uint64
	CurrentFrameMatrix []float64
}

type Skin struct {
	Name string
	Path string
}

/*The object structure which describes
each object on the map
*/
type Base struct {
	*sources.ModelCombination

	Animation
	Skin
	Physics

	Parent *Base

	// Names parentid the object referes to
	ID uuid.UUID

	RawPos, RawOffset, LastRawPos struct {
		X, Y float64
	}

	Direction, SubDirection keycodes.Direction

	Spawn image.Point

	Role Role

	TranslationMovementXBlocked, TranslationMovementYBlocked bool

	//Only client fields

	PositionHistorySequence zeroshifter.IZeroShifter
}

func (o *Base) IsEqualTo(ob Base) bool {
	return o.ID == ob.ID
}

func (o *Base) UpdateDirection() {
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

func (o *Base) IsDirectionUP() bool {
	return o.Direction == keycodes.UP || o.SubDirection == keycodes.UP
}
func (o *Base) IsDirectionLEFT() bool {
	return o.Direction == keycodes.LEFT || o.SubDirection == keycodes.LEFT
}
func (o *Base) IsDirectionRIGHT() bool {
	return o.Direction == keycodes.RIGHT || o.SubDirection == keycodes.RIGHT
}
func (o *Base) IsDirectionDOWN() bool {
	return o.Direction == keycodes.DOWN || o.SubDirection == keycodes.DOWN
}

func (o *Base) SaveLastPosition() {
	o.LastRawPos = struct {
		X float64
		Y float64
	}{
		X: o.RawPos.X, Y: o.RawPos.Y}
	o.PositionHistorySequence.Add(image.Point{X: int(o.RawPos.X), Y: int(o.RawPos.Y)})
}

func (o *Base) SaveAnimationStartPosition() {
	o.Animation.AnimationStartPosition = struct {
		X float64
		Y float64
	}{o.RawPos.X, o.RawPos.Y}
}

func (o *Base) SetRawX(x float64) {
	o.RawPos.X = x
}

func (o *Base) SetRawY(y float64) {
	o.RawPos.Y = y
}

func (o *Base) SetRawOffsetY(y float64) {
	o.RawOffset.Y = y
}
func (o *Base) SetRawOffsetX(x float64) {
	o.RawOffset.X = x
}

//Checks if x pos has been changed
//in comparison to the last x poses
func (o *Base) IsXChanged() bool {
	var prevX int
	for _, v := range o.PositionHistorySequence.Get() {
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
func (o *Base) IsYChanged() bool {
	var prevY int
	for _, v := range o.PositionHistorySequence.Get() {
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
func (o *Base) SetSpawn(spawns []image.Point) {
	o.RawPos.X = 500
	o.RawPos.Y = 500
}

//Checks if pc executes animation
func (o *Base) IsAnimatied() bool {
	return len(o.Physics.Jump) != 0
}

//Returns last saved position before animation was executed
func (o *Base) GetAnimationStartPosition() (float64, float64) {
	return o.Animation.AnimationStartPosition.X, o.Animation.AnimationStartPosition.Y
}

//Sets skin for the object
func (o *Base) SetSkin(path string) {
	o.Path = path
	_, file := filepath.Split(path)
	o.Name = file

	o.ModelCombination = sources.UseSources().Metadata().GetMetadata(o.Path)
	// fmt.Println(o.ModelCombination.Modified.Animation, "KAKA")
}

//Returns images for the skin selected
func (o *Base) GetImage() *ebiten.Image {
	return sources.UseSources().Images().GetImage(o.Path)
}

func (o *Base) GetCopyOfImage() *ebiten.Image {
	return ebiten.NewImageFromImage(sources.UseSources().Images().GetImage(o.Path))
}

//Returns image where animation properties applied to
func (o *Base) GetAnimatedImage() *ebiten.Image {
	i := o.GetImage()
	sx, sy := int((o.ModelCombination.Modified.Animation.FrameX+float64(o.Animation.FrameCount))*o.ModelCombination.Modified.Animation.FrameWidth), int(o.ModelCombination.Modified.Animation.FrameY)
	return i.SubImage(image.Rect(sx, sy, sx+int(o.ModelCombination.Modified.Animation.FrameWidth), sy+int(o.ModelCombination.Modified.Animation.FrameHeight))).(*ebiten.Image)
}

//API//

func (o *Base) ToAPIMessage() *api.Base {
	m := &api.Base{
		Animation: &api.Animation{
			PositionBeforeAnimation: &api.Position{
				X: o.Animation.AnimationStartPosition.X,
				Y: o.Animation.AnimationStartPosition.Y,
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
		Role:      int64(o.Role),
	}
	if o.Parent != nil {
		m.Parent = o.Parent.ToAPIMessage()
	}
	return m
}

func (o *Base) FromAPIMessage(m *api.Base) {
	o.Animation.AnimationStartPosition.X = m.Animation.PositionBeforeAnimation.X
	o.Animation.AnimationStartPosition.Y = m.Animation.PositionBeforeAnimation.Y
	o.Animation.FrameCount = m.Animation.FrameCount
	o.Animation.FrameDelayCounter = m.Animation.FrameDelayCounter
	o.Animation.CurrentFrameMatrix = m.Animation.CurrentFrameMatrix

	o.Skin.Name = m.Skin.Name
	o.Skin.Path = m.Skin.Path

	o.Physics.Jump = *(*[]keycodes.Direction)(unsafe.Pointer(&m.Physics.Jump))

	if o.Parent != nil {
		o.Parent.FromAPIMessage(m.Parent)
		o.ID = uuid.MustParse(m.Id)
	}

	o.RawPos.X = m.RawPos.X
	o.RawPos.Y = m.RawPos.Y
	o.Spawn.X = int(m.Spawn.X)
	o.Spawn.Y = int(m.Spawn.Y)
	o.Direction = keycodes.Direction(m.Direction)
	o.Role = Role(m.Role)
}

func (o *Base) GetScaledPosX() float64 {
	return (o.RawPos.X * o.Parent.Modified.RuntimeDefined.ZoomedScale.X) - o.Modified.Offset.X
}

func (o *Base) GetScaledPosY() float64 {
	return (o.RawPos.Y * o.Parent.Modified.RuntimeDefined.ZoomedScale.Y) - o.Modified.Offset.Y
}

func (o *Base) GetScaledOffsetX() float64 {
	return (o.RawOffset.X * o.Parent.Modified.RuntimeDefined.ZoomedScale.X) - o.Modified.Offset.X
}

func (o *Base) GetScaledOffsetY() float64 {
	// fmt.Println(o.Parent.Modified.RuntimeDefined.ZoomedScale.Y)
	return (o.RawOffset.Y * o.Parent.Modified.RuntimeDefined.ZoomedScale.Y) - o.Modified.Offset.Y
}

func (o *Base) SetTranslationXMovementBlocked(s bool) {
	o.TranslationMovementXBlocked = s
}

func (o *Base) SetTranslationYMovementBlocked(s bool) {
	o.TranslationMovementYBlocked = s
}

func (o *Base) IsTranslationMovementBlocked() bool {
	return o.TranslationMovementXBlocked || o.TranslationMovementYBlocked
}

func NewBase() *Base {
	return &Base{Parent: new(Base)}
}
