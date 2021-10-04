package ui

import (
	"github.com/YarikRevich/HideSeek-Client/internal/buffers/text"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/audio"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/applyer"
	audiomiddleware "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/audio"
)

func cleanBuffers(){
	text.UseBuffer().Clean()
}

func setSuspendedMusicDone(){
	applyer.ApplyMiddlewares(
		statemachine.UseStateMachine().Audio().SetState(audio.DONE),
		audiomiddleware.UseAudioMiddleware,
	)
}

func UseUIMiddleware(c func()){
	cleanBuffers()
	setSuspendedMusicDone()
	c()
}