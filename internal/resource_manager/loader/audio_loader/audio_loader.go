package audioloader

import (
	"log"
	"os"
	"regexp"
	"sync"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var (
	Audio = make(map[string]func())
	mu    = sync.Mutex{}
)

func Load(motherDir, extension, path string, wg *sync.WaitGroup) {
	if extension != "mp3" {
		return
	}

	wg.Add(1)
	go func() {
		sound, err := os.Open(path)
		if err != nil {
			log.Fatalln(err)
		}
		defer sound.Close()

		streamer, format, err := mp3.Decode(sound)
		if err != nil {
			log.Fatalln(err)
		}

		path = path[len(motherDir):]
		reg := regexp.MustCompile(`\.[a-z]*$`)
		if reg.MatchString(path) {
			mu.Lock()
			defer mu.Unlock()
			Audio[reg.Split(path, -1)[0]] = func() {
				speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
				speaker.Play(beep.Seq(streamer, nil))
			}
		}

		wg.Done()
	}()
}
