package mousepress

import (
	metadataloader "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/loader/metadata_loader"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// var (
// 	lastPressedButton ebiten.MouseButton
// )

func IsMousePressLeftOnce(m metadataloader.Metadata) bool {
	currX, currY := ebiten.CursorPosition()
	return inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) &&
		(currX >= int(m.Margins.LeftMargin) && currX <= int(m.Margins.LeftMargin+m.Size.Width)) &&
		(currY >= int(m.Margins.TopMargin) && currY <= int(m.Margins.TopMargin+m.Size.Height))
}
