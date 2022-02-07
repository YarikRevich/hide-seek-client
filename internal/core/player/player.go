package player

import (
	"math"
	"time"

	"github.com/YarikRevich/hide-seek-client/internal/core/player/trackmanager"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/faiface/beep/speaker"
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
	// ticker := time.NewTicker(time.Millisecond * 500)
	// for range ticker.C {
	// 	streamer := track.Buffer.Streamer(0, track.Buffer.Len())
	// 	fmt.Println(streamer.Position(), streamer.Len())
	// 	if streamer.Position() == streamer.Len() {
	// 		break
	// 	}
	// }
	// ticker.Stop()

	// fmt.Println(track.Ctrl.Paused, track.TrackPath)
	// track.Ctrl.Paused = true

	// statemachine.UseStateMachine().Audio().SetState(statemachine.AUDIO_DONE)
}

//TODO: create effect of volume lowering when music is stoped
func (a *Player) Play(trackPath string, opts PlayerOpts) {
	go func() {
		track := sources.UseSources().Audio().GetAudioController(trackPath)
		track.Lock()

		if opts.Fading {
			a.silentlyStopCurrentTrack()
		}

		streamer := track.Track.Buffer.Streamer(0, track.Track.Buffer.Len())

		speaker.Play(streamer)

		a.trackManager.Push(track.Track)

		// if !opts.Infinite {
		// 	a.waitTrackEnds(track)
		// }
		track.Unlock()
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
