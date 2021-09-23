package loader

import (
	_ "image/png"
	"io/fs"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/YarikRevich/HideSeek-Client/internal/asset_manager/paths"
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
			processAssetDir(path)
		} else {
			nameSplit := strings.Split(v.Name(), ".")
			if nameSplit[len(nameSplit)-1] != "png"{
				continue
			}

			img, _, err := ebitenutil.NewImageFromFile(path)
			if err != nil {
				log.Fatalln(err)
			}

			path = path[len(paths.GAME_ASSETS_DIR):]
			reg := regexp.MustCompile(`\.[a-z]*$`)
			if reg.MatchString(path){
					Images[reg.Split(path, -1)[0]] = img
			}
		}
	}
}

func processAssetDir(dirname string) {
	d, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatalln(err)
	}
	processImagesInDir(dirname, d)
}

func LoadAssets() {
	processAssetDir(paths.GAME_ASSETS_DIR)
}


