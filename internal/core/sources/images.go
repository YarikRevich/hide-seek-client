package sources

import (
	"bytes"
	"embed"
	"fmt"
	"image"
	_ "image/png"
	"path/filepath"
	"regexp"
	"sync"

	"github.com/YarikRevich/HideSeek-Client/internal/core/runtime"
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
		fmt.Println(err)
		logrus.Fatal("error happened decoding image file from embedded fs to ebiten image", err)
	}

	img := ebiten.NewImageFromImage(image)

	reg := regexp.MustCompile(`\.[a-z0-9]*$`)
	if reg.MatchString(path) {
		i.Lock()
		i.Collection[reg.Split(path, -1)[0]] = img
		i.Unlock()
	}
}

func (i *Images) Load(fs embed.FS, path string, wg *sync.WaitGroup) {
	NewParser(fs, path, i.loadFile).Parse()
	wg.Done()
}

func (i *Images) GetImage(path string) *ebiten.Image {
	path = filepath.Join("assets/images", path)

	image, ok := i.Collection[path]
	if !ok {
		logrus.Fatal(fmt.Sprintf("image with path '%s' not found", path))
	}

	if runtime.UseRuntime().IsPrepared() {
		return ebiten.NewImageFromImage(image)
	}
	return image
}

func NewImages() *Images {
	return &Images{Collection: make(map[string]*ebiten.Image)}
}
