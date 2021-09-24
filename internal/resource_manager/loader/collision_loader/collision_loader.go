package collisionloader

import (
	// "log"
	// "regexp"
	"log"
	"regexp"
	"sync"

	"github.com/BurntSushi/toml"
	// "github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	RawCollisions = make(map[string]Collision)
	mu = sync.Mutex{}
)

type Collision struct {
	X float64
	Y float64
}

func Load(motherDir, extension, path string, wg *sync.WaitGroup) {
	if extension != "toml" {
		return
	}

	wg.Add(1)
	go func(){
		var ds Collision
		if _, err := toml.DecodeFile(path, &ds); err != nil{
			log.Fatalln(err)
		}

		path = path[len(motherDir):]
		reg := regexp.MustCompile(`\.[a-z]*$`)
		if reg.MatchString(path) {
			mu.Lock()
			defer mu.Unlock()
			RawCollisions[reg.Split(path, -1)[0]] = ds
		}

		wg.Done()
	}()
}
