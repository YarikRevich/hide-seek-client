package audiocontroller

import (
	"math"
	"time"

	"github.com/YarikRevich/HideSeek-Client/internal/core/middlewares"
	"github.com/YarikRevich/HideSeek-Client/internal/core/sources"
	"github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"
	"github.com/faiface/beep/speaker"
	"github.com/sirupsen/logrus"
)

var instance *AudioController

type AudioController struct {
	LastTrackPath string
	track         *sources.Track
}

func (a *AudioController) Wrap(path string) {
	a.track = sources.UseSources().Audio().GetAudioController(path)
}

func (a *AudioController) Start() {
	a.LastTrackPath = a.track.TrackPath

	go func() {
		tick := time.NewTicker(time.Microsecond * 500)
		for math.Abs(math.Ceil(a.track.Volume.Volume*100)/100) != 2 {
			select {
			case <-tick.C:
				speaker.Lock()
				a.track.Volume.Volume += 0.001
				speaker.Unlock()
			default:
			}
		}
		tick.Stop()

		a.track.Ctrl.Paused = false

		if err := speaker.Init(a.track.Format.SampleRate, a.track.Format.SampleRate.N(time.Second/100)); err != nil {
			logrus.Fatal("error happened initiating audion in audio loader")
		}

		speaker.Play(a.track.Volume)
		go func() {
			ticker := time.NewTicker(time.Millisecond * 500)
			for range ticker.C {
				if a.track.Streamer.Position() == a.track.Streamer.Len() {
					break
				}
			}
			ticker.Stop()

			a.track.Ctrl.Paused = true

			middlewares.UseMiddlewares().Audio().UseAfter(func() {
				statemachine.UseStateMachine().Audio().SetState(statemachine.AUDIO_DONE)
			})
		}()
	}()
}

func (a *AudioController) Stop() {
	go func() {
		tick := time.NewTicker(time.Microsecond * 500)
		
		for math.Abs(math.Ceil(a.track.Volume.Volume*100)/100) != 0 {
			select {
			case <-tick.C:
				speaker.Lock()
				a.track.Volume.Volume -= 0.001
				speaker.Unlock()
			default:
			}
		}
		tick.Stop()
		
		speaker.Lock()
		a.track.Ctrl.Paused = !a.track.Ctrl.Paused
		speaker.Unlock()
	}()
}

func UseAudioController() *AudioController {
	if instance == nil {
		instance = new(AudioController)
	}
	return instance
}