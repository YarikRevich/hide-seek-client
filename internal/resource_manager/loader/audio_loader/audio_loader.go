package audioloader

import (
	"fmt"
	"log"
	"os"
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
			AudioCollection[reg.Split(path, -1)[0]] = func() {
				speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
				speaker.Play(beep.Seq(streamer, nil))
			}
		}

		wg.Done()
	}()
}
