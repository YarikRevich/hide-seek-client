package collisions

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	// "github.com/YarikRevich/HideSeek-Client/internal/core/world"
)

//Returns the objects the object collides with
//if it collides though one object it will return ok as true
func IsPCCollideWithPCs(o objects.Base) ([]*objects.PC, bool) {
	// w := objects.UseObjects().World()
	// worldMap := world.UseWorld().GetWorldMap()
	// worldMap.ModelCombination.Modified.Size
	// m := w.GetMetadata().Modified
	// r := []*objects.PC{}
	// var ok bool

	// oMaxX := (m.Margins.LeftMargin + m.Size.Width)
	// oMinX := (m.Margins.LeftMargin)
	// oMaxY := (m.Margins.TopMargin + m.Size.Height)
	// oMinY := (m.Margins.TopMargin)
	// for _, v := range w.PCs {
	// 	vMaxX := (m.Margins.LeftMargin + m.Size.Width)
	// 	vMinX := (m.Margins.LeftMargin)
	// 	vMaxY := (m.Margins.TopMargin + m.Size.Height)
	// 	vMinY := (m.Margins.TopMargin)

	// 	if (oMinX <= vMaxX && vMinX <= oMaxX) ||
	// 		(oMinY <= vMaxY && vMinY <= oMaxY) {
	// 		r = append(r, v)
	// 		ok = true
	// 	}
	// }
	// return r, ok
	return nil, false
}

//Returns the objects the object collides with
//if it collides though one object it will return ok as true
func IsAmmoCollideWithObject(o objects.Weapon) ([]*objects.Weapon, bool) {
	// w := objects.UseObjects().World()
	// r := []*{}
	// var ok bool

	// oMaxX := (m.Margins.LeftMargin + m.Size.Width)
	// oMinX := (m.Margins.LeftMargin)
	// oMaxY := (m.Margins.TopMargin + m.Size.Height)
	// oMinY := (m.Margins.TopMargin)
	// for _, v := range w.Ammo {
	// 	vMaxX := (m.Margins.LeftMargin + m.Size.Width)
	// 	vMinX := (m.Margins.LeftMargin)
	// 	vMaxY := (m.Margins.TopMargin + m.Size.Height)
	// 	vMinY := (m.Margins.TopMargin)

	// 	if (oMinX <= vMaxX && vMinX <= oMaxX) &&
	// 		(oMinY <= vMaxY && vMinY <= oMaxY) {
	// 		r = append(r, v)
	// 		ok = true
	// 	}
	// }
	// return r, ok
	return nil, false
}
