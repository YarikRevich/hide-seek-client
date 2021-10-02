package loop

import (
	"github.com/YarikRevich/HideSeek-Client/internal/audio"
	"github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard"
	"github.com/YarikRevich/HideSeek-Client/internal/hid/mouse"
	"github.com/YarikRevich/HideSeek-Client/internal/history"
	"github.com/YarikRevich/HideSeek-Client/internal/networking"
	metadataloader "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/loader/metadata_loader"

	"github.com/YarikRevich/HideSeek-Client/internal/render"
	"github.com/YarikRevich/HideSeek-Client/internal/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

type Loop struct{}

var _ ebiten.Game = (*Loop)(nil)

func (g *Loop) Update() error {
	ui.Process()
	mouse.Process()
	keyboard.Process()
	audio.Process()
	networking.Process()
	return nil
}

func (g *Loop) Draw(screen *ebiten.Image) {
	metadataloader.SyncMetadata(screen)

	for _, dt := range render.GetToRender(){
		dt(screen)
	}
	render.CleanRenderPool()

	w, h := screen.Size()
	history.SetScreenSize(history.ScreenSize{
		Height: h, Width: w})
}

func (g *Loop) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func New() *Loop {
	return new(Loop)
}
