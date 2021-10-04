package audioloader

import (
	"embed"
	"regexp"
	"sync"
	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/sirupsen/logrus"

	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/audio_loader/collection"
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/audio_loader/controllers"
)

var mu = sync.Mutex{}

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

		ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}
		volume := &effects.Volume{
			Streamer: ctrl,
			Base: 2,
			Volume: 0.001,
		}

		reg := regexp.MustCompile(`\.[a-z0-9]*$`)

		if reg.MatchString(path) {
			mu.Lock()
			collection.AudioControllers[reg.Split(path, -1)[0]] = 
				controllers.NewController(volume, ctrl, format, streamer)
			mu.Unlock()
		}

		wg.Done()
	}()
}
