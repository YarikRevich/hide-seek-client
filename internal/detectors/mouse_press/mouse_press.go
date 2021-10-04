package mousepress

import (
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/models"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func IsMousePressLeftOnce(m models.Metadata) bool {
	currX, currY := ebiten.CursorPosition()
	return inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) &&
		(currX >= int(m.Margins.LeftMargin) && currX <= int(m.Margins.LeftMargin+m.Size.Width)) &&
		(currY >= int(m.Margins.TopMargin) && currY <= int(m.Margins.TopMargin+m.Size.Height))
}
