package collection

import (
	"crypto/sha256"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sirupsen/logrus"
)

var (
	ImageCollection = make(map[[sha256.Size]byte]*ebiten.Image)
	PathsToHash     = make(map[string][sha256.Size]byte)
)

func GetImage(path string) *ebiten.Image {
	i, ok := ImageCollection[PathsToHash[path]]
	if !ok {
		logrus.Fatal(fmt.Sprintf("image with path '%s' not found", path))
	}
	return i
}
