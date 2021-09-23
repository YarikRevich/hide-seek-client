package audio

import (
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"

	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
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
	switch statemachine.GetInstance().GetState() {
	case statemachine.GAME:
		play("assets/audio/game.mp3")
	}
}
