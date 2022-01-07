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
	Streamer  beep.StreamSeekCloser
	TrackPath string
}

type Audio struct {
	sync.Mutex

	Collection map[string]*Track
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
		)
		if regexp.MustCompile(`\.(mp3)*$`).MatchString(path) {
			streamer, format, err = mp3.Decode(sound)
			if err != nil {
				logrus.Fatal("error happened decoding mp3 audio file from embedded fs", err)
			}
		} else if regexp.MustCompile(`\.(waw)*$`).MatchString(path) {
			streamer, format, err = wav.Decode(sound)
			if err != nil {
				logrus.Fatal("error happened decoding wav audio file from embedded fs", err)
			}
		}

		ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}
		volume := &effects.Volume{
			Streamer: ctrl,
			Base:     2,
			Volume:   0.001,
		}

		a.Lock()
		trackPath := reg.Split(path, -1)[0]
		a.Collection[trackPath] = &Track{volume, ctrl, format, streamer, trackPath}
		a.Unlock()
	}
}

func (a *Audio) Load(fs embed.FS, path string, wg *sync.WaitGroup) {
	NewParser(fs, path, a.loadFile).Parse()
	wg.Done()
}

func (a *Audio) GetAudioController(path string) *Track {
	fmt.Println(path)
	path = filepath.Join("dist/audio", path)

	audio, ok := a.Collection[path]
	if !ok {
		logrus.Fatal(fmt.Sprintf("audio with path '%s' not found", path))
	}
	return audio
}

func NewAudio() *Audio {
	return &Audio{Collection: make(map[string]*Track)}
}
