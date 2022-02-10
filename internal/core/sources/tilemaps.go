package sources

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"path/filepath"
	"strings"

	"github.com/YarikRevich/hide-seek-client/assets"
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	"github.com/sirupsen/logrus"
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

type AnimationFrame struct {
	Duration uint32
	TileID   image.Point
}

type Animation struct {
	Frames                     []*AnimationFrame
	CurrentFrame, DelayTrigger int
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

	Tiles      map[image.Point]*Tile
	Properties struct {
		//Contains IDs of Spawn Tiles
		Spawns []int64
	}

	MapSize, TileSize types.Vec2
}

func (tm *Tilemap) load(path string) error {
	gameMap, err := tiled.LoadFile(fmt.Sprintf("%s.%s", path, "tmx"), tiled.WithFileSystem(assets.Assets))
	if err != nil {
		return err
	}

	baseDir := filepath.Dir(path)

	tm.MapSize = types.Vec2{
		X: float64(gameMap.Width * gameMap.TileWidth),
		Y: float64(gameMap.Height * gameMap.TileHeight)}
	tm.TileSize = types.Vec2{
		X: float64(gameMap.TileWidth),
		Y: float64(gameMap.TileHeight)}

	tempTileCollection := make(map[image.Point]*Tile)
	tempTileImageCache := make(map[string]*ebiten.Image)

	for n, l := range gameMap.Layers {
		y := 0
		for i, t := range l.Tiles {
			// fmt.Println(t)
			if (i%gameMap.Width)*gameMap.TileWidth == ((gameMap.Width * gameMap.TileWidth) - gameMap.TileWidth) {
				y += gameMap.TileHeight
			}

			if !t.IsNil() {
				tile := new(Tile)

				if _, ok := tempTileImageCache[t.Tileset.Image.Source]; !ok {
					file, err := assets.Assets.Open(filepath.Join(baseDir, t.Tileset.Image.Source))
					if err != nil {
						logrus.Fatalln(err)
					}
					// fmt.Println(file.Stat())
					pngFile, _, err := image.Decode(file)
					if err != nil {
						logrus.Fatalln(err)
					}
					if err := file.Close(); err != nil {
						logrus.Fatalln(err)
					}
					ebitenImage := ebiten.NewImageFromImage(pngFile)
					tempTileImageCache[t.Tileset.Image.Source] = ebitenImage
				}

				ebitenImage := tempTileImageCache[t.Tileset.Image.Source]
				subImage := ebitenImage.SubImage(t.GetTileRect())

				tile.Image = ebiten.NewImageFromImage(subImage)

				tile.Layer = l.Name
				tile.LayerNum = n

				x := (i % gameMap.Width) * gameMap.TileWidth
				// fmt.Println(t.Tileset.Image, len(t.Tileset.Tiles))

				for _, w := range t.Tileset.Tiles {
					if w.ID == t.ID {
						// tile.Image = w.Image
						animation := new(Animation)
						// fmt.Println("HERLLO")
						for _, a := range w.Animation {

							animation.Frames = append(animation.Frames, &AnimationFrame{
								Duration: a.Duration,
								TileID:   image.Point{X: x, Y: y},
							})
						}
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

				// fmt.Println(t.ID<<x, t.Tileset.FirstGID)
				tempTileCollection[image.Point{X: x, Y: y}] = tile
				tm.Tiles[image.Point{X: x, Y: y}] = tile
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
	// fmt.Println("START")
	for _, v := range t.Tiles {
		drawOpts := &ebiten.DrawImageOptions{}
		drawOpts.GeoM.Translate(float64(v.Rect.Max.X)-t.TileSize.X, float64(v.Rect.Max.Y)-t.TileSize.Y)
		sm.Image.DrawImage(v.Image, drawOpts)
	}
	// fmt.Println("END")
	// screenSize := sm.GetSize()
	// for _, v := range t.Tiles {
	// 	v.Rect
	// 	// sm.RenderTile(v.Rect, v)
	// }
}

func NewTilemap() *Tilemap {
	return &Tilemap{
		Tiles:      make(map[image.Point]*Tile),
		Animations: make(map[int]*Animation),
		Graph:      make(Graph),
	}
}
