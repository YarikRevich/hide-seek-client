package audio

import (
	"fmt"

	"github.com/YarikRevich/HideSeek-Client/internal/audio/game"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/audio"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/ui"
)

func Process() {
	if audio.UseStatus().GetState() == audio.DONE {
		switch ui.UseStatus().GetState() {
		case ui.GAME:
			fmt.Println(audio.UseStatus().GetState())
			audio.UseStatus().SetState(audio.UNDONE)
			fmt.Println(audio.UseStatus().GetState())
			game.Exec()
		}
	}
}
