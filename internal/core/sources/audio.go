package sources

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/YarikRevich/hide-seek-client/assets"
	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/wav"
)

type Audio struct {
	Name   string
	Volume *effects.Volume
	Ctrl   *beep.Ctrl
	Format beep.Format
	Buffer *beep.Buffer
}

func (a *Audio) load(path string) error {
	sound, err := assets.Assets.Open(path)
	if err != nil {
		return fmt.Errorf("error happened opening audio file from embedded fs: %w", err)
	}
	defer sound.Close()

	var (
		streamer beep.StreamSeekCloser
		format   beep.Format
		loop     beep.Streamer
	)

	if regexp.MustCompile(`\.(mp3)*$`).MatchString(path) {
		streamer, format, err = mp3.Decode(sound)
		if err != nil {
			return fmt.Errorf("error happened decoding mp3 audio file from embedded fs: %w", err)
		}
		loop = beep.Loop(-1, streamer)
	} else if regexp.MustCompile(`\.(wav)*$`).MatchString(path) {
		streamer, format, err = wav.Decode(sound)
		if err != nil {
			return fmt.Errorf("error happened decoding wav audio file from embedded fs: %w", err)
		}
		loop = beep.Loop(-1, streamer)
	}

	a.Ctrl = &beep.Ctrl{Streamer: loop, Paused: false}
	a.Volume = &effects.Volume{
		Streamer: a.Ctrl,
		Base:     2,
		Volume:   0.001,
	}
	a.Format = format

	a.Name = strings.Split(path, ".")[0]

	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	if err := streamer.Close(); err != nil {
		return err
	}
	a.Buffer = buffer

	audioCollection[a.Name] = a

	return nil
}

func (a *Audio) Start() {

}

func (a *Audio) Stop() {

}

// //Checks if there is any unloaded audio
// func (a *Audio) IsAllAudioLoaded() bool {
// 	for _, v := range a.Collection {
// 		if v.Track == nil {
// 			return false
// 		}
// 	}
// 	return true
// }
