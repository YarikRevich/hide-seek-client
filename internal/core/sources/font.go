package sources

import (
	"fmt"
	"image/color"
	"math"
	"strconv"
	"strings"

	"github.com/YarikRevich/hide-seek-client/assets"
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Font struct {
	Name string
	Font font.Face
}

func (f *Font) load(path string) error {
	file, err := assets.Assets.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error happened opening font file from embedded fs: %s", err)
	}

	parsedFont, err := truetype.Parse(file)
	if err != nil {
		return fmt.Errorf("error happened parsing font file from embedded fs: %s", err)
	}

	name := strings.Split(path, ".")[0]

	size, err := strconv.Atoi(strings.Split(path, "_")[1])
	if err != nil {
		return err
	}

	f.Name = name
	f.Font = truetype.NewFace(parsedFont, &truetype.Options{
		Size:    float64(size),
		Hinting: font.HintingFull,
	})

	fontCollection[f.Name] = f
	return nil
}

type RenderTextCharachterOpts struct {
	Text                                string
	Position                            types.Vec2
	FontAdvance, FontDistance, RowWidth float64
	Color                               color.Color
}

func (f *Font) Render(sm screen.ScreenManager, opts RenderTextCharachterOpts) {
	for i, c := range opts.Text {
		var yOffset float64
		if c != '\n' {
			yOffset = math.Floor(opts.Position.X*float64(i) + opts.FontDistance*float64(i-1)/opts.RowWidth)
		} else {
			delta := float64(i) - opts.RowWidth
			cNumInc := opts.RowWidth
			if delta > 0 {
				cNumInc = math.Ceil(float64(i) / cNumInc)
			}
			yOffset = math.Floor(opts.Position.X*float64(cNumInc) + opts.FontDistance*float64(cNumInc-1)/opts.RowWidth)
		}

		text.Draw(
			sm.GetImage(),
			string(c),
			f.Font,
			int(opts.Position.X*float64(i)+opts.FontDistance*float64(i/(i+1))),
			int(opts.Position.Y*yOffset+opts.FontAdvance*float64(yOffset-1)),
			opts.Color)
	}
}
