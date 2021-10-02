package imageloader

import (
	"bytes"
	"crypto/sha256"
	"embed"
	"fmt"
	"image"
	_ "image/png"
	"regexp"
	"sync"


	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sirupsen/logrus"
)

var (
	ImageCollection = make(map[[sha256.Size]byte]*ebiten.Image)
	PathsToHash = make(map[string][sha256.Size]byte)
	mu     = sync.Mutex{}
)

func GetImage(path string)*ebiten.Image{
	i, ok := ImageCollection[PathsToHash[path]]
	if !ok{
		logrus.Fatal(fmt.Sprintf("image with path '%s' not found", path))
	}
	return i
}

func Load(e embed.FS, extension, path string, wg *sync.WaitGroup) {
	if extension != "png" {
		return
	}

	wg.Add(1)
	go func() {
		f, err := e.ReadFile(path)
		if err != nil{
			logrus.Fatal("error happened opening image file from embedded fs", err)
		}

		imageHash := sha256.Sum256(f)

		i, _, err := image.Decode(bytes.NewReader(f))
		if err != nil{
			logrus.Fatal("error happened decoding image file from embedded fs to ebiten image", err)
		}

		
		img := ebiten.NewImageFromImage(i)
		// if err != nil {
		// 	logrus.Fatal("error happened converting image file from embedded fs to ebiten image", err)
		// }

		reg := regexp.MustCompile(`\.[a-z]*$`)
		if reg.MatchString(path) {
			mu.Lock()
			defer mu.Unlock()
			ImageCollection[imageHash] = img
			PathsToHash[reg.Split(path, -1)[0]] = imageHash
		}

		wg.Done()
	}()

}
