package sources

import (
	"embed"
	"fmt"
	"regexp"
	"sync"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/sirupsen/logrus"
)

type Track struct {
	Volume *effects.Volume
	Ctrl *beep.Ctrl
	Format beep.Format
	Streamer beep.StreamSeekCloser
	TrackPath string
}

type Audio struct {
	sync.Mutex

	Collection  map[string]*Track
}

func (a *Audio) loadFile(fs embed.FS, path string) {
	sound, err := fs.Open(path)
	if err != nil {
		logrus.Fatal("error happened opening audio file from embedded fs", err)
	}
	defer sound.Close()

	streamer, format, err := mp3.Decode(sound)
	if err != nil {
		logrus.Fatal("error happened decoding audio file from embedded fs", err)
	}

	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}
	volume := &effects.Volume{
		Streamer: ctrl,
		Base:     2,
		Volume:   0.001,
	}

	reg := regexp.MustCompile(`\.[a-z0-9]*$`)

	if reg.MatchString(path) {
		a.Lock()
		trackPath := reg.Split(path, -1)[0]
		a.Collection[trackPath] = &Track{volume, ctrl, format, streamer, trackPath}
		a.Unlock()
	}
}

func (a *Audio) Load(fs embed.FS, path string) {
	NewParser(fs, path, a.loadFile).Parse()
}
func (a *Audio)GetAudioController(path string) *Track {
	audio, ok := a.Collection[path]
	if !ok {
		logrus.Fatal(fmt.Sprintf("audio with path '%s' not found", path))
	}
	return audio
}


func NewAudio() *Audio{
	return new(Audio)
}