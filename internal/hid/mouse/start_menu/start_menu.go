package startmenu

import (
	mousepress "github.com/YarikRevich/HideSeek-Client/internal/detectors/mouse_press"
	statemachine "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
	metadataloader "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/loader/metadata_loader"
)



func Exec(){
	if mousepress.IsMousePressLeftOnce(*metadataloader.GetMetadata("assets/images/menues/buttons/settingswheel")){
		statemachine.GetInstance().SetState(statemachine.GAME)
	}
}