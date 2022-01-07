package player

import (
	"math"
	"time"

	"github.com/YarikRevich/hide-seek-client/internal/core/player/trackmanager"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"
	"github.com/faiface/beep/speaker"
	"github.com/sirupsen/logrus"
)

var instance *Player

type PlayerOpts struct {
	Infinite, Fading bool
}

type Player struct {
	trackManager *trackmanager.TrackManager
}

func (a *Player) silentlyStopCurrentTrack() {
	if track := a.trackManager.TopCurrentTrack(); track != nil {
		tick := time.NewTicker(time.Microsecond * 500)
		for math.Abs(math.Ceil(track.Volume.Volume*100)/100) != 2 {
			select {
			case <-tick.C:
				speaker.Lock()
				track.Volume.Volume += 0.001
				speaker.Unlock()
			default:
			}
		}
		tick.Stop()
	}
}

func (a *Player) waitTrackEnds(track *sources.Track) {
	ticker := time.NewTicker(time.Millisecond * 500)
	for range ticker.C {
		if track.Streamer.Position() == track.Streamer.Len() {
			break
		}
	}
	ticker.Stop()

	track.Ctrl.Paused = true

	statemachine.UseStateMachine().Audio().SetState(statemachine.AUDIO_DONE)
}

func (a *Player) Play(trackPath string, opts PlayerOpts) {
	go func() {
		if track := a.trackManager.Find(trackPath); track != nil {
			track.Ctrl.Paused = false
			return
		}

		track := sources.UseSources().Audio().GetAudioController(trackPath)

		if opts.Fading {
			a.silentlyStopCurrentTrack()
		}

		if err := speaker.Init(track.Format.SampleRate, track.Format.SampleRate.N(time.Second/100)); err != nil {
			logrus.Fatal("error happened initializing audio speaker")
		}
		speaker.Play(track.Volume)

		a.trackManager.Push(track)

		if !opts.Infinite {
			a.waitTrackEnds(track)
		}
	}()
}
func (a *Player) Pause(trackPath string) {
	if track := a.trackManager.Find(trackPath); track != nil {
		track.Ctrl.Paused = true
	}
}

//Stops playing of current track(if track is set)
func (a *Player) Stop(trackPath string) {
	go func() {
		if track := a.trackManager.Find(trackPath); track != nil {
			tick := time.NewTicker(time.Microsecond * 500)
			for math.Abs(math.Ceil(track.Volume.Volume*100)/100) != 0 {
				select {
				case <-tick.C:
					speaker.Lock()
					track.Volume.Volume -= 0.001
					speaker.Unlock()
				default:
				}
			}
			tick.Stop()

			speaker.Lock()
			track.Ctrl.Paused = !track.Ctrl.Paused
			speaker.Unlock()

			a.trackManager.Remove(track)
		}
	}()
}

//Stops all the played tracks
func (p *Player) StopAll() {
	p.trackManager.RemoveAll()
}

func UsePlayer() *Player {
	if instance == nil {
		instance = &Player{trackManager: trackmanager.New()}
	}
	return instance
}
