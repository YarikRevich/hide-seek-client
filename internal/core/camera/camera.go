package camera

import (
	"fmt"

	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	"github.com/hajimehoshi/ebiten/v2"
)

var instance *Camera

type Hero struct {
	attachedPos struct {
		X, Y float64
	}
	followed *objects.PC
	followedMatrix ebiten.GeoM
}

func (h *Hero) Follow(p *objects.PC){
	h.followed = p
}

func (h *Hero) GetMatrixFor(p *objects.PC) ebiten.GeoM{
	if p == h.followed{
		defer h.followedMatrix.Reset()
		return h.followedMatrix
	}
	g := ebiten.GeoM{}
	g.Scale(p.GetZoomForSkin())
	g.Translate(p.GetZoomedPos())
	return g
}

type Map struct {
	matrix ebiten.GeoM
}

func (m *Map) GetMatrix()ebiten.GeoM{
	w := objects.UseObjects().World()
	g := ebiten.GeoM{}
	g.Scale(w.GetMapScale())
	return g
} 

type Camera struct {
	Hero
	Map

	// mapAxis struct {
	// 	X, Y float64
	// }

	// lastMapScale struct {
	// 	X, Y float64
	// }

	// connectedHeroPos struct {
	// 	X, Y float64
	// }

	// lastConnectedHeroPos struct {
	// 	X, Y float64
	// }

	// connectedMapPos struct {
	// 	X, Y float64
	// }

	// lastConnectedMapPos struct {
	// 	X, Y float64
	// }

	// scaledConnectedHeroPos struct {
	// 	X, Y float64
	// }

	// //States if pc crossed x axis
	// isHeroMovementBlockedX bool

	// //States if pc crossed y axis
	// isHeroMovementBlockedY bool

	// scaledMapTranslation struct {
	// 	X, Y float64
	// }

	// lastScaledMapTranslation struct {
	// 	X, Y float64
	// }

	// scaledHeroTranslation struct {
	// 	X, Y float64
	// }

	// isScaledHeroTranslationBlockedX, isScaledHeroTranslationBlockedY bool

	// lastHeroScale struct {
	// 	X, Y float64
	// }

	// lastHeroTranslation struct {
	// 	X, Y float64
	// }
}

func (c *Camera) isHeroTranslationBlocked() bool {
	return c.isHeroMovementBlockedX || c.isHeroMovementBlockedY
}

//Updates general metrics for map matrix
func (c *Camera) updateMapMatrix() {
	c.MapMatrix.Scale(c.mapScale.X, c.mapScale.Y)

	p := objects.UseObjects().PC()
	w := objects.UseObjects().World()
	pm := p.GetMetadata().Modified
	wm := w.GetMetadata().Modified

	if c.isHeroTranslationBlocked() {
		if c.isHeroMovementBlockedX {
			// if (c.scaledHeroTranslation.X+c.scaledConnectedHeroPos.X+p.Metadata.Size.Width <= w.Metadata.Size.Width*c.mapScale.X) &&
			// 	(c.scaledHeroTranslation.X-c.scaledConnectedHeroPos.X) >= 0 {
			// fmt.Println(c.scaledHeroTranslation.X+c.scaledConnectedHeroPos.X+p.Metadata.Size.Width <= (w.Metadata.Size.Width * c.maxMapScale.X) - ((w.Metadata.Size.Width * c.mapScale.X) / c.maxMapScale.X))
			// fmt.Println(w.GetMapScale())
			fmt.Println(c.scaledHeroTranslation.X+c.scaledConnectedHeroPos.X+pm.Size.Width, (wm.Size.Width * c.maxMapScale.X) - ((wm.Size.Width * c.mapScale.X) / c.maxMapScale.X))
			//(1550 * 2.83) - ((1550 * 2.06) / 2.83)
			// fmt.Println(c.scaledHeroTranslation.X+c.scaledConnectedHeroPos.X+p.Metadata.Size.Width, w.Metadata.Size.Width, c.mapScale.X, c.maxMapScale.X)
			c.MapMatrix.Translate(-(c.scaledHeroTranslation.X - c.scaledConnectedHeroPos.X), 0)
			// } else {
				// cx := c.scaledHeroTranslation.X - c.scaledConnectedHeroPos.X
				// fmt.Println(cx)
				// if cx < 0 {
				// 	c.connectedMapPos.X = 0
				// } else {
				// 	mapWidth := w.Metadata.Size.Width * c.mapScale.X
				// 	if c.scaledHeroTranslation.X+c.scaledConnectedHeroPos.X >= mapWidth {
				// 		c.connectedMapPos.X = (cx - (mapWidth - (c.scaledHeroTranslation.X + c.scaledConnectedHeroPos.X))) - p.Metadata.Size.Width
				// 	} else {
				// 		c.connectedMapPos.X = cx
				// 	}
				// }
				// c.isHeroMovementBlockedX = false
				// c.isScaledHeroTranslationBlockedX = true
				// if c.scaledConnectedHeroPos.X != 0 {
				// 	c.scaledHeroTranslation.X = c.scaledConnectedHeroPos.X
				// }
				// c.connectedHeroPos.X = 0
			// }
		}

		if c.isHeroMovementBlockedY {
			// if (c.scaledHeroTranslation.Y+c.scaledConnectedHeroPos.Y+p.Metadata.Size.Height < w.Metadata.Size.Height*c.mapScale.Y) &&
				// (c.scaledHeroTranslation.Y-c.scaledConnectedHeroPos.Y) > 0 {
					// fmt.Println(c.scaledHeroTranslation.Y+c.scaledConnectedHeroPos.Y+p.Metadata.Size.Height < w.Metadata.Size.Height*c.mapScale.Y)
				c.MapMatrix.Translate(0, -(c.scaledHeroTranslation.Y - c.scaledConnectedHeroPos.Y))
			// } else {
				// mapHeight := w.Metadata.Size.Height * c.mapScale.Y
				// cy := c.scaledHeroTranslation.Y - c.scaledConnectedHeroPos.Y
				// if cy < 0 {
				// 	c.connectedMapPos.Y = 0
				// } else {
				// 	if c.scaledHeroTranslation.Y+c.scaledConnectedHeroPos.Y >= mapHeight {
				// 		c.connectedMapPos.Y = (cy - (mapHeight - (c.scaledHeroTranslation.Y + c.scaledConnectedHeroPos.Y))) - p.Metadata.Size.Height
				// 	} else {

				// 		c.connectedMapPos.Y = cy
				// 	}
				// }
				// c.isHeroMovementBlockedY = false
				// c.isScaledHeroTranslationBlockedY = true
				// if c.scaledConnectedHeroPos.Y != 0 {
				// 	c.scaledHeroTranslation.Y = c.scaledConnectedHeroPos.Y
				// }
				// c.connectedHeroPos.Y = 0
			// }
		}
	}

	// fmt.Println(c.connectedMapPos, c.scaledHeroTranslation, c.scaledConnectedHeroPos)

	// fmt.Println(c.connectedMapPos)

	if !c.isHeroMovementBlockedX && !c.isHeroMovementBlockedY {
		c.MapMatrix.Translate(-c.connectedMapPos.X, -c.connectedMapPos.Y)
	}

	if !c.isHeroMovementBlockedX && c.isHeroMovementBlockedY {
		c.MapMatrix.Translate(-c.connectedMapPos.X, 0)
	}

	if c.isHeroMovementBlockedX && !c.isHeroMovementBlockedY {
		c.MapMatrix.Translate(0, -c.connectedMapPos.Y)
	}
}

//Updates general metrics for hero matrix
func (c *Camera) updateHeroMatrix() {
	p := objects.UseObjects().PC()
	c.HeroMatrix.Scale(p.GetMovementRotation(), 1)
	c.HeroMatrix.Scale(c.heroScale.X, c.heroScale.Y)

	if !c.isHeroMovementBlockedX && c.isCrossedAxisX() {
		c.connectedHeroPos.X = p.RawPos.X
		c.isHeroMovementBlockedX = true
	}

	if !c.isHeroMovementBlockedY && c.isCrossedAxisY() {
		c.connectedHeroPos.Y = p.RawPos.Y
		c.isHeroMovementBlockedY = true
	}

	
	if !c.isHeroTranslationBlocked() {
		// fmt.Println(c.scaledHeroTranslation, c.isHeroMovementBlockedX, c.isHeroMovementBlockedY, c.mapScale.X, "GOOD")
		c.HeroMatrix.Translate(c.scaledHeroTranslation.X, c.scaledHeroTranslation.Y)
	} else {
		// fmt.Println(c.isHeroMovementBlockedX, c.isHeroMovementBlockedY, "BAD")
		// fmt.Println(c.scaledConnectedHeroPos, c.scaledConnectedHeroPos)
		if c.isHeroMovementBlockedX && c.isHeroMovementBlockedY {
			c.HeroMatrix.Translate(c.scaledConnectedHeroPos.X, c.scaledConnectedHeroPos.Y)
		}

		if c.isHeroMovementBlockedX && !c.isHeroMovementBlockedY {
			c.HeroMatrix.Translate(c.scaledConnectedHeroPos.X, c.scaledHeroTranslation.Y)
		}

		if !c.isHeroMovementBlockedX && c.isHeroMovementBlockedY {
			c.HeroMatrix.Translate(c.scaledHeroTranslation.X, c.scaledConnectedHeroPos.Y)
		}
	}
}

/*
Last property saves
*/

// func (c *Camera) updateLastHeroTranslation() {
// 	p := objects.UseObjects().PC()

// 	if c.lastHeroTranslation.X == 0 && c.lastHeroTranslation.Y == 0 {
// 		c.lastHeroTranslation.X = p.RawPos.X / (c.heroScale.X / c.lastHeroScale.X)
// 		c.lastHeroTranslation.Y = p.RawPos.Y / (c.heroScale.Y / c.lastHeroScale.Y)
// 		return
// 	} else {
// 		c.lastHeroTranslation.X = p.RawPos.X / (c.heroScale.X / c.lastHeroScale.X)
// 		c.lastHeroTranslation.Y = p.RawPos.Y / (c.heroScale.Y / c.lastHeroScale.Y)
// 	}

// 	if p.IsXChanged() || p.IsYChanged() || p.IsAnimatied() {
// 		c.lastHeroTranslation.X = p.RawPos.X / (c.heroScale.X / c.lastHeroScale.X)
// 		c.lastHeroTranslation.Y = p.RawPos.Y / (c.heroScale.Y / c.lastHeroScale.Y)
// 	}
// }

// func (c *Camera) updateLastScaledMapTranslation() {
// 	// if c.lastScaledMapTranslation.X == 0 && c.lastScaledMapTranslation.Y == 0 {
// 	// 	p := objects.UseObjects().PC()
// 	// 	c.scaledMapTranslation.X = p.RawPos.X * c.heroScale.X / c.maxHeroScale.X
// 	// 	c.scaledMapTranslation.Y = p.RawPos.Y * c.heroScale.Y / c.maxHeroScale.Y
// 	// } else {
// 	c.lastScaledMapTranslation = c.scaledMapTranslation
// 	// }
// }

/*
Update pipeline components
*/

// //Updates used scales for matrices
// func (c *Camera) updateScales() {
// 	// c.updateMapScale()
// 	// c.updateHeroScale()
// }

//Clears all the metrics for map and hero matrices
// func (c *Camera) clearMatrices() {
// 	// c.MapMatrix = ebiten.GeoM{}
// 	// c.HeroMatrix = ebiten.GeoM{}
// }

//Updates the properties used for other
//camera calculations
// func (c *Camera) updatePreDeps() {
	// c.updateMapAxis()

	// c.updateLastHeroTranslation()
	// c.updateScaledHeroTranslation()

	// c.updateLastHeroScale()

	// c.updateScaledMapTranslation()
	// c.updateLastScaledMapTranslation()

// }
// func (c *Camera) updatePostDeps() {
// 	c.updateScaledConnectedPos()
// }

//Updates used matrices
func (c *Camera) updateMatrices() {
	c.updateMapMatrix()
	c.updateHeroMatrix()
}

//Updates camera properties
func (c *Camera) UpdateCamera() {
	// c.updateScales()
	// c.clearMatrices()
	// c.updatePreDeps()
	c.updateMatrices()
	// c.updatePostDeps()
}

// func (c *Camera) SyncCameraPosition() {
// 	// w := world.UseWorld()

// 	// if c.Position.X == 0{
// 	// c.Position.X = w.Metadata.Size.Width / c.Scale.X / 2
// 	// }

// 	// fmt.Println(p.Y <= float64(c.MapBoarders.Max.Y), p.Y, float64(c.MapBoarders.Max.Y))

// 	fmt.Println(p.X, c.Position.X)
// 	// if p.X > c.Position.X {
// 	//w.Metadata.Size.Width / c.Scale.X / 2
// 	// c.Position.X++
// 	// }
// 	// c.Position.X += p.Buffs.SpeedX

// 	// fmt.Println()
// }

//Increments zoom property
func (c *Camera) ZoomIn() {
	w := objects.UseObjects().World()
	m := w.GetMetadata().Modified
	if m.Camera.Zoom <  m.Camera.MaxZoom{
		m.Camera.Zoom++
	}
}

//Decrements zoom property
func (c *Camera) ZoomOut() {
	w := objects.UseObjects().World()
	m := w.GetMetadata().Modified
	if m.Camera.Zoom > m.Camera.MinZoom {
		m.Camera.Zoom--
	}
}

//Uses or creates a new instance of camera
func UseCamera() *Camera {
	if instance == nil {
		instance = new(Camera)
	}
	return instance
}
