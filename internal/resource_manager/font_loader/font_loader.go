package fontloader

import (
	"embed"
	"regexp"
	"sync"

	"github.com/golang/freetype/truetype"
	"github.com/sirupsen/logrus"
	"golang.org/x/image/font"

	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/font_loader/collection"
)

var mu = sync.Mutex{}

//Process loading of font assets
func Load(e embed.FS, extension, path string, wg *sync.WaitGroup) {
	if extension != "ttf" {
		return
	}

	wg.Add(1)
	go func() {
		f, err := e.ReadFile(path)
		if err != nil {
			logrus.Fatal("error happened opening font file from embedded fs: ", err)
		}

		ff, err := truetype.Parse(f)
		if err != nil {
			logrus.Fatal("error happened parsing font file from embedded fs: ", err)
		}

		reg := regexp.MustCompile(`\.[a-z0-9]*$`)
		if reg.MatchString(path) {
			fontPath := reg.Split(path, -1)[0]
			mu.Lock()

			collection.FontCollection[fontPath] =
				truetype.NewFace(ff, &truetype.Options{
					Size:    9,
					Hinting: font.HintingFull,
				})

			mu.Unlock()
		}

		wg.Done()
	}()
}
