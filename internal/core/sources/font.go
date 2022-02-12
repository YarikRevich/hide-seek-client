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

type RenderTextCharachterOpts struct {
	Tilemap                             *Tilemap
	Text                                string
	SurfacePosition                     types.Vec2
	FontAdvance, FontDistance, RowWidth float64
	TextPosition                        types.Vec2
	Color                               color.Color
}

func (f *Font) Render(sm *screen.ScreenManager, opts RenderTextCharachterOpts) {
	// screenScale := sm.GetScale()
	if opts.RowWidth == 0 {
		logrus.Fatalln("RowWidth should be greather than zero")
	}

	var spaceOffset float64
	var lineNum, breakIndex int
	for i, c := range opts.Text {
		fontHeight := f.Font.Metrics().Height

		fontAdvance, ok := f.Font.GlyphAdvance(c)
		if !ok {
			logrus.Fatalln("can't get advance of the font")
		}

		maxSymbols := int(opts.RowWidth / float64(fontAdvance.Round()))

		var yOffset, xOffset float64

		currentSymbolsInRow := int(fontAdvance.Round() * i)
		maxSymbolsInRow := int((fontAdvance.Round() * maxSymbols))

		if (currentSymbolsInRow != 0 && currentSymbolsInRow%maxSymbolsInRow == 0) || c == '\n' {
			lineNum = (currentSymbolsInRow / maxSymbolsInRow)
			breakIndex = i
			if c == ' ' {
				spaceOffset = float64(fontAdvance.Round())
			} else {
				spaceOffset = 0
			}
		}
		yOffset = opts.SurfacePosition.Y + opts.TextPosition.Y + float64(fontHeight.Round()*lineNum)
		fmt.Println(string(c), currentSymbolsInRow, maxSymbolsInRow, yOffset)

		xOffset = opts.SurfacePosition.X + opts.TextPosition.X + (float64(fontAdvance.Round() * (i - breakIndex))) - spaceOffset

		text.Draw(
			sm.GetImage(),
			string(c),
			f.Font,
			int(xOffset),
			int(yOffset),
			color.Opaque)
	}
}
