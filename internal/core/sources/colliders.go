package sources

import (
	"embed"
	"encoding/json"
	"fmt"
	_ "image/png"
	"path/filepath"
	"regexp"
	"sync"

	"github.com/sirupsen/logrus"
)

type CollidersModel struct {
	ImageHeight int `json:"imageheight"`
	ImageWidth  int `json:"imagewidth"`
	TileHeight  int `json:"tileheight"`
	TileWidth   int `json:"tilewidth"`
	Tiles       []struct {
		Objectgroup struct {
			Id      int `json:"id"`
			Objects []struct {
				Id     int `json:"id"`
				Height int `json:"height"`
				Width  int `json:"width"`
				X      int `json:"x"`
				Y      int `json:"y"`
			} `json:"objects"`
		} `json:"objectgroup"`
	} `json:"tiles"`
}

type Colliders struct {
	sync.Mutex

	Collection map[string][]struct{ X, Y float64 }
}

func (c *Colliders) loadFile(fs embed.FS, path string) {
	if reg := regexp.MustCompile(`\.json*$`); reg.MatchString(path) {
		file, err := fs.ReadFile(path)
		if err != nil {
			logrus.Fatal("error happened opening image file from embedded fs", err)
		}

		var collidersModel CollidersModel
		if err := json.Unmarshal(file, &collidersModel); err != nil {
			logrus.Fatal(err)
		}

		for _, v := range collidersModel.Tiles {
			for _, o := range v.Objectgroup.Objects {
				c.Lock()
				path = reg.Split(path, -1)[0]
				c.Collection[path] = append(c.Collection[path], struct {
					X float64
					Y float64
				}{
					X: float64((int(o.Id-v.Objectgroup.Id) % (int(collidersModel.ImageWidth / collidersModel.TileWidth)))) * float64(collidersModel.TileWidth),
					Y: float64((int(o.Id-v.Objectgroup.Id) / (int(collidersModel.ImageHeight / collidersModel.TileHeight)))) * float64(collidersModel.TileHeight),
				})
				c.Unlock()
			}
		}
	}
}

func (c *Colliders) Load(fs embed.FS, path string, wg *sync.WaitGroup) {
	NewParser(fs, path, c.loadFile).Parse()
	wg.Done()
}

func (c *Colliders) GetCollider(path string) []struct{ X, Y float64 } {
	path = filepath.Join("dist/colliders", path)

	image, ok := c.Collection[path]
	if !ok {
		logrus.Fatal(fmt.Sprintf("image with path '%s' not found", path))
	}

	return image
}

func NewColliders() *Colliders {
	return &Colliders{Collection: make(map[string][]struct{ X, Y float64 })}
}
