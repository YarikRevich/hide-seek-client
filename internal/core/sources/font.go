package sources

import (
	"embed"
	"fmt"
	"path/filepath"
	"regexp"
	"sync"

	"github.com/golang/freetype/truetype"
	"github.com/sirupsen/logrus"
	"golang.org/x/image/font"
)

type Font struct {
	sync.Mutex

	Collection map[string]font.Face
}

func (f *Font) loadFile(fs embed.FS, path string) {
	file, err := fs.ReadFile(path)
	if err != nil {
		logrus.Fatal("error happened opening font file from embedded fs: ", err)
	}

	ff, err := truetype.Parse(file)
	if err != nil {
		logrus.Fatal("error happened parsing font file from embedded fs: ", err)
	}

	reg := regexp.MustCompile(`\.[a-z0-9]*$`)
	if reg.MatchString(path) {
		fontPath := reg.Split(path, -1)[0]
		f.Lock()

		f.Collection[fontPath] =
			truetype.NewFace(ff, &truetype.Options{
				Size:    9,
				Hinting: font.HintingFull,
			})

		f.Unlock()
	}
}

func (f *Font) Load(fs embed.FS, path string, wg *sync.WaitGroup) {
	NewParser(fs, path, f.loadFile).Parse()
	wg.Done()
}

func (f *Font) GetFont(path string) font.Face {
	path = filepath.Join("assets/fonts", path)

	font, ok := f.Collection[path]
	if !ok {
		logrus.Fatal(fmt.Sprintf("font with path '%s' not found", path))
	}
	return font
}

func NewFont() *Font {
	return &Font{Collection: make(map[string]font.Face)}
}
