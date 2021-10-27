package controllers

import (
	"math"
	"time"

	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/audio_loader/collection"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/audio"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/applyer"
	audiomiddleware "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/audio"
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/audio_loader/models"
	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/speaker"
	"github.com/sirupsen/logrus"
)

func StopCallback(effect *effects.Volume, ctrl *beep.Ctrl) func() {
	return func() {
		go func() {
			tick := time.NewTicker(time.Microsecond * 500)

			for math.Abs(math.Ceil(effect.Volume*100)/100) != 0 {
				select {
				case <-tick.C:
					speaker.Lock()
					effect.Volume -= 0.001
					speaker.Unlock()
				default:
				}
			}
			tick.Stop()

			speaker.Lock()
			ctrl.Paused = !ctrl.Paused
			speaker.Unlock()
		}()
	}
}

func StartCallback(
	effect *effects.Volume, ctrl *beep.Ctrl, format beep.Format, streamer beep.StreamSeekCloser, path string) func() {
	return func() {
		collection.SetLastAudioTrackPath(path)

		go func() {
			tick := time.NewTicker(time.Microsecond * 500)

			for math.Abs(math.Ceil(effect.Volume*100)/100) != 2 {
				select {
				case <-tick.C:
					speaker.Lock()
					effect.Volume += 0.001
					speaker.Unlock()
				default:
				}
			}
			tick.Stop()

			ctrl.Paused = false
			if err := speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/100)); err != nil {
				logrus.Fatal("error happened initiating audion in audio loader")
			}
			speaker.Play(effect)
			go func() {
				ticker := time.NewTicker(time.Millisecond * 500)
				for range ticker.C {
					if streamer.Position() == streamer.Len() {
						break
					}
				}
				ticker.Stop()

				ctrl.Paused = true

				applyer.ApplyMiddlewares(
					statemachine.UseStateMachine().Audio().SetState(audio.DONE),
					audiomiddleware.UseAudioMiddleware,
				)

			}()
		}()
	}
}

func NewController(
	effect *effects.Volume, ctrl *beep.Ctrl, format beep.Format, streamer beep.StreamSeekCloser, path string) models.Controller {
	return models.Controller{
		Stop:  StopCallback(effect, ctrl),
		Start: StartCallback(effect, ctrl, format, streamer, path),
	}
}
