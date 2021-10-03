package history

import "github.com/YarikRevich/HideSeek-Client/internal/direction"

type ScreenSize struct {
	Height int
	Width  int
}

var (
	lastDirection  direction.Direction
	lastScreenSize ScreenSize
	lastAudioPosition int
)

func SetDirection(d direction.Direction) {
	lastDirection = d
}

func GetDirection() direction.Direction {
	return lastDirection
}

func SetScreenSize(s ScreenSize) {
	lastScreenSize = s
}

func GetScreenSize() (int, int) {
	return lastScreenSize.Width, lastScreenSize.Height
}

func SetAudioPosition(p int){
	lastAudioPosition = p
}

func GetAudioPosition() int {
	return lastAudioPosition
}
