package sources

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"path/filepath"
	"strings"

	"github.com/YarikRevich/hide-seek-client/assets"
	"github.com/YarikRevich/hide-seek-client/internal/core/camera"
	"github.com/YarikRevich/hide-seek-client/internal/core/primitives"
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/lafriks/go-tiled"
	"github.com/sirupsen/logrus"
)

type Tile struct {
	Image    *ebiten.Image
	Position image.Point

	Triangles []*types.Triangle

	Faces       []TileFace
	Layer       string
	LayerNum    int
	TileNum     int
	ColorMatrix []color.Color
	Properties  struct {
		Collision, Spawn bool
	}
}

func (t *Tile) GetTriangleIndicies() int {
	return len(t.Triangles)
}

type AnimationFrame struct {
	Duration uint32
	TileID   image.Point
}

type Animation struct {
	Frames                     []*AnimationFrame
	CurrentFrame, DelayTrigger int
}

//Starts animation if hasn't been started
//or continues if not
func (a *Animation) Proceed() {
	a.DelayTrigger++
	a.DelayTrigger %= int(a.Frames[a.CurrentFrame].Duration)
	if a.DelayTrigger == 0 {
		a.CurrentFrame++
		a.CurrentFrame %= len(a.Frames)
	}
}

//Stops animation and returns frame and frame
//change trigger to start position
func (a *Animation) Reset() {
	a.DelayTrigger = 0
	a.CurrentFrame = 0
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

type OrthographicTile struct {
	Tile            *Tile
	Rotation, Pitch float64
}

type TileFace int

const (
	Floor TileFace = iota
	South
	North
	East
	West
	Top
)

type OrthographicTilebatch struct {
	IsWall bool
	Tiles  map[TileFace]*OrthographicTile
	// Floor, Top, North, South, West, East *OrthographicTile
}

//Extension for tilemap
type OrthographicTilemap map[image.Point]*OrthographicTilebatch

type Quad struct {
	Tile
	Points [4]types.Vec3
}

type CubeOpts struct {
	sm *screen.ScreenManager

	Scale, Position types.Vec2
	Angle, Pitch    float64
	CameraPosition  types.Vec3
}

//Returns orthographic projection of the cube
func CreateOrthographicCube(opts CubeOpts) [8]types.Vec3 {
	var unitCube, rotCube, worldCube, projCube [8]types.Vec3

	unitCube[0] = types.Vec3{}
	unitCube[1] = types.Vec3{X: opts.Scale.X}
	unitCube[2] = types.Vec3{X: opts.Scale.X, Y: -opts.Scale.Y}
	unitCube[3] = types.Vec3{Y: -opts.Scale.Y}
	unitCube[4] = types.Vec3{Z: opts.Scale.Y}
	unitCube[5] = types.Vec3{X: opts.Scale.X, Z: opts.Scale.Y}
	unitCube[6] = types.Vec3{X: opts.Scale.X, Y: -opts.Scale.Y, Z: opts.Scale.Y}
	unitCube[7] = types.Vec3{Y: -opts.Scale.Y, Z: opts.Scale.Y}

	for i := 0; i < 8; i++ {
		unitCube[i].X += (opts.Position.X*opts.Scale.X - opts.CameraPosition.X)
		unitCube[i].Y += -opts.CameraPosition.Y
		unitCube[i].Z += (opts.Position.Y*opts.Scale.Y - opts.CameraPosition.Z)
	}

	s := math.Sin(opts.Angle)
	c := math.Cos(opts.Angle)
	for i := 0; i < 8; i++ {
		rotCube[i].X = unitCube[i].X*c + unitCube[i].Z*s
		rotCube[i].Y = unitCube[i].Y
		rotCube[i].Z = unitCube[i].X*-s + unitCube[i].Z*c
	}

	s = math.Sin(opts.Pitch)
	c = math.Cos(opts.Pitch)
	for i := 0; i < 8; i++ {
		worldCube[i].X = rotCube[i].X
		worldCube[i].Y = rotCube[i].Y*c - rotCube[i].Z*s
		worldCube[i].Z = rotCube[i].Y*s + rotCube[i].Z*c
	}

	screenSize := opts.sm.GetSize()

	for i := 0; i < 8; i++ {
		projCube[i].X = worldCube[i].X + screenSize.X*0.5
		projCube[i].Y = worldCube[i].Y + screenSize.Y*0.5
		projCube[i].Z = worldCube[i].Z
	}

	return projCube
}

//Returns perspective projection of the cube
func CreatePerspectiveCube(opts CubeOpts) [8]types.Vec3 {
	var unitCube, rotCube, worldCube, projCube [8]types.Vec3

	unitCube[0] = types.Vec3{}
	unitCube[1] = types.Vec3{X: opts.Scale.X}
	unitCube[2] = types.Vec3{X: opts.Scale.X, Y: -opts.Scale.Y}
	unitCube[3] = types.Vec3{Y: -opts.Scale.Y}
	unitCube[4] = types.Vec3{Z: opts.Scale.Y}
	unitCube[5] = types.Vec3{X: opts.Scale.X, Z: opts.Scale.Y}
	unitCube[6] = types.Vec3{X: opts.Scale.X, Y: -opts.Scale.Y, Z: opts.Scale.Y}
	unitCube[7] = types.Vec3{Y: -opts.Scale.Y, Z: opts.Scale.Y}

	for i := 0; i < 8; i++ {
		unitCube[i].X += (opts.Position.X*opts.Scale.X - opts.CameraPosition.X)
		unitCube[i].Y += -opts.CameraPosition.Y
		unitCube[i].Z += (opts.Position.Y*opts.Scale.Y - opts.CameraPosition.Z)
	}

	s := math.Sin(opts.Angle)
	c := math.Cos(opts.Angle)
	for i := 0; i < 8; i++ {
		rotCube[i].X = unitCube[i].X*c + unitCube[i].Z*s
		rotCube[i].Y = unitCube[i].Y
		rotCube[i].Z = unitCube[i].X*-s + unitCube[i].Z*c
	}

	s = math.Sin(opts.Pitch)
	c = math.Cos(opts.Pitch)
	for i := 0; i < 8; i++ {
		worldCube[i].X = rotCube[i].X
		worldCube[i].Y = rotCube[i].Y*c - rotCube[i].Z*s
		worldCube[i].Z = rotCube[i].Y*s + rotCube[i].Z*c
	}

	screenSize := opts.sm.GetSize()

	for i := 0; i < 8; i++ {
		projCube[i].X = worldCube[i].X + screenSize.X*0.5
		projCube[i].Y = worldCube[i].Y + screenSize.Y*0.5
		projCube[i].Z = worldCube[i].Z
	}

	return projCube
}

type FaceQuadOpts struct {
	CubeOpts
	Faces                                         []TileFace
	OrthographicProjection, PerspectiveProjection bool
}

func GetFaceQuad(opts FaceQuadOpts) []Quad {
	var r []Quad
	var projectionCube [8]types.Vec3

	cubeOpts := CubeOpts{
		sm:             opts.sm,
		Position:       opts.Position,
		Angle:          opts.Angle,
		Pitch:          opts.Pitch,
		CameraPosition: opts.CameraPosition,
		Scale:          opts.Scale,
	}

	if opts.OrthographicProjection {
		projectionCube = CreateOrthographicCube(cubeOpts)
	} else if opts.PerspectiveProjection {
		projectionCube = CreatePerspectiveCube(cubeOpts)
	}

	for _, face := range opts.Faces {
		switch face {
		case Floor:
			r = append(r, Quad{Points: [4]types.Vec3{projectionCube[4], projectionCube[0], projectionCube[1], projectionCube[5]}})
		case South:
			r = append(r, Quad{Points: [4]types.Vec3{projectionCube[3], projectionCube[0], projectionCube[1], projectionCube[2]}})
		case North:
			r = append(r, Quad{Points: [4]types.Vec3{projectionCube[6], projectionCube[5], projectionCube[4], projectionCube[7]}})
		case East:
			r = append(r, Quad{Points: [4]types.Vec3{projectionCube[7], projectionCube[4], projectionCube[0], projectionCube[3]}})
		case West:
			r = append(r, Quad{Points: [4]types.Vec3{projectionCube[2], projectionCube[1], projectionCube[5], projectionCube[6]}})
		case Top:
			r = append(r, Quad{Points: [4]types.Vec3{projectionCube[7], projectionCube[3], projectionCube[2], projectionCube[6]}})
		}
	}

	return r
}

type Tilemap struct {
	Name string

	//Collection of animations which
	//can be applied to the tilemap
	//SAVES STATE
	Animations map[int]*Animation
	Graph
	OrthographicTilemap

	Tiles      map[image.Point]*Tile
	Properties struct {
		//PHYSICS

		//Gravity acceleration
		G int

		//Acceleration
		A int

		//WORLD MAP properties

		//Contains IDs of Spawn Tiles
		Spawns []int64
	}

	MapSize, TileSize, TileCount types.Vec2
}

//Returns if tilemap has any animation
func (tm *Tilemap) IsAnimated() bool {
	return len(tm.Animations) != 0
}

func (tm *Tilemap) ToAPIMessage() {

}

func (tm *Tilemap) load(path string) error {
	gameMap, err := tiled.LoadFile(fmt.Sprintf("%s.%s", path, "tmx"), tiled.WithFileSystem(assets.Assets))
	if err != nil {
		return err
	}

	baseDir := filepath.Dir(path)

	tm.MapSize = types.Vec2{
		X: float64((gameMap.Width) * gameMap.TileWidth),
		Y: float64((gameMap.Height) * gameMap.TileHeight)}
	tm.TileSize = types.Vec2{
		X: float64(gameMap.TileWidth),
		Y: float64(gameMap.TileHeight)}
	tm.TileCount = types.Vec2{
		X: float64(gameMap.Width),
		Y: float64(gameMap.Height),
	}

	tempTileCollection := make(map[image.Point]*Tile)
	tempTileImageCache := make(map[string]*ebiten.Image)
	tempTileColorMatrixCache := make(map[string][]color.Color)
	tempTileTriangleCache := make(map[string][]*types.Triangle)

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

				if _, ok := tempTileColorMatrixCache[t.Tileset.Image.Source]; !ok {
					w, h := tile.Image.Size()
					colorMatrix := make([]color.Color, w*h)

					for y := 0; y < h; y++ {
						for x := 0; x < w; x++ {
							colorMatrix[y*int(tm.TileSize.X)+x] = tile.Image.At(x, y)
						}
					}

					tempTileColorMatrixCache[t.Tileset.Image.Source] = colorMatrix
				}

				if _, ok := tempTileTriangleCache[t.Tileset.Image.Source]; !ok {
					var tris []*types.Triangle

					tris = append(tris, primitives.CreateBottomQuad()...)

					tempTileTriangleCache[t.Tileset.Image.Source] = tris
				}

				tile.Triangles = tempTileTriangleCache[t.Tileset.Image.Source]

				tile.ColorMatrix = tempTileColorMatrixCache[t.Tileset.Image.Source]

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

				if left, ok := tempTileCollection[image.Point{X: x + gameMap.TileWidth, Y: y}]; ok {
					tm.Graph.AddNode(
						tile, left)
				}

				tempTileCollection[image.Point{X: x, Y: y}] = tile

				tm.OrthographicTilemap[image.Point{X: x, Y: y}] = &OrthographicTilebatch{
					Tiles: map[TileFace]*OrthographicTile{
						Floor: {Tile: tile},
						Top:   {Tile: tile},
						South: {Tile: tile},
						North: {Tile: tile},
						West:  {Tile: tile},
						East:  {Tile: tile},
					},
					IsWall: false,
				}
				tile.TileNum = len(tm.Tiles) + 1
				tile.Position = image.Point{X: x, Y: y}
				tm.Tiles[tile.Position] = tile

				if (i%gameMap.Width)*gameMap.TileWidth == ((gameMap.Width * gameMap.TileWidth) - gameMap.TileWidth) {
					y += gameMap.TileHeight
				}
			}
		}
	}

	tm.Name = strings.Split(path, ".")[0]
	tileMapCollection[tm.Name] = *tm

	return nil
}

type RenderTilemapOptsContext struct {
	Camera *camera.Camera
}

type RenderTilemapOpts struct {
	StickedTo                            *Tilemap
	StickedToPosition                    types.Vec2
	SurfacePosition, Scale               types.Vec2
	AutoScaleForbidden, CenterizedOffset bool
	AvailableFaces                       []TileFace

	OrthigraphicProjection, PerspectiveProjection bool

	RenderTilemapOptsContext
}

func (t *Tilemap) Render(sm *screen.ScreenManager, opts RenderTilemapOpts) {
	screenSize := sm.GetSize()
	screenScale := sm.GetScale()

	for k, v := range t.Tiles {
		if (float64(k.X)+opts.SurfacePosition.X-(t.TileSize.X) < screenSize.X && float64(k.Y)+opts.SurfacePosition.Y-(t.TileSize.Y) < screenSize.Y) &&
			(float64(k.X)+opts.SurfacePosition.X+t.TileSize.X > 0 && float64(k.Y)+opts.SurfacePosition.Y+t.TileSize.Y > 0) {
			drawOpts := &ebiten.DrawImageOptions{}

			if opts.OrthigraphicProjection {
				for y := k.Y; y < k.Y+(int(t.TileSize.Y)); y++ {
					for x := k.X; x < k.X+int(t.TileSize.X); x++ {
						quads := GetFaceQuad(FaceQuadOpts{
							CubeOpts: CubeOpts{
								sm:             sm,
								Scale:          types.Vec2{X: opts.Camera.Zoom, Y: opts.Camera.Zoom},
								Position:       types.Vec2{X: float64(x), Y: float64(y)},
								Angle:          opts.Camera.Angle,
								Pitch:          opts.Camera.Pitch,
								CameraPosition: opts.Camera.GetPosition()},
							Faces:                  v.Faces,
							OrthographicProjection: true})
						for _, quad := range quads {
							color := v.ColorMatrix[((y-k.Y)*int(t.TileSize.X) + (x - k.X))]
							var (
								l1X, l1Y = float32(quad.Points[0].X), float32(quad.Points[0].Y)
								l2X, l2Y = float32(quad.Points[1].X), float32(quad.Points[1].Y)
								l3X, l3Y = float32(quad.Points[2].X), float32(quad.Points[2].Y)
								l4X, l4Y = float32(quad.Points[3].X), float32(quad.Points[3].Y)
							)
							if ((l1X < float32(screenSize.X+t.TileSize.X) && l2X < float32(screenSize.X+t.TileSize.X) && l3X < float32(screenSize.X+t.TileSize.X) && l4X < float32(screenSize.X+t.TileSize.X)) &&
								(l1Y < float32(screenSize.Y+t.TileSize.Y) && l2Y < float32(screenSize.Y+t.TileSize.Y) && l3Y < float32(screenSize.Y+t.TileSize.Y) && l4Y < float32(screenSize.Y+t.TileSize.Y))) &&
								((l1X > float32(-t.TileSize.X) && l2X > float32(-t.TileSize.X) && l3X > float32(-t.TileSize.X) && l4X > float32(-t.TileSize.X)) &&
									(l1Y > float32(-t.TileSize.Y) && l2Y > float32(-t.TileSize.Y) && l3Y > float32(-t.TileSize.Y) && l4Y > float32(-t.TileSize.Y))) {
								var path vector.Path
								path.LineTo(l1X, l1Y)
								path.LineTo(l2X, l2Y)
								path.LineTo(l3X, l3Y)
								path.LineTo(l4X, l4Y)
								path.LineTo(l1X, l1Y)
								path.Fill(sm.Image, &vector.FillOptions{Color: color})
							}
						}
					}
				}
			} else if opts.PerspectiveProjection {

			} else {
				if !opts.AutoScaleForbidden {
					drawOpts.GeoM.Scale(1/screenScale.X, 1/screenScale.Y)
				}

				if opts.StickedTo != nil {
					drawOpts.GeoM.Translate(opts.StickedToPosition.X, opts.StickedToPosition.Y)
				}

				drawOpts.GeoM.Translate(float64(k.X), float64(k.Y))

				if opts.CenterizedOffset {
					drawOpts.GeoM.Translate(-t.MapSize.X/2, -t.MapSize.Y/2)
				}

				if opts.Scale.X != 0 && opts.Scale.Y != 0 {
					drawOpts.GeoM.Scale(opts.Scale.X, opts.Scale.Y)
				}

				drawOpts.GeoM.Translate(opts.SurfacePosition.X, opts.SurfacePosition.Y)
				sm.Image.DrawImage(v.Image, drawOpts)
			}
		}
	}

	if opts.OrthigraphicProjection {
		// vp := opts.RenderTilemapOptsContext.Camera.GetView().GetMultiplied(opts.RenderTilemapOptsContext.Camera.GetProjection(sm))

		// for _, v := range orthographicPostRender {
		// 	var verts []ebiten.Vertex
		// 	for _, t := range v.Triangles {
		// 		// fmt.Println(t.Vertices[0].Position)
		// 		v0 := vp.GetMultipiedVector(t.Vertices[0].Position)
		// 		v1 := vp.GetMultipiedVector(t.Vertices[1].Position)
		// 		v2 := vp.GetMultipiedVector(t.Vertices[2].Position)
		// 		fmt.Println(v0, v1, v2)

		// 		var ev0, ev1, ev2 ebiten.Vertex

		// 		ev0.DstX = float32(math.Round(v0[0]))
		// 		ev0.DstY = float32(math.Round(v0[1]))
		// 		ev0.ColorR = 1
		// 		ev0.ColorG = 1
		// 		ev0.ColorB = 1
		// 		ev0.ColorA = 1

		// 		ev1.DstX = float32(math.Round(v1[0]))
		// 		ev1.DstY = float32(math.Round(v1[1]))
		// 		ev1.ColorR = 1
		// 		ev1.ColorG = 1
		// 		ev1.ColorB = 1
		// 		ev1.ColorA = 1

		// 		ev2.DstX = float32(math.Round(v2[0]))
		// 		ev2.DstY = float32(math.Round(v2[1]))
		// 		ev2.ColorR = 1
		// 		ev2.ColorG = 1
		// 		ev2.ColorB = 1
		// 		ev2.ColorA = 1

		// 		verts = append(verts, ev0, ev1, ev2)
		// 	}
		// 	var indices []uint16
		// 	// for q := 0; q < v.GetTriangleIndicies()*3/2; q++ {
		// 	// 	indices = append(indices, uint16(q))
		// 	// }
		// 	fmt.Println(verts[0].DstX, indices)
		// 	sm.Image.DrawTriangles(verts, indices, t.Tiles[image.Point{0, 0}].Image, nil)
		// // }
		// var verts []ebiten.Vertex

		// var vx1, vx2, vx3, vx4, vx5, vx6 ebiten.Vertex

		// // NewVertex(1, -1, -1, 1, 0),
		// NewVertex(1, -1, 1, 1, 1),
		// NewVertex(-1, -1, -1, 0, 0),

		// NewVertex(-1, -1, -1, 0, 0),
		// NewVertex(1, -1, 1, 1, 1),
		// NewVertex(-1, -1, 1, 0, 1),

		// var path vector.Path
		// vector.Path
		// path.MoveTo(10, 10)

		// path.LineTo(10, 10)
		// path.LineTo(10, 20)
		// path.LineTo(20, 20)
		// path.LineTo(20, 10)

		// path.MoveTo(20, 20)
		// path.LineTo(20, 70)
		// path.LineTo(70, 70)
		// path.LineTo(70, 60)
		// path.LineTo(30, 60)
		// path.LineTo(30, 50)
		// path.LineTo(70, 50)
		// path.LineTo(70, 40)
		// path.LineTo(30, 40)
		// path.LineTo(30, 30)
		// path.LineTo(70, 30)
		// path.LineTo(70, 20)

		// path.LineTo(10, 10)
		// path.LineTo(20, 10)
		// path.LineTo(20, -10)
		// path.MoveTo(10, 10)
		// path.LineTo(1, 1)
		// path.LineTo(1, -1)
		// path.LineTo(-1, -1)
		// path.LineTo(-1, -1)
		// path.LineTo(1, 1)
		// path.LineTo(-1, -1)
		// var path vector.Path
		// // path.MoveTo(500, 500)
		// z := float32(opts.RenderTilemapOptsContext.Camera.Zoom)
		// path.LineTo(1*z, 1*z)
		// path.LineTo(1*z, -1*z)
		// path.LineTo(-1*z, -1*z)
		// path.LineTo(-1*z, 1*z)
		// path.LineTo(1*z, 1*z)

		// path.Fill(sm.Image, &vector.FillOptions{color.Opaque})

		// // const z = 100
		// op := &ebiten.DrawImageOptions{}
		// op.GeoM.Scale(4, 4)
		// op.GeoM.Translate(100, 100)
		// opts.GeoM.Translate(100/z, 100/z)
		// opts.GeoM.Translate()

		// sm.Image.DrawImage(t.Tiles[image.Point{0, 0}].Image, op)

		// const z = 3
		// opts1 := &ebiten.DrawImageOptions{}

		// opts1.GeoM.Scale(4/opts.RenderTilemapOptsContext.Camera.Zoom, 4/opts.RenderTilemapOptsContext.Camera.Zoom)
		// // opts1.GeoM.Translate(100, 100)
		// opts1.GeoM.Translate(opts.RenderTilemapOptsContext.Camera.Position.X/opts.RenderTilemapOptsContext.Camera.Zoom, opts.RenderTilemapOptsContext.Camera.Position.Y/opts.RenderTilemapOptsContext.Camera.Zoom)
		// // opts.GeoM.Translate()

		// sm.Image.DrawImage(t.Tiles[image.Point{0, 0}].Image, opts1)
		// opts := &vector.FillOptions{}
		// fmt.Println(path)
		// opts.Color = color.RGBA{120, 214, 192, 255}
		// // image.
		// path.Fill(sm.Image, opts)

		// op := ebiten.DrawImageOptions{}
		// sm.Image.DrawImage(t.Tiles[image.Point{0, 0}].Image, &op)

		// vx1.ColorR = 1
		// vx1.ColorG = 1
		// vx1.ColorB = 1
		// vx1.ColorA = 1

		// vx2.ColorR = 1
		// vx2.ColorG = 1
		// vx2.ColorB = 1
		// vx2.ColorA = 1

		// vx3.ColorR = 1
		// vx3.ColorG = 1
		// vx3.ColorB = 1
		// vx3.ColorA = 1

		// vx4.ColorR = 1
		// vx4.ColorG = 1
		// vx4.ColorB = 1
		// vx4.ColorA = 1

		// vx5.ColorR = 1
		// vx5.ColorG = 1
		// vx5.ColorB = 1
		// vx5.ColorA = 1

		// vx6.ColorR = 1
		// vx6.ColorG = 1
		// vx6.ColorB = 1
		// vx6.ColorA = 1

		// verts = append(verts, vx1, vx2, vx3, vx4, vx5, vx6)

		// var indices []uint16
		// for q := 0; q < 3; q++ {
		// 	indices = append(indices, uint16(q))
		// }
		// sm.Image.DrawTriangles(verts, indices, t.Tiles[image.Point{0, 0}].Image, nil)

		// op := &ebiten.DrawTrianglesOptions{}
		// // op.Address = ebiten.AddressUnsafe

		// img := ebiten.NewImage(3, 3)
		// img.Fill(color.Opaque)

		// var (
		// 	centerX = screenSize.X / 2
		// 	centerY = screenSize.Y / 2
		// 	// r       = 50
		// )

		// vs := []ebiten.Vertex{}
		// // for i := 0; i < 4; i++ {
		// // 	rate := float64(i) / float64(4)
		// // 	fmt.Println(rate)
		// // cr := 0.0
		// // cg := 0.0
		// // cb := 0.0
		// // if rate < 1.0/3.0 {
		// // 	cb = 2 - 2*(rate*3)
		// // 	cr = 2 * (rate * 3)
		// // }
		// // if 1.0/3.0 <= rate && rate < 2.0/3.0 {
		// // 	cr = 2 - 2*(rate-1.0/3.0)*3
		// // 	cg = 2 * (rate - 1.0/3.0) * 3
		// // }
		// // if 2.0/3.0 <= rate {
		// // 	cg = 2 - 2*(rate-2.0/3.0)*3
		// // 	cb = 2 * (rate - 2.0/3.0) * 3
		// // }
		// vs = append(vs, ebiten.Vertex{
		// 	DstX: 0,
		// 	DstY: float32(centerY),
		// 	SrcX: 0,
		// 	SrcY: 0,
		// 	// ColorR: float32(cr),
		// 	// ColorG: float32(cg),
		// 	// ColorB: float32(cb),
		// 	// ColorA: 1,
		// 	ColorR: 1,
		// 	ColorG: 1,
		// 	ColorB: 1,
		// 	ColorA: 1,
		// })

		// vs = append(vs, ebiten.Vertex{
		// 	DstX: float32(centerX),
		// 	DstY: float32(centerY),
		// 	SrcX: 0,
		// 	SrcY: 0,
		// 	// ColorR: float32(cr),
		// 	// ColorG: float32(cg),
		// 	// ColorB: float32(cb),
		// 	// ColorA: 1,
		// 	ColorR: 1,
		// 	ColorG: 1,
		// 	ColorB: 1,
		// 	ColorA: 1,
		// })

		// vs = append(vs, ebiten.Vertex{
		// 	DstX: float32(centerX),
		// 	DstY: 0,
		// 	SrcX: 0,
		// 	SrcY: 0,
		// 	// ColorR: float32(cr),
		// 	// ColorG: float32(cg),
		// 	// ColorB: float32(cb),
		// 	// ColorA: 1,
		// 	ColorR: 1,
		// 	ColorG: 1,
		// 	ColorB: 1,
		// 	ColorA: 1,
		// })
		// // }

		// vs = append(vs, ebiten.Vertex{
		// 	DstX:   0,
		// 	DstY:   0,
		// 	SrcX:   0,
		// 	SrcY:   0,
		// 	ColorR: 1,
		// 	ColorG: 1,
		// 	ColorB: 1,
		// 	ColorA: 1,
		// })
		// // var verts []ebiten.Vertex
		// // for i := 0; i < 3; i++ {
		// // 	verts = append(verts, ebiten.Vertex{
		// // 		SrcX: 0, SrcY: 0, DstX: float32(160*math.Cos(float64(2)*math.Pi*float64(i/3)) + (screenSize.X / 2)), DstY: float32(160*math.Sin(float64(2)*math.Pi*float64(i/3)) + (screenSize.Y / 2)), ColorR: 120, ColorG: 192, ColorB: 255, ColorA: 1})
		// // }
		// // verts = append(verts, ebiten.Vertex{
		// // 	DstX:   float32(screenSize.X / 2),
		// // 	DstY:   float32(screenSize.Y / 2),
		// // 	SrcX:   0,
		// // 	SrcY:   0,
		// // 	ColorR: 1,
		// // 	ColorG: 1,
		// // 	ColorB: 1,
		// // 	ColorA: 1,
		// // })

		// indices := []uint16{}
		// for i := 0; i < 4; i++ {
		// 	// fmt.Println(uint16(i), uint16(i+1)%uint16(3), uint16(3))
		// 	indices = append(indices, uint16(i), uint16(i+1)%uint16(3), uint16(3))
		// }
		// sm.Image.DrawTriangles(vs, indices, t.Tiles[image.Point{0, 0}].Image, op)
		// sort.Slice(orthographicPostRender, func(i, j int) bool {
		// 	return orthographicPostRender[i].TileNum < orthographicPostRender[j].TileNum
		// })

		// orthographicSurface := ebiten.NewImage(int(t.MapSize.X), int(t.MapSize.Y))
		// for _, v := range orthographicPostRender {
		// 	w, h := v.Image.Size()
		// 	for y := 0; y < h; y++ {
		// 		for x := 0; x < w; x++ {
		// 			orthographicSurface.Set(v.Position.X+x, v.Position.Y+y, v.ColorMatrix[y*int(t.TileSize.X)+x])
		// 		}
		// 	}
		// }

		// w, h := orthographicSurface.Size()
		// op := &ebiten.DrawImageOptions{}
		// for i := 0; i < h; i++ {
		// 	op.GeoM.Reset()

		// 	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
		// 	lineW := (w + i*3/4)
		// 	x := -float64(lineW) / float64(w) / 2

		// 	fmt.Println(opts.CameraPitch)
		// 	// op.GeoM.Scale(float64(lineW)/float64(w), opts.CameraPitch)
		// 	if opts.CameraZoom > 0 {
		// 		op.GeoM.Scale(float64(lineW)/float64(w)*(opts.CameraZoom), opts.CameraZoom)
		// 	} else {
		// 		op.GeoM.Scale(float64(lineW)/float64(w), opts.CameraZoom)
		// 	}

		// 	//Makes rotation around Y but with offset, so it rounds around circle
		// 	//#####
		// 	// op.GeoM.Rotate(opts.CameraAngle)
		// 	//#####

		// 	op.GeoM.Translate(x+(float64(i)/opts.CameraAngle), float64(i)*(opts.CameraZoom))

		// 	op.GeoM.Translate(opts.CameraPosition.X, opts.CameraPosition.Y)

		// 	//Makes rotation around Y axis
		// 	//#####
		// 	op.GeoM.Rotate(opts.CameraAngle)
		// 	//#####

		// 	op.GeoM.Translate(screenSize.X/2, screenSize.Y/2)
		// 	// sm.Image.DrawImage(t.OrthographicTilemap[k].Tiles[OrthographicTileFace(c[4])].Tile.Image, opts)
		// 	sm.Image.DrawImage(orthographicSurface.SubImage(image.Rect(0, i, w, i+1)).(*ebiten.Image), op)
		// }
	}
}

func NewTilemap() Tilemap {
	return Tilemap{
		Tiles:               make(map[image.Point]*Tile),
		Animations:          make(map[int]*Animation),
		OrthographicTilemap: make(OrthographicTilemap),
		Graph:               make(Graph),
	}
}
