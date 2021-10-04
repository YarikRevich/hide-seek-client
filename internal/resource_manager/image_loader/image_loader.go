package imageloader

import (
	"bytes"
	"crypto/sha256"
	"embed"
	"image"
	_ "image/png"
	"regexp"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sirupsen/logrus"

	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/image_loader/collection"
)

var mu = sync.Mutex{}



func Load(e embed.FS, extension, path string, wg *sync.WaitGroup) {
	if extension != "png" {
		return
	}

	wg.Add(1)
	go func() {
		f, err := e.ReadFile(path)
		if err != nil {
			logrus.Fatal("error happened opening image file from embedded fs", err)
		}

		imageHash := sha256.Sum256(f)

		i, _, err := image.Decode(bytes.NewReader(f))
		if err != nil {
			logrus.Fatal("error happened decoding image file from embedded fs to ebiten image", err)
		}

		img := ebiten.NewImageFromImage(i)

		reg := regexp.MustCompile(`\.[a-z0-9]*$`)
		if reg.MatchString(path) {
			mu.Lock()
			collection.ImageCollection[imageHash] = img
			collection.PathsToHash[reg.Split(path, -1)[0]] = imageHash
			mu.Unlock()
		}

		wg.Done()
	}()

}
