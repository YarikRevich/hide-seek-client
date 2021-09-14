package imageloader

import (
	_ "image/png"
	"io/fs"
	"io/ioutil"
	"log"
	"strings"

	"github.com/YarikRevich/Hide-Seek-with-Guns/internal/paths"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	Images = make(map[string]*ebiten.Image)
)

func processImagesInDir(dirname string, files []fs.FileInfo) {
	for _, v := range files {
		path := dirname + "/" + v.Name()
		if v.IsDir() {
			processImageDir(path)
		} else {
			nameSplit := strings.Split(v.Name(), ".")
			if nameSplit[len(nameSplit)-1] != "png"{
				continue
			}

			img, _, err := ebitenutil.NewImageFromFile(path)
			if err != nil {
				log.Fatalln(err)
			}
			Images[path] = img
		}
	}
}

func processImageDir(dirname string) {
	d, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatalln(err)
	}
	processImagesInDir(dirname, d)
}

func LoadImages() {
	processImageDir(paths.GAME_ASSETS_DIR)
}
