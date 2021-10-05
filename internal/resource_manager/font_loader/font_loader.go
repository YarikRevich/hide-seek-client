package fontloader

import (
	"embed"
	"fmt"
	"regexp"
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"

	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/font_loader/collection"
)

var mu = sync.Mutex{}

func Load(e embed.FS, extension, path string, wg *sync.WaitGroup) {
	if extension != "ttf" {
		return
	}

	wg.Add(1)
	go func() {
		f, err := e.ReadFile(path)
		if err != nil {
			logrus.Errorf("error happened opening font file from embedded fs: %w", err)
		}

		ff, err := truetype.Parse(f)
		if err != nil{
			logrus.Errorf("error happened parsing font file from embedded fs: %w", err)
		}

		reg := regexp.MustCompile(`\.[a-z0-9]*$`)
		if reg.MatchString(path) {
			fontPath := reg.Split(path, -1)[0]
			mu.Lock()

			for s := 0; s < 100; s++{
				face := truetype.NewFace(ff, &truetype.Options{
					Size: float64(s),
					DPI: 72,
					Hinting: font.HintingFull,
				})

				collection.FontCollection[fmt.Sprintf("%s_%d", fontPath, s)] = face
			}
	
			mu.Unlock()
		}

		wg.Done()
	}()
}
