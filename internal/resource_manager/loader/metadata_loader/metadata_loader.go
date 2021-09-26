package collisionloader

import (
	"log"
	"regexp"
	"sync"

	"github.com/BurntSushi/toml"
)

var (
	Metadata = make(map[string]M)
	mu            = sync.Mutex{}
)

type M struct {
	Size struct {
		Height float64
		Width  float64
	}
	Margins struct {
		LeftMargin float64
		TopMargin float64
	}
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
