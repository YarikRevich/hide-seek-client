package game

import (
	audioloader "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/loader/audio_loader"
)

func Exec(){
	audioloader.GetAudio("assets/audio/game")()
}