package imageloader

import (
	"log"
	"regexp"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	Images = make(map[string]*ebiten.Image)
	mu     = sync.Mutex{}
)

func Load(motherDir, extension, path string, wg *sync.WaitGroup) {
	if extension != "png" {
		return
	}

	wg.Add(1)
	go func() {
		img, _, err := ebitenutil.NewImageFromFile(path)
		if err != nil {
			log.Fatalln(err)
		}

		path = path[len(motherDir):]
		reg := regexp.MustCompile(`\.[a-z]*$`)
		if reg.MatchString(path) {
			mu.Lock()
			defer mu.Unlock()
			Images[reg.Split(path, -1)[0]] = img
		}

		wg.Done()
	}()

}
