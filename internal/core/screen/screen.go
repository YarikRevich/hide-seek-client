package screen

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var fullWidth, fullHeight = ebiten.ScreenSizeInFullscreen()
var lastScreenSizeWidth, lastScreenSizeHeight float64 

func GetMinWidth() int {
	return int((GetMaxWidth() * 60) / 100)
}

func GetMinHeight() int {
	return int((GetMaxHeight() * 60) / 100)
}

func GetMaxWidth() int {
	return int(float64(fullWidth) / 1.15)
}

func GetMaxHeight() int {
	return int(float64(fullHeight) / 1.15)
}

func SetLastScreenSize(screen *ebiten.Image){
	w, h := screen.Size()
	lastScreenSizeWidth = float64(w)
	lastScreenSizeHeight = float64(h)
}

func GetLastScreenSize()(float64, float64){
	return lastScreenSizeWidth, lastScreenSizeHeight
}
