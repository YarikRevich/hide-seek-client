package middlewares

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/events"
	"github.com/YarikRevich/hide-seek-client/internal/core/latency"

	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"
)

type UI struct {
	blocked bool
}

func (u *UI) cleanTimings() {
	latency.UseLatency().Timings().CleanEachTimings(statemachine.UseStateMachine().UI().GetState())
}

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

func (u *UI) UseAfter(c func()) {
	if !u.blocked {
		u.blocked = true
		u.cleanTimings()

		c()

		u.cleanBuffers()
		u.setSuspendedMusicDone()
		u.cleanLatencyOnce()
		u.blocked = false
	}
}

func NewUI() *UI {
	return new(UI)
}
