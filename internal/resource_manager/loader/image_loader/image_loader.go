package imageloader

import (
	"log"
	"os"
	"regexp"
	"crypto/sha256"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	Images = make(map[string]*Image)
	mu     = sync.Mutex{}
)

type Image struct {
	Id [sha256.Size]byte
	Image *ebiten.Image
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

		img, _, err := ebitenutil.NewImageFromFile(path)
		if err != nil {
			log.Fatalln(err)
		}

		path = path[len(motherDir):]
		reg := regexp.MustCompile(`\.[a-z]*$`)
		if reg.MatchString(path) {
			mu.Lock()
			defer mu.Unlock()
			Images[reg.Split(path, -1)[0]] = &Image{
				Id: sha256.Sum256(f),
				Image: img,
			}
		}

		wg.Done()
	}()

}
