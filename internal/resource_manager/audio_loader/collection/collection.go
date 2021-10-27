package collection

import (
	"fmt"

	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/audio_loader/models"
	"github.com/sirupsen/logrus"
)

var (
	lastAudioPath    string
	AudioControllers = make(map[string]models.Controller)
)

func SetLastAudioTrackPath(p string) {
	lastAudioPath = p
}

func GetLastAudioTrackPath() string {
	return lastAudioPath
}

func GetAudioController(path string) models.Controller {
	i, ok := AudioControllers[path]
	if !ok {
		logrus.Fatal(fmt.Sprintf("audio with path '%s' not found", path))
	}
	return i
}
