package middlewares

import (
	"fmt"

	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	"github.com/YarikRevich/HideSeek-Client/internal/core/latency"

	// "github.com/YarikRevich/HideSeek-Client/internal/core/networking"
	"github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"
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

		fmt.Println("AFTER CLEANING TIMINGS")
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
