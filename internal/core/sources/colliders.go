package sources

import (
	"embed"
	"fmt"
	"image"
	_ "image/png"
	"path/filepath"
	"regexp"
	"sync"

	"github.com/lafriks/go-tiled"
	"github.com/sirupsen/logrus"
)

type CollidersModel struct {
	Layer      string
	LayerNum   int
	Properties struct {
		Collision, Walkable bool
	}
}

func (cm *CollidersModel) IsCollision() bool {
	return cm.Properties.Collision
}

func (cm *CollidersModel) IsWalkable() bool {
	return cm.Properties.Walkable
}

type CollidersBatch struct {
	RawData map[image.Rectangle]CollidersModel
	Graph   map[image.Rectangle]map[image.Rectangle]CollidersModel
}

type Colliders struct {
	sync.Mutex

	Collection map[string]CollidersBatch
}

func (c *Colliders) loadFile(fs embed.FS, path string) {
	if reg := regexp.MustCompile(`\.tmx*$`); reg.MatchString(path) {
		gameMap, err := tiled.LoadFile(path, tiled.WithFileSystem(fs))
		if err != nil {
			logrus.Fatal(err)
		}

		colliders := CollidersBatch{
			RawData: make(map[image.Rectangle]CollidersModel),
			Graph:   make(map[image.Rectangle]map[image.Rectangle]CollidersModel),
		}

		for n, l := range gameMap.Layers {
			y := 0
			for i, t := range l.Tiles {
				if (i%gameMap.Width)*gameMap.TileWidth == ((gameMap.Width * gameMap.TileWidth) - gameMap.TileWidth) {
					y += gameMap.TileHeight
				}
				if !t.IsNil() {
					var collider CollidersModel

					collider.Layer = l.Name
					collider.LayerNum = n

					x := (i % gameMap.Width) * gameMap.TileWidth
					for _, w := range t.Tileset.Tiles {
						if w.ID == t.ID {
							collider.Properties.Walkable = w.Properties.GetBool("walkable")
							collider.Properties.Collision = w.Properties.GetBool("collision")
						}
					}

					fmt.Println(l.Tiles[0].Tileset.Tiles[0].Image)

					colliders.RawData[image.Rect(x-gameMap.TileWidth, y, x, y+gameMap.TileHeight)] = collider
				}
			}
		}

		for k := range colliders.RawData {
			colliders.Graph[k] = make(map[image.Rectangle]CollidersModel)

			topRect := image.Rect(k.Min.X, k.Min.Y-(gameMap.TileHeight*2), k.Max.X, k.Min.Y)
			bottomRect := image.Rect(k.Min.X, k.Max.Y, k.Max.X, k.Max.Y+(gameMap.TileHeight*2))
			rightRect := image.Rect(k.Max.X, k.Min.Y, k.Max.X+(gameMap.TileWidth*2), k.Max.Y)
			leftRect := image.Rect(k.Max.X, k.Min.Y-(gameMap.TileWidth*2), k.Min.X, k.Max.Y)

			if top, ok := colliders.RawData[topRect]; ok {
				colliders.Graph[k][topRect] = top
			}
			if bottom, ok := colliders.RawData[bottomRect]; ok {
				colliders.Graph[k][bottomRect] = bottom
			}

			if right, ok := colliders.RawData[rightRect]; ok {
				colliders.Graph[k][rightRect] = right
			}

			if left, ok := colliders.RawData[leftRect]; ok {
				colliders.Graph[k][leftRect] = left
			}
		}

		c.Lock()
		path = reg.Split(path, -1)[0]
		c.Collection[path] = colliders
		c.Unlock()
	}
}

func (c *Colliders) Load(fs embed.FS, path string, wg *sync.WaitGroup) {
	NewParser(fs, path, c.loadFile).Parse()
	wg.Done()
}

func (c *Colliders) GetCollider(path string) (CollidersBatch, error) {
	path = filepath.Join("dist/colliders", path)

	colliders, ok := c.Collection[path]
	if !ok {
		// logrus.Fatal(fmt.Sprintf("image with path '%s' not found", path))
		return CollidersBatch{}, fmt.Errorf("image with path '%s' not found", path)
	}

	return colliders, nil
}

func NewColliders() *Colliders {
	return &Colliders{Collection: make(map[string]CollidersBatch)}
}
