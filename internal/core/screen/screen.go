package screen

import (
	"fmt"

	"github.com/YarikRevich/hide-seek-client/internal/core/types"
	"github.com/hajimehoshi/ebiten/v2"
)

var instance *Screen

type Screen struct {
	//Describes full screen size
	max, min, lastSize types.Vec2

	screen *ebiten.Image
}

func (s *Screen) SetScreen(i *ebiten.Image) {
	s.screen = i
}

func (s *Screen) GetScreen() *ebiten.Image {
	return s.screen
}

func (s *Screen) CleanScreen() {
	s.screen = nil
}

func (s *Screen) SetLastSize() {
	width, height := s.screen.Size()
	s.lastSize = types.Vec2{X: float64(width), Y: float64(height)}
}

func (s *Screen) GetLastSize() types.Vec2 {
	if s.lastSize.X == 0 || s.lastSize.Y == 0 {
		return s.GetSize()
	}
	return s.lastSize
}

func (s *Screen) GetSize() types.Vec2 {
	if s.screen != nil {
		width, height := s.screen.Size()
		return types.Vec2{X: float64(width), Y: float64(height)}
	}
	return s.lastSize
}

func (s *Screen) GetMaxSize() types.Vec2 {
	return s.max
}

func (s *Screen) GetMinSize() types.Vec2 {
	return s.min
}

func (s *Screen) GetAxis() types.Vec2 {
	if s.screen != nil {
		x, y := s.screen.Size()
		return types.Vec2{X: float64(x) / 2, Y: float64(y) / 2}
	}
	return types.Vec2{X: s.lastSize.X / 2, Y: s.lastSize.Y / 2}
}

func (s *Screen) GetAxisSleepingZones() types.Vec2 {
	a := s.GetAxis()
	return types.Vec2{X: (a.X * 5) / 100, Y: (a.Y * 5) / 100}
}

func (s *Screen) IsResized() bool {
	size := s.GetSize()
	return size.X != s.lastSize.X || size.Y != s.lastSize.Y
}

func (s *Screen) GetHUDOffset() float64 {
	return s.GetSize().Y / 12
}

func (s *Screen) IsLessAxisXCrossed(x float64, speedX float64) bool {
	a := s.GetAxis()
	asz := s.GetAxisSleepingZones()

	return x < a.X+(speedX+asz.X) && x > a.X-(speedX+asz.X)
}

func (s *Screen) IsHigherAxisXCrossed(x float64, speedX float64) bool {
	a := s.GetAxis()
	asz := s.GetAxisSleepingZones()

	fmt.Println(x, a.X)
	return x > a.X-(speedX+asz.X) && x < a.X+(speedX+asz.X)
}

func (s *Screen) IsLessAxisYCrossed(y float64, speedY float64) bool {
	a := s.GetAxis()
	asz := s.GetAxisSleepingZones()

	return y < a.Y+(speedY+asz.Y) && y > a.Y-(speedY+asz.Y)
}

func (s *Screen) IsHigherAxisYCrossed(y float64, speedY float64) bool {
	a := s.GetAxis()
	asz := s.GetAxisSleepingZones()

	return y > a.Y-(speedY+asz.Y) && y < a.Y+(speedY+asz.Y)
}

// func (s *Screen) GetOffset() types.Vec2 {
// 	return types.Vec2{X: math.Ceil(s.GetAxisX()), Y: math.Ceil(s.GetAxisY())}
// }

func UseScreen() *Screen {
	if instance == nil {
		fullWidth, fullHeight := ebiten.ScreenSizeInFullscreen()
		instance = &Screen{
			max: types.Vec2{X: float64(fullWidth) / 1.15, Y: float64(fullHeight) / 1.15},
			min: types.Vec2{X: ((float64(fullWidth) / 1.15) * 60) / 100, Y: ((float64(fullHeight) / 1.15) * 60) / 100},
		}
	}
	return instance
}
