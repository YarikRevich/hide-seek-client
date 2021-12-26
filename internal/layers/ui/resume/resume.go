package resume

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/effects/lines"
	"github.com/YarikRevich/hide-seek-client/internal/core/render"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var linesPool = lines.NewLinesPool(1024)

func Draw() {
	render.UseRender().SetToRender(func(i *ebiten.Image) {
		linesPool.ForEach(func(l lines.Line) {
			x0, y0, x1, y1, color := l.Raw()
			ebitenutil.DrawLine(i, x0, y0, x1, y1, color)
		})
	})
}
