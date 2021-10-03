package startmenu

import (
	mousepress "github.com/YarikRevich/HideSeek-Client/internal/detectors/mouse_press"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/ui"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/input"
	metadataloader "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/loader/metadata_loader"
)



func Exec(){
	if mousepress.IsMousePressLeftOnce(*metadataloader.GetMetadata("assets/images/menues/buttons/settingswheel")){
		ui.UseStatus().SetState(ui.GAME)
		input.UseStatus().SetState(input.GAME)
	}
}