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

	MapSize, TileSize, TileCount types.Vec2
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
	tm.TileCount = types.Vec2{
		X: float64(gameMap.TileWidth),
		Y: float64(gameMap.TileHeight),
	}

	tempTileCollection := make(map[image.Point]*Tile)
	tempTileImageCache := make(map[string]*ebiten.Image)

	for n, l := range gameMap.Layers {
		y := 0
		for i, t := range l.Tiles {

			if !t.IsNil() {
				tile := new(Tile)

				if _, ok := tempTileImageCache[t.Tileset.Image.Source]; !ok {
					file, err := assets.Assets.Open(filepath.Join(baseDir, t.Tileset.Image.Source))
					if err != nil {
						logrus.Fatalln(err)
					}
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

				for _, w := range t.Tileset.Tiles {
					if w.ID == t.ID {
						animation := new(Animation)
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
				tm.Tiles[image.Point{X: x, Y: y}] = tile

				if (i%gameMap.Width)*gameMap.TileWidth == ((gameMap.Width * gameMap.TileWidth) - gameMap.TileWidth) {
					y += gameMap.TileHeight
				}
			}
		}
	}

	tm.Name = strings.Split(path, ".")[0]
	tileMapCollection[tm.Name] = tm

	return nil
}

type RenderTilemapOpts struct {
	SurfacePosition, Scale               types.Vec2
	AutoScaleForbidden, CenterizedOffset bool
}

func (t *Tilemap) Render(sm *screen.ScreenManager, opts RenderTilemapOpts) {
	screenSize := sm.GetSize()
	screenScale := sm.GetScale()

	for k, v := range t.Tiles {
		if (float64(k.X)+opts.SurfacePosition.X-t.TileSize.X < screenSize.X && float64(k.Y)+opts.SurfacePosition.Y-t.TileSize.Y < screenSize.Y) &&
			(float64(k.X)+opts.SurfacePosition.X+t.TileSize.X > 0 && float64(k.Y)+opts.SurfacePosition.Y+t.TileSize.Y > 0) {
			drawOpts := &ebiten.DrawImageOptions{}

			if !opts.AutoScaleForbidden {
				drawOpts.GeoM.Scale(1/screenScale.X, 1/screenScale.Y)
			}

			if opts.Scale.X != 0 && opts.Scale.Y != 0 {
				drawOpts.GeoM.Scale(opts.Scale.X, opts.Scale.Y)
			}

			if opts.CenterizedOffset {
				drawOpts.GeoM.Translate(-t.TileSize.X*(t.TileCount.X/2), -t.TileSize.Y*(t.TileCount.Y/2))
			}

			drawOpts.GeoM.Translate(float64(k.X)+opts.SurfacePosition.X, float64(k.Y)+opts.SurfacePosition.Y)
			sm.Image.DrawImage(v.Image, drawOpts)
		}
	}
}

func NewTilemap() *Tilemap {
	return &Tilemap{
		Tiles:      make(map[image.Point]*Tile),
		Animations: make(map[int]*Animation),
		Graph:      make(Graph),
	}
}
