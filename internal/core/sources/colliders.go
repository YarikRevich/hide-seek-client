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

type CollidersBatch map[image.Rectangle]CollidersModel

//Returns collider object collides with
func (cb *CollidersBatch) GetCollider(r image.Rectangle) (CollidersModel, bool) {
	for q, v := range *cb {
		if ((q.Min.X <= r.Max.X &&
			q.Min.X >= r.Min.X) || (q.Min.X >= r.Max.X &&
			q.Min.X <= r.Min.X)) &&
			((q.Min.Y <= r.Max.Y &&
				q.Min.Y >= r.Min.Y) || (q.Min.Y >= r.Max.Y &&
				q.Min.Y <= r.Min.Y)) {
			return v, true
		}
	}
	return CollidersModel{}, false
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

		colliders := make(map[image.Rectangle]CollidersModel)

		for _, l := range gameMap.Layers {
			for _, t := range l.Tiles {
				if !t.IsNil() {
					var collider CollidersModel

					tileID := int(t.ID - t.Tileset.FirstGID)
					x := (tileID % gameMap.Width) * gameMap.TileWidth
					y := (tileID % gameMap.Height) * gameMap.TileHeight

					for _, e := range t.Tileset.Tiles {
						if int(e.ID) == tileID {
							collider.Properties.Walkable = e.Properties.GetBool("walkable")
							collider.Properties.Collision = e.Properties.GetBool("collision")
						}
					}

					collider.Layer = l.Name

					colliders[image.Rect(x, y, x+gameMap.TileWidth, y+gameMap.TileHeight)] = collider
					// collider.Rect =

					// colliders = append(colliders, collider)
				}
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
		return nil, fmt.Errorf("image with path '%s' not found", path)
	}

	return colliders, nil
}

func NewColliders() *Colliders {
	return &Colliders{Collection: make(map[string]CollidersBatch)}
}
