package sources

import (
	"fmt"
	"image"
	_ "image/png"
	"strings"

	"github.com/YarikRevich/hide-seek-client/assets"
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
)

type Tile struct {
	Rect       image.Rectangle
	Image      *ebiten.Image
	Layer      string
	LayerNum   int
	Properties struct {
		Collision, Spawn bool
	}
}

type Animation struct {
	Frames       []*tiled.AnimationFrame
	CurrentFrame int
}

type Graph map[*Tile][]*Tile

func (g *Graph) AddNode(key, value *Tile) {
	_, ok := (*g)[key]
	if !ok {
		(*g)[key] = append((*g)[key], value)
	}

	_, ok = (*g)[value]
	if !ok {
		(*g)[value] = append((*g)[value], key)
	}
}

type Tilemap struct {
	Name       string
	Animations map[int]*Animation
	Graph

	Tiles      map[uint32]*Tile
	Properties struct {
		//Contains IDs of Spawn Tiles
		Spawns []uint32
	}

	Size types.Vec2
}

func (tm *Tilemap) load(path string) error {
	gameMap, err := tiled.LoadFile(path, tiled.WithFileSystem(assets.Assets))
	if err != nil {
		return err
	}

	tm.Size = types.Vec2{
		X: float64(gameMap.Width * gameMap.TileWidth),
		Y: float64(gameMap.Height * gameMap.TileHeight)}

	tempTileCollection := make(map[image.Point]*Tile)
	for n, l := range gameMap.Layers {
		y := 0
		for i, t := range l.Tiles {
			if (i%gameMap.Width)*gameMap.TileWidth == ((gameMap.Width * gameMap.TileWidth) - gameMap.TileWidth) {
				y += gameMap.TileHeight
			}
			if !t.IsNil() {
				tile := new(Tile)

				tile.Layer = l.Name
				tile.LayerNum = n

				x := (i % gameMap.Width) * gameMap.TileWidth
				for _, w := range t.Tileset.Tiles {
					if w.ID == t.ID {
						animation := new(Animation)
						animation.Frames = w.Animation
						tm.Animations[len(tm.Animations)+1] = animation
						tile.Properties.Collision = w.Properties.GetBool("collision")
					}
				}

				tile.Rect = image.Rect(x, y, x+gameMap.TileWidth, y+gameMap.TileHeight)

				if top, ok := tempTileCollection[image.Point{X: x, Y: y - gameMap.TileHeight}]; ok {
					tm.Graph.AddNode(
						tile, top)
				}

				if bottom, ok := tempTileCollection[image.Point{X: x, Y: y + gameMap.TileHeight}]; ok {
					tm.Graph.AddNode(
						tile, bottom)
				}

				if right, ok := tempTileCollection[image.Point{X: x + gameMap.TileWidth, Y: y}]; ok {
					tm.Graph.AddNode(
						tile, right)
				}

				if left, ok := tempTileCollection[image.Point{X: x - gameMap.TileWidth, Y: y}]; ok {
					tm.Graph.AddNode(
						tile, left)
				}

				tempTileCollection[image.Point{X: x, Y: y}] = tile
				tm.Tiles[t.ID] = tile
			}
		}
	}

	tm.Name = strings.Split(path, ".")[0]
	tileMapCollection[tm.Name] = tm

	return nil
}

func (t *Tilemap) OnInteraction(s screen.ScreenManager) {

}

func (t *Tilemap) OnCollision(s screen.ScreenManager) {

}

type RenderTilemapOpts struct {
	Position types.Vec2
}

func (t *Tilemap) Render(sm *screen.ScreenManager, opts RenderTilemapOpts) {
	fmt.Println(sm)
	// screenSize := sm.GetSize()
	// for _, v := range t.Tiles {
	// 	v.Rect
	// 	// sm.RenderTile(v.Rect, v)
	// }
}
