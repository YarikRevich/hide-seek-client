package audioloader

import (
	"embed"
	"fmt"
	"regexp"
	"sync"
	"time"

	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/audio"
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

		ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}
		
		reg := regexp.MustCompile(`\.[a-z0-9]*$`)

		if reg.MatchString(path) {
			mu.Lock()
			defer mu.Unlock()
			AudioCollection[reg.Split(path, -1)[0]] = func() {
				ctrl.Paused = false
				if err := speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/100)); err != nil{
					logrus.Fatal("error happened initiating audion in audio loader")
				}
				speaker.Play(ctrl)
				go func(){
					for streamer.Position() != streamer.Len(){}
					ctrl.Paused = true
					audio.UseStatus().SetState(audio.DONE)
				}()	
			}
		}

		wg.Done()
	}()
}
