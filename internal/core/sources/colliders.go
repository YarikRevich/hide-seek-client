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

//Model used only for input files parsing
type CollidersParseModel struct {
	ImageHeight int `json:"imageheight"`
	ImageWidth  int `json:"imagewidth"`
	TileHeight  int `json:"tileheight"`
	TileWidth   int `json:"tilewidth"`
	Tiles       []struct {
		Objectgroup struct {
			Id      int `json:"id"`
			Objects []struct {
				Id     int     `json:"id"`
				Height float64 `json:"height"`
				Width  float64 `json:"width"`
				X      float64 `json:"x"`
				Y      float64 `json:"y"`
			} `json:"objects"`
		} `json:"objectgroup"`
		Properties []struct {
			Name  string      `json:"name"`
			Type  string      `json:"type"`
			Value interface{} `json:"value"`
		} `json:"properties"`
	} `json:"tiles"`
}

type CollidersModel struct {
	Name                  string
	X, Y                  float64
	TileHeight, TileWidth int
	Properties            struct {
		IsCollision, IsWalkable bool
	}
}

func (m *CollidersModel) GetSizeMaxX() float64 {
	return (m.X + float64(m.TileWidth/2))
}

func (m *CollidersModel) GetSizeMaxY() float64 {
	return (m.Y + float64(m.TileHeight/2))
}

func (m *CollidersModel) GetSizeMinX() float64 {
	return (m.X - float64(m.TileWidth/2))
}

func (m *CollidersModel) GetSizeMinY() float64 {
	return (m.Y - float64(m.TileHeight/2))
}

type Colliders struct {
	sync.Mutex

	Collection map[string][]*CollidersModel
}

func (c *Colliders) loadFile(fs embed.FS, path string) {
	if reg := regexp.MustCompile(`\.json*$`); reg.MatchString(path) {
		file, err := fs.ReadFile(path)
		if err != nil {
			logrus.Fatal("error happened opening image file from embedded fs", err)
		}

		var collidersParseModel CollidersParseModel
		if err := json.Unmarshal(file, &collidersParseModel); err != nil {
			logrus.Fatal(err)
		}

		for _, v := range collidersParseModel.Tiles {
			for _, o := range v.Objectgroup.Objects {

				collidersModel := &CollidersModel{
					X:          float64((int(o.Id-v.Objectgroup.Id) % (int(collidersParseModel.ImageWidth / collidersParseModel.TileWidth)))) * float64(collidersParseModel.TileWidth),
					Y:          float64((int(o.Id-v.Objectgroup.Id) / (int(collidersParseModel.ImageHeight / collidersParseModel.TileHeight)))) * float64(collidersParseModel.TileHeight),
					TileHeight: collidersParseModel.TileHeight,
					TileWidth:  collidersParseModel.TileWidth,
					Name:       path,
				}
				for _, v := range v.Properties {
					switch v.Name {
					case "walkable":
						collidersModel.Properties.IsWalkable = v.Value.(bool)
					case "collision":
						collidersModel.Properties.IsCollision = v.Value.(bool)
					}
				}

				c.Lock()
				path = reg.Split(path, -1)[0]
				c.Collection[path] = append(c.Collection[path], collidersModel)
				c.Unlock()
			}
		}
	}
}

func (c *Colliders) Load(fs embed.FS, path string, wg *sync.WaitGroup) {
	NewParser(fs, path, c.loadFile).Parse()
	wg.Done()
}

func (c *Colliders) GetCollider(path string) ([]*CollidersModel, error) {
	path = filepath.Join("dist/colliders", path)

	image, ok := c.Collection[path]
	if !ok {
		// logrus.Fatal(fmt.Sprintf("image with path '%s' not found", path))
		return nil, fmt.Errorf("image with path '%s' not found", path)
	}

	return image, nil
}

func NewColliders() *Colliders {
	return &Colliders{Collection: make(map[string][]*CollidersModel)}
}
