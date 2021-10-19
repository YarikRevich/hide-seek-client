package mousepress

import (
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/models"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func IsMousePressLeftOnce(m models.Metadata) bool {
	currX, currY := ebiten.CursorPosition()
	mx, my := m.FastenMarginsWithCoefficients()

	return inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) &&
	(currX >= int(mx) && currX <= int((m.Size.Width)+(mx))) &&
	(currY >= int(my) && currY <= int((m.Size.Height)+(my)))
}
