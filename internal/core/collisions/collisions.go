package collisions

import "github.com/YarikRevich/HideSeek-Client/internal/core/objects"

//Returns the objects the object collides with
//if it collides though one object it will return ok as true
func IsPCCollideWithPCs(o objects.Object) ([]*objects.PC, bool) {
	w := objects.UseObjects().World()
	r := []*objects.PC{}
	var ok bool

	oMaxX := (o.Metadata.Margins.LeftMargin + o.Metadata.Size.Width)
	oMinX := (o.Metadata.Margins.LeftMargin)
	oMaxY := (o.Metadata.Margins.TopMargin + o.Metadata.Size.Height)
	oMinY := (o.Metadata.Margins.TopMargin)
	for _, v := range w.PCs {
		vMaxX := (v.Metadata.Margins.LeftMargin + v.Metadata.Size.Width)
		vMinX := (v.Metadata.Margins.LeftMargin)
		vMaxY := (v.Metadata.Margins.TopMargin + v.Metadata.Size.Height)
		vMinY := (v.Metadata.Margins.TopMargin)

		if (oMinX <= vMaxX && vMinX <= oMaxX) &&
			(oMinY <= vMaxY && vMinY <= oMaxY) {
			r = append(r, v)
			ok = true
		}
	}
	return r, ok
}

//Returns the objects the object collides with
//if it collides though one object it will return ok as true
func IsAmmoCollideWithObject(o objects.Weapon) ([]*objects.Weapon, bool) {
	w := objects.UseObjects().World()
	r := []*objects.PC{}
	var ok bool

	oMaxX := (o.Metadata.Margins.LeftMargin + o.Metadata.Size.Width)
	oMinX := (o.Metadata.Margins.LeftMargin)
	oMaxY := (o.Metadata.Margins.TopMargin + o.Metadata.Size.Height)
	oMinY := (o.Metadata.Margins.TopMargin)
	for _, v := range w {
		vMaxX := (v.Metadata.Margins.LeftMargin + v.Metadata.Size.Width)
		vMinX := (v.Metadata.Margins.LeftMargin)
		vMaxY := (v.Metadata.Margins.TopMargin + v.Metadata.Size.Height)
		vMinY := (v.Metadata.Margins.TopMargin)

		if (oMinX <= vMaxX && vMinX <= oMaxX) &&
			(oMinY <= vMaxY && vMinY <= oMaxY) {
			r = append(r, v)
			ok = true
		}
	}
	return r, ok
}
