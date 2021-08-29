package sound

import (
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"

	"github.com/YarikRevich/Hide-Seek-with-Guns/internal/status"
)

func play(name string) {
	sound, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer sound.Close()
	streamer, format, err := mp3.Decode(sound)
	if err != nil {
		panic(err)
	}
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
}

func Process() {
	switch status.GetInstance().GetState() {
	case status.GAME:
		play("assets/audio/game.mp3")
	}
}
