package ui

import (
	"sync"

	buffercollection "github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard/buffers/collection"
	networkingcollection "github.com/YarikRevich/HideSeek-Client/internal/networking/collection"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/audio"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/applyer"
	audiomiddleware "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/audio"
)

func updateNetworkCollection(){
	for k := range networkingcollection.OnceCollection{
		networkingcollection.OnceCollection[k] = new(sync.Once)
	}
}

func cleanBuffers() {
	buffercollection.SettingsMenuNameBuffer.Clean()
}

func setSuspendedMusicDone() {
	applyer.ApplyMiddlewares(
		statemachine.UseStateMachine().Audio().SetState(audio.DONE),
		audiomiddleware.UseAudioMiddleware,
	)
}

func UseUIMiddleware(c func()) {
	updateNetworkCollection()
	cleanBuffers()
	setSuspendedMusicDone()
	c()
}
