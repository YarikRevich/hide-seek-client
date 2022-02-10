package screen

import (
	"fmt"

	// "github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
	"github.com/hajimehoshi/ebiten/v2"
)

//All metrics of the ScreenManager are shown
//in tiles
type ScreenManager struct {
	SizeOnStartup types.Vec2
	// Pixels       []types.Vec2
	// OrigTileSize int
	// Scale        int

	// TileSize int

	//Describes full ScreenManager size
	MaxScreenSize, MinScreenSize, lastSize types.Vec2

	Image *ebiten.Image
}

// NEW API

// type RenderTextCharachterOpts struct {
// 	Position                            types.Vec2
// 	FontAdvance, FontDistance, RowWidth float64
// 	Font                                font.Face
// 	Color                               color.Color
// }

//Renders text charachters for ui elements
// func (s *ScreenManager) RenderTextCharachter(cNum int, c rune, opts RenderTextCharachterOpts) {
// var yOffset float64
// if c != '\n' {
// 	yOffset = math.Floor(opts.Position.X*float64(cNum) + opts.FontDistance*float64(cNum-1)/opts.RowWidth)
// } else {
// 	delta := float64(cNum) - opts.RowWidth
// 	cNumInc := opts.RowWidth
// 	if delta > 0 {
// 		cNumInc = math.Ceil(float64(cNum) / cNumInc)
// 	}
// 	yOffset = math.Floor(opts.Position.X*float64(cNumInc) + opts.FontDistance*float64(cNumInc-1)/opts.RowWidth)
// }

// text.Draw(
// 	s.GetImage(),
// 	string(c),
// 	opts.Font,
// 	int(opts.Position.X*float64(cNum)+opts.FontDistance*float64(cNum/(cNum+1))),
// 	int(opts.Position.Y*yOffset+opts.FontAdvance*float64(yOffset-1)),
// 	opts.Color)
// }

// func (s *ScreenManager) RenderTile(position types.Vec2, tile sources.Tile) {
// 	// ebitenutil.DrawRect(screenImage)

// 	tileSize := tile.GetSize()
// yLoop:
// 	for y := 0; y < tileSize.Y; y++ {
// 		ya := y + int(position.Y)
// 		for x := 0; x < tileSize.X; x++ {
// 			xa := x + int(position.X)
// 			screenWidth, screenHeight := s.GetImage().Size()
// 			if xa < 0 || xa >= screenWidth || ya < 0 || ya >= screenHeight {
// 				break yLoop
// 			}
// 			s.Pixels[x+y*screenWidth] = tile.Pixels[x+y*tileSize.X]
// 		}
// 	}
// }

func (s *ScreenManager) SetImage(i *ebiten.Image) {
	s.Image = i
}

func (s *ScreenManager) GetImage() *ebiten.Image {
	return s.Image
}

func (s *ScreenManager) GetScale() types.Vec2 {
	currentSize := s.GetSize()
	return types.Vec2{
		X: s.SizeOnStartup.X / currentSize.X,
		Y: s.SizeOnStartup.Y / currentSize.Y}
}

func (s *ScreenManager) CleanScreen() {
	s.Image = nil
}

func (s *ScreenManager) SetLastSize() {
	s.lastSize = s.GetSize()
}

func (s *ScreenManager) GetLastSize() types.Vec2 {
	if s.lastSize.X == 0 || s.lastSize.Y == 0 {
		return s.GetSize()
	}
	return s.lastSize
}

func (s *ScreenManager) GetSize() types.Vec2 {
	if s.Image != nil {
		width, height := s.Image.Size()
		return types.Vec2{X: float64(width), Y: float64(height)}
	}
	return s.lastSize
}

func (s *ScreenManager) GetMaxSize() types.Vec2 {
	return s.MaxScreenSize
}

func (s *ScreenManager) GetMinSize() types.Vec2 {
	return s.MinScreenSize
}

func (s *ScreenManager) GetAxis() types.Vec2 {
	if s.Image != nil {
		x, y := s.Image.Size()
		return types.Vec2{X: float64(x) / 2, Y: float64(y) / 2}
	}
	return types.Vec2{X: s.lastSize.X / 2, Y: s.lastSize.Y / 2}
}

func (s *ScreenManager) GetAxisSleepingZones() types.Vec2 {
	a := s.GetAxis()
	return types.Vec2{X: (a.X * 5) / 100, Y: (a.Y * 5) / 100}
}

func (s *ScreenManager) IsResized() bool {
	size := s.GetSize()
	return size.X != s.lastSize.X || size.Y != s.lastSize.Y
}

func (s *ScreenManager) GetHUDOffset() float64 {
	return s.GetSize().Y / 12
}

func (s *ScreenManager) IsLessAxisXCrossed(x float64, speedX float64) bool {
	a := s.GetAxis()
	asz := s.GetAxisSleepingZones()

	fmt.Println(x, a.X)
	return x < a.X+(speedX+asz.X) && x > a.X-(speedX+asz.X)
}

func (s *ScreenManager) IsHigherAxisXCrossed(x float64, speedX float64) bool {
	a := s.GetAxis()
	asz := s.GetAxisSleepingZones()

	return x > a.X-(speedX+asz.X) && x < a.X+(speedX+asz.X)
}

func (s *ScreenManager) IsLessAxisYCrossed(y float64, speedY float64) bool {
	a := s.GetAxis()
	asz := s.GetAxisSleepingZones()

	return y < a.Y+(speedY+asz.Y) && y > a.Y-(speedY+asz.Y)
}

func (s *ScreenManager) IsHigherAxisYCrossed(y float64, speedY float64) bool {
	a := s.GetAxis()
	asz := s.GetAxisSleepingZones()

	return y > a.Y-(speedY+asz.Y) && y < a.Y+(speedY+asz.Y)
}

// func (s *ScreenManager) GetOffset() types.Vec2 {
// 	return types.Vec2{X: math.Ceil(s.GetAxisX()), Y: math.Ceil(s.GetAxisY())}
// }

func NewScreenManager() *ScreenManager {
	fullScreenWidth, fullScreenHeight := ebiten.ScreenSizeInFullscreen()
	return &ScreenManager{
		MaxScreenSize: types.Vec2{
			X: float64(fullScreenWidth),
			Y: float64(fullScreenHeight)},
		SizeOnStartup: types.Vec2{
			X: float64(fullScreenWidth),
			Y: float64(fullScreenHeight)},
		MinScreenSize: types.Vec2{
			X: float64(fullScreenWidth * 60 / 100),
			Y: float64(fullScreenHeight * 60 / 100)},
	}
}

// 		fullScreenWidth, fullScreenHeight := ebiten.ScreenSizeInFullscreen()
// 		origTileSize := 17
// 		scale := 1
// 		tileSize := origTileSize * scale
// 		maxScreenSize := types.Vec2{X: float64(fullScreenWidth / tileSize), Y: float64(fullScreenHeight / tileSize)}
// 		minScreenSize := types.Vec2{X: math.Floor(float64(fullScreenWidth * 60 / 100 / tileSize)), Y: math.Floor(float64(fullScreenHeight * 60 / 100 / tileSize))}

// 		instance = &ScreenManager{
// 			OrigTileSize:  origTileSize,
// 			Scale:         scale,
// 			TileSize:      tileSize,
// 			MaxScreenSize: maxScreenSize,
// 			MinScreenSize: minScreenSize,
