package middlewares

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	"github.com/YarikRevich/HideSeek-Client/internal/core/latency"
	"github.com/YarikRevich/HideSeek-Client/internal/core/sources"
	"github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"
)

type UI struct{}

func (u *UI) cleanLatencyOnce() {
	latency.UseLatency().Once().Reset()
}

func (u *UI) cleanBuffers() {
	events.UseEvents().Input().SettingsMenuNameBuffer.Clean()
}

func (u *UI) setSuspendedMusicDone() {
	// UseMiddlewares().Audio().UseAfter(func() {
		statemachine.UseStateMachine().Audio().SetState(statemachine.AUDIO_DONE)
	// })
}

func (u *UI) cleanUsedMetadata(){
	sources.UseSources().Metadata().CleanUsedMetadata()
}


func (u *UI) UseAfter(c func()) {
	c()

	u.cleanBuffers()
	u.setSuspendedMusicDone()
	u.cleanUsedMetadata()
}

func NewUI() *UI {
	return new(UI)
}
