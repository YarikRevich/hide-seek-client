package middlewares

import (
	// "sync"

	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	// networkingcollection "github.com/YarikRevich/HideSeek-Client/internal/networking/collection"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/audio"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/applyer"
	audiomiddleware "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/audio"
)

type UI struct {}

func (u *UI)cleanBuffers() {
	events.UseEvents().Input().SettingsMenuNameBuffer.Clean()
}

func (u *UI)setSuspendedMusicDone() {
	// applyer.ApplyMiddlewares(
	// 	statemachine.UseStateMachine().Audio().SetState(audio.DONE),
	// 	audiomiddleware.UseAudioMiddleware,
	// )
}

func (u *UI) UseAfter (c func()){
	c()

	u.cleanBuffers()
	u.setSuspendedMusicDone()
}

func NewUI() *UI{
	return new(UI)
}
