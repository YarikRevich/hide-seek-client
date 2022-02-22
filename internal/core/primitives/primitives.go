package primitives

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
	"github.com/hajimehoshi/ebiten/v2"
)

//Create square by ebiten tools
func CreateSquare(a int) *ebiten.Image {
	return ebiten.NewImage(a, a)
}

func CreateBottomQuad() []*types.Triangle {
	var triangles []*types.Triangle

	verts := []*types.Vertex{
		types.NewVertex(1, -1, -1, 1, 0),
		types.NewVertex(1, -1, 1, 1, 1),
		types.NewVertex(-1, -1, -1, 0, 0),

		types.NewVertex(-1, -1, -1, 0, 0),
		types.NewVertex(1, -1, 1, 1, 1),
		types.NewVertex(-1, -1, 1, 0, 1),
	}

	for i := 0; i < len(verts); i += 3 {
		triangle := new(types.Triangle)
		triangle.Vertices = append(triangle.Vertices, verts[i], verts[i+1], verts[i+2])
		triangles = append(triangles, triangle)
	}

	return triangles
}

// // Top

// NewVertex(-1, 1, -1, 0, 0),
// NewVertex(1, 1, 1, 1, 1),
// NewVertex(1, 1, -1, 1, 0),

// NewVertex(-1, 1, 1, 0, 1),
// NewVertex(1, 1, 1, 1, 1),
// NewVertex(-1, 1, -1, 0, 0),

// // Front

// // NewVertex(-1, 1, 1, 0, 0),
// NewVertex(1, -1, 1, 1, 1),
// NewVertex(1, 1, 1, 1, 0),

// NewVertex(-1, -1, 1, 0, 1),
// NewVertex(1, -1, 1, 1, 1),
// NewVertex(-1, 1, 1, 0, 0),

// // Back

// NewVertex(1, 1, -1, 1, 0),
// NewVertex(1, -1, -1, 1, 1),
// NewVertex(-1, 1, -1, 0, 0),

// NewVertex(-1, 1, -1, 0, 0),
// NewVertex(1, -1, -1, 1, 1),
// NewVertex(-1, -1, -1, 0, 1),

// // Right

// NewVertex(1, 1, -1, 1, 0),
// NewVertex(1, 1, 1, 1, 1),
// NewVertex(1, -1, -1, 0, 0),

// NewVertex(1, -1, -1, 0, 0),
// NewVertex(1, 1, 1, 1, 1),
// NewVertex(1, -1, 1, 0, 1),

// // Left

// NewVertex(-1, -1, -1, 0, 0),
// NewVertex(-1, 1, 1, 1, 1),
// NewVertex(-1, 1, -1, 1, 0),

// NewVertex(-1, -1, 1, 0, 1),
// NewVertex(-1, 1, 1, 1, 1),
// NewVertex(-1, -1, -1, 0, 0),
