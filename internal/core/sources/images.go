package sources

import (
	"bytes"
	"embed"
	"fmt"
	"image"
	"regexp"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sirupsen/logrus"
)

type Images struct {
	sync.Mutex

	Collection map[string]*ebiten.Image
}

func (i *Images) loadFile(fs embed.FS, path string) {
	file, err := fs.ReadFile(path)
	if err != nil {
		logrus.Fatal("error happened opening image file from embedded fs", err)
	}

	image, _, err := image.Decode(bytes.NewReader(file))
	if err != nil {
		logrus.Fatal("error happened decoding image file from embedded fs to ebiten image", err)
	}

	img := ebiten.NewImageFromImage(image)

	reg := regexp.MustCompile(`\.[a-z0-9]*$`)
	if reg.MatchString(path) {
		i.Lock()
		i.Collection[path] = img
		i.Unlock()
	}
}

func (i *Images) Load(fs embed.FS, path string) {
	NewParser(fs, path, i.loadFile).Parse()
}

func (i *Images) GetImage(path string) *ebiten.Image {
	image, ok := i.Collection[path]
	if !ok {
		logrus.Fatal(fmt.Sprintf("image with path '%s' not found", path))
	}
	
	return ebiten.NewImageFromImage(image)
}

func NewImages() *Images {
	return new(Images)
}
