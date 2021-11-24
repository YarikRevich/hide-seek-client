package screen

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	"github.com/hajimehoshi/ebiten/v2"
)

var instance *Screen

type Screen struct {
	fullWidth, fullHeight int
	lastWidth, lastHeight int

	screen *ebiten.Image
}

func (s *Screen) SetLastSize() {
	width, height := s.screen.Size()
	s.lastWidth = width
	s.lastHeight = height
}

func (s *Screen) GetLastSize() (float64, float64) {
	if s.lastWidth == 0 || s.lastHeight == 0{
		return s.GetSize()
	}
	return float64(s.lastWidth), float64(s.lastHeight)
}

func (s *Screen) GetMinWidth() int {
	return int((s.GetMaxWidth() * 60) / 100)
}

func (s *Screen) GetMinHeight() int {
	return int((s.GetMaxHeight() * 60) / 100)
}

func (s *Screen) GetMaxWidth() int {
	return int(float64(s.fullWidth) / 1.15)
}

func (s *Screen) GetMaxHeight() int {
	return int(float64(s.fullHeight) / 1.15)
}

func (s *Screen) GetAxisX() float64 {
	if s.screen != nil {
		x, _ := s.screen.Size()
		return float64(x) / 2
	}
	return float64(s.lastWidth) / 2
}

func (s *Screen) GetAxisY() float64 {
	if s.screen != nil {
		_, y := s.screen.Size()
		return float64(y) / 2
	}
	return float64(s.lastHeight) / 2
}

func (s *Screen) SetScreen(i *ebiten.Image) {
	s.screen = i
}

func (s *Screen) CleanScreen() {
	s.screen = nil
}

func (s *Screen) GetWidth() float64 {
	if s.screen != nil {
		return float64(s.screen.Bounds().Max.X)
	}
	return float64(s.lastWidth)
}

func (s *Screen) GetHeight() float64 {
	if s.screen != nil {
		return float64(s.screen.Bounds().Max.Y)
	}
	return float64(s.lastHeight)
}

func (s *Screen) GetScreen() *ebiten.Image {
	return s.screen
}

func (s *Screen) IsResized() bool {
	return s.GetWidth() != float64(s.lastWidth) || s.GetHeight() != float64(s.lastHeight)
}

func (s *Screen) GetSize() (float64, float64) {
	if s.screen != nil {
		width, height := s.screen.Size()
		return float64(width), float64(height)
	}
	return 0, 0
}

func (s *Screen) IsAxisXCrossedByPC() bool {
	p := objects.UseObjects().PC()
	x := p.GetScaledOffsetX()
	ax := s.GetAxisX()

	return (x-p.ModelCombination.Modified.Buffs.Speed.X-p.ModelCombination.Modified.Size.Width/2) <= ax &&
		ax <= (x+p.ModelCombination.Modified.Buffs.Speed.X+p.ModelCombination.Modified.Size.Width/2)
}

func (s *Screen) IsAxisYCrossedByPC() bool {
	p := objects.UseObjects().PC()
	y := p.GetScaledOffsetY()
	ay := s.GetAxisY()

	return (y-p.ModelCombination.Modified.Buffs.Speed.Y-p.ModelCombination.Modified.Size.Height/2) <= ay &&
		ay <= (y+p.ModelCombination.Modified.Buffs.Speed.Y+p.ModelCombination.Modified.Size.Height/2)
}

func UseScreen() *Screen {
	if instance == nil {
		fullWidth, fullHeight := ebiten.ScreenSizeInFullscreen()
		instance = &Screen{
			fullWidth:  fullWidth,
			fullHeight: fullHeight,
		}
	}
	return instance
}
