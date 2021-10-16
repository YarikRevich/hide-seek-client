package mousepress

import (
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/models"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func IsMousePressLeftOnce(m models.Metadata) bool {
	currX, currY := ebiten.CursorPosition()
	return inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) &&
	(currX >= int(m.Margins.LeftMargin*m.Scale.CoefficiantX) && currX <= int((m.Size.Width*m.Scale.CoefficiantX)+(m.Margins.LeftMargin*m.Scale.CoefficiantX))) &&
	(currY >= int(m.Margins.TopMargin*m.Scale.CoefficiantY) && currY <= int((m.Size.Height*m.Scale.CoefficiantY)+(m.Margins.TopMargin*m.Scale.CoefficiantY)))
}
