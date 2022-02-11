package sources

import (
	"fmt"
	"image/color"
	"math"
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
	screenScale := sm.GetScale()
	if opts.RowWidth == 0 {
		logrus.Fatalln("RowWidth should be greather than zero")
	}

	var lastYOffset float64
	var lastIndex int
	var spaceOffset float64
	for i, c := range opts.Text {
		fontHeight := f.Font.Metrics().Height

		fontAdvance, ok := f.Font.GlyphAdvance(c)
		if !ok {
			logrus.Fatalln("can't get advance of the font")
		}

		var yOffset, xOffset float64

		if c != '\n' {
			yOffset = opts.SurfacePosition.X + float64(fontHeight.Round())*1.5*(math.Floor(((opts.SurfacePosition.X)+(opts.TextPosition.X)+(float64(fontAdvance.Round()*(i)))/screenScale.X)/(opts.RowWidth+float64(fontAdvance.Round())))-1)
		} else {
			delta := float64(i) - opts.RowWidth
			cNumInc := opts.RowWidth
			if delta > 0 {
				cNumInc = math.Ceil(float64(i) / cNumInc)
			}
			yOffset = math.Floor(opts.SurfacePosition.X*float64(cNumInc) + opts.FontDistance*float64(cNumInc-1)/opts.RowWidth)
		}
		if lastYOffset != yOffset {
			lastIndex = i

			if c == ' ' {
				spaceOffset = float64(fontAdvance.Round())
			} else {
				spaceOffset = 0
			}
		}

		xOffset = (opts.SurfacePosition.X) + (opts.TextPosition.X) + (float64(fontAdvance.Round() * (i - lastIndex))) - spaceOffset/screenScale.X
		lastYOffset = yOffset

		text.Draw(
			sm.GetImage(),
			string(c),
			f.Font,
			int(xOffset),
			int(yOffset),
			color.Opaque)
	}
}
