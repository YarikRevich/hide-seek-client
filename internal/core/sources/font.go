package sources

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/YarikRevich/hide-seek-client/assets"
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/sirupsen/logrus"
	"golang.org/x/image/font"
)

type Font struct {
	Name string
	Font font.Face
}

func (f *Font) load(path string, size int) error {
	file, err := assets.Assets.ReadFile(fmt.Sprintf("%s.%s", path, "ttf"))
	if err != nil {
		return fmt.Errorf("error happened opening font file from embedded fs: %s", err)
	}

	parsedFont, err := truetype.Parse(file)
	if err != nil {
		return fmt.Errorf("error happened parsing font file from embedded fs: %s", err)
	}

	name := strings.Split(path, ".")[0]

	f.Name = name
	f.Font = truetype.NewFace(parsedFont, &truetype.Options{
		Size:    float64(size),
		Hinting: font.HintingFull,
	})

	fontCollection[f.Name] = f
	return nil
}

type TextFormatterOpts struct {
	Text     string
	RowWidth float64
}

func (f *Font) formatTextToRender(opts TextFormatterOpts) [][]rune {
	fontAdvance, ok := f.Font.GlyphAdvance(rune(opts.Text[0]))
	if !ok {
		logrus.Fatalln("can't get advance of the font")
	}

	maxSymbolsPerRow := int(opts.RowWidth / float64(fontAdvance.Round()))

	var r [][]rune
	var q []rune

	var breakIndex int
	var spaceShift int
	for i, c := range opts.Text {
		currentSymbolsInRow := int(fontAdvance.Round() * i)
		maxRowWidth := int((fontAdvance.Round() * maxSymbolsPerRow))

		if currentSymbolsInRow != 0 && currentSymbolsInRow%maxRowWidth == spaceShift && breakIndex != i-1 {
			r = append(r, q)
			q = make([]rune, 0)
			breakIndex = i
		}

		if !(c == ' ' && breakIndex != 0 && breakIndex == i-1) {
			q = append(q, c)
		} else {
			spaceShift = fontAdvance.Round()
		}

	}

	r = append(r, q)
	return r
}

type Align string

const (
	None   Align = ""
	Center       = "center"
	Left         = "left"
	Right        = "right"
)

type RenderTextCharachterOpts struct {
	Align

	Tilemap                       *Tilemap
	SurfacePosition, SurfaceScale types.Vec2

	Text         string
	RowWidth     float64
	TextPosition types.Vec2
	Color        color.Color
}

func (f *Font) Render(sm *screen.ScreenManager, opts RenderTextCharachterOpts) {
	if len(opts.Text) == 0 {
		return
	}

	if opts.RowWidth == 0 && opts.Align == None {
		logrus.Fatalln("RowWidth should be greather than zero or Align should be set")
	}

	fontAdvance, ok := f.Font.GlyphAdvance(rune(opts.Text[0]))
	if !ok {
		logrus.Fatalln("can't get advance of the font")
	}

	screenScale := sm.GetScale()
	fontHeight := float64(f.Font.Metrics().Height.Round()) / (screenScale.Y)

	var textPosition types.Vec2
	var textRowWidth float64

	switch opts.Align {
	case Center:
		textPosition.X = (-opts.Tilemap.MapSize.X * opts.SurfaceScale.X / 2) + opts.Tilemap.TileSize.X
		textPosition.Y = -opts.Tilemap.MapSize.Y*opts.SurfaceScale.Y/2 + fontHeight
		textRowWidth = opts.Tilemap.MapSize.X*opts.SurfaceScale.X - (opts.Tilemap.TileSize.X * 2)
	case Right:
		textPosition.X = (-opts.Tilemap.MapSize.X * opts.SurfaceScale.X / 2) + (opts.Tilemap.TileSize.X * 2)
		textPosition.Y = -opts.Tilemap.MapSize.Y*opts.SurfaceScale.Y/2 + fontHeight
		textRowWidth = opts.Tilemap.MapSize.X*opts.SurfaceScale.X - (opts.Tilemap.TileSize.X * 3)
	case Left:
		textPosition.X = (-opts.Tilemap.MapSize.X * opts.SurfaceScale.X / 2) + (opts.Tilemap.TileSize.X)
		textPosition.Y = -opts.Tilemap.MapSize.Y*opts.SurfaceScale.Y/2 + fontHeight
		textRowWidth = opts.Tilemap.MapSize.X*opts.SurfaceScale.X - (opts.Tilemap.TileSize.X * 4)
	default:
		textRowWidth = opts.RowWidth
		textPosition = opts.TextPosition
	}

	formattedText := f.formatTextToRender(TextFormatterOpts{
		Text:     opts.Text,
		RowWidth: textRowWidth,
	})
	for y, p := range formattedText {
		for x, c := range p {
			yOffset := opts.SurfacePosition.Y + textPosition.Y + float64(int(fontHeight)*(y+1))
			xOffset := opts.SurfacePosition.X + textPosition.X + ((float64(fontAdvance.Round()) / screenScale.X) * float64(x))

			text.Draw(
				sm.GetImage(),
				string(c),
				f.Font,
				int(xOffset),
				int(yOffset),
				color.Opaque)
		}
	}
}
