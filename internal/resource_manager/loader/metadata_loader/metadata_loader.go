package collisionloader

import (
	"fmt"
	"log"
	"regexp"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

var (
	Metadata = make(map[string]M)
	mu       = sync.Mutex{}
)

type M struct {
	Size struct {
		Width  float64
		Height float64
	}
	Margins struct {
		LeftMargin float64
		TopMargin  float64
	}
	Animation struct {
		Delay float64
		FrameNum float64
		FrameX      float64
		FrameY      float64
		FrameWidth  float64
		FrameHeight float64
	}
}

func GetMetadata(path string)M{
	i, ok := Metadata[path]
	if !ok{
		logrus.Fatal(fmt.Sprintf("image with path '%s' not found", path))
	}
	return i
}

func Load(motherDir, extension, path string, wg *sync.WaitGroup) {
	if extension != "toml" {
		return
	}

	wg.Add(1)
	go func() {
		var ds M
		if _, err := toml.DecodeFile(path, &ds); err != nil {
			log.Fatalln(err)
		}

		path = path[len(motherDir):]
		reg := regexp.MustCompile(`\.[a-z]*$`)
		if reg.MatchString(path) {
			mu.Lock()
			defer mu.Unlock()
			Metadata[reg.Split(path, -1)[0]] = ds
		}

		wg.Done()
	}()
}
