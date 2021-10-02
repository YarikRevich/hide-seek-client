package audioloader

import (
	"embed"
	"fmt"
	"regexp"
	"sync"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/sirupsen/logrus"
)

var (
	AudioCollection = make(map[string]func())
	mu    = sync.Mutex{}
)

func GetAudio(path string)func(){
	i, ok := AudioCollection[path]
	if !ok{
		logrus.Fatal(fmt.Sprintf("audio with path '%s' not found", path))
	}
	return i
}

func Load(e embed.FS, extension, path string, wg *sync.WaitGroup) {
	if extension != "mp3" {
		return
	}

	wg.Add(1)
	go func() {
		sound, err := e.Open(path)
		if err != nil {
			logrus.Fatal("error happened opening audio file from embedded fs", err)
		}
		defer sound.Close()

		streamer, format, err := mp3.Decode(sound)
		if err != nil {
			logrus.Fatal("error happened decoding audio file from embedded fs", err)
		}

		reg := regexp.MustCompile(`\.[a-z]*$`)
		if reg.MatchString(path) {
			mu.Lock()
			defer mu.Unlock()
			AudioCollection[reg.Split(path, -1)[0]] = func() {
				speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
				speaker.Play(beep.Seq(streamer, nil))
			}
		}

		wg.Done()
	}()
}
