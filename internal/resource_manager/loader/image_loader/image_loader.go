package imageloader

import (
	"crypto/sha256"
	"log"
	"os"
	"regexp"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	Images = make(map[[sha256.Size]byte]*ebiten.Image)
	PathsToHash = make(map[string][sha256.Size]byte)
	mu     = sync.Mutex{}
)

func GetPathByHash(hash [sha256.Size]byte)string{
	for k, v := range PathsToHash{
		if v == hash{
			return k
		}
	}
	return ""
}

// type Image struct {
// 	Id [sha256.Size]byte
// 	Image *ebiten.Image
// }

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
			// Images[reg.Split(path, -1)[0]] = &Image{
				// Id: sha256.Sum256(f),
				// Image: img,
			// }
		}

		wg.Done()
	}()

}
