package collection

import (
	"fmt"

	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/audio_loader/models"
	"github.com/sirupsen/logrus"
)

var AudioControllers = make(map[string]models.Controller)

func GetAudioController(path string)models.Controller{
	i, ok := AudioControllers[path]
	if !ok{
		logrus.Fatal(fmt.Sprintf("audio with path '%s' not found", path))
	}
	return i
}