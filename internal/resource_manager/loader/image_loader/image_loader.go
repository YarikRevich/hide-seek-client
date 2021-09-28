package imageloader

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"regexp"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/sirupsen/logrus"
)

var (
	Images = make(map[[sha256.Size]byte]*ebiten.Image)
	PathsToHash = make(map[string][sha256.Size]byte)
	mu     = sync.Mutex{}
)

func GetImage(path string)*ebiten.Image{
	i, ok := Images[PathsToHash[path]]
	if !ok{
		logrus.Fatal(fmt.Sprintf("image with path '%s' not found", path))
	}
	return i
}

func Load(motherDir, extension, path string, wg *sync.WaitGroup) {
	if extension != "png" {
		return
	}

	wg.Add(1)
	go func() {
		f, err := os.ReadFile(path)
		if err != nil{
			log.Fatalln(err)
		}

		imageHash := sha256.Sum256(f)

		img, _, err := ebitenutil.NewImageFromFile(path)
		if err != nil {
			log.Fatalln(err)
		}

		path = path[len(motherDir):]
		reg := regexp.MustCompile(`\.[a-z]*$`)
		if reg.MatchString(path) {
			mu.Lock()
			defer mu.Unlock()
			Images[imageHash] = img
			PathsToHash[reg.Split(path, -1)[0]] = imageHash
		}

		wg.Done()
	}()

}
