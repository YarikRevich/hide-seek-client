package sources

import (
	"embed"
	"fmt"
	"path/filepath"
	"regexp"
	"sync"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/wav"
	"github.com/sirupsen/logrus"
)

type Track struct {
	Volume    *effects.Volume
	Ctrl      *beep.Ctrl
	Format    beep.Format
	Buffer    *beep.Buffer
	TrackPath string
}

type Cell struct {
	sync.Mutex
	Track *Track
}

type Audio struct {
	sync.Mutex

	Collection map[string]*Cell
}

func (a *Audio) loadFile(fs embed.FS, path string) {
	if reg := regexp.MustCompile(`\.(mp3|wav)*$`); reg.MatchString(path) {
		sound, err := fs.Open(path)
		if err != nil {
			logrus.Fatal("error happened opening audio file from embedded fs", err)
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
				logrus.Fatal("error happened decoding mp3 audio file from embedded fs", err)
			}
			loop = beep.Loop(-1, streamer)
		} else if regexp.MustCompile(`\.(wav)*$`).MatchString(path) {
			streamer, format, err = wav.Decode(sound)
			if err != nil {
				logrus.Fatal("error happened decoding wav audio file from embedded fs", err)
			}
			loop = beep.Loop(1, streamer)
		}

		ctrl := &beep.Ctrl{Streamer: loop, Paused: false}
		volume := &effects.Volume{
			Streamer: ctrl,
			Base:     2,
			Volume:   0.001,
		}

		trackPath := reg.Split(path, -1)[0]
		a.Lock()
		a.Collection[trackPath] = &Cell{}
		a.Unlock()

		a.Collection[trackPath].Lock()
		go func() {
			buffer := beep.NewBuffer(format)
			buffer.Append(streamer)
			if err := streamer.Close(); err != nil {
				logrus.Fatal(err)
			}

			c := a.Collection[trackPath]
			c.Track = &Track{volume, ctrl, format, buffer, trackPath}
			a.Collection[trackPath].Unlock()
		}()

	}
}

//Checks if there is any unloaded audio
func (a *Audio) IsAllAudioLoaded() bool {
	for _, v := range a.Collection {
		if v.Track == nil {
			return false
		}
	}
	return true
}

func (a *Audio) Load(fs embed.FS, path string, wg *sync.WaitGroup) {
	NewParser(fs, path, a.loadFile).Parse()
	wg.Done()
}

func (a *Audio) GetAudioController(path string) *Cell {
	path = filepath.Join("dist/audio", path)
	audio, ok := a.Collection[path]
	fmt.Println(ok)
	if !ok {
		logrus.Fatal(fmt.Sprintf("audio with path '%s' not found", path))
	}
	return audio
}

func NewAudio() *Audio {
	return &Audio{Collection: make(map[string]*Cell)}
}
