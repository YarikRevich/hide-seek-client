package objects

import (
	"image"
	"unsafe"

	"github.com/YarikRevich/hide-seek-client/internal/core/networking/api/server_external"
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// const (
// 	//Players are separated on two teams
// 	//If player dies it will respawn
// 	//If the biggest part of the players is
// 	//on the oposite teritory, spawns will be swapped
// 	deathmatch regime = iota

// 	//Players are separated on two teams
// 	//The game will end if all members of the team
// 	//will dir
// 	teamToTeam
// )

type Map struct {
	Base

	// Regime regime

	// //Describes the objects are on the map
	// PCs      []*PC
	// Elements []*Object
	// Weapons  []*Weapon
	// Ammo     []*Ammo
	// LootSet []*LootSet
}

// func (w *World) AddPC(p *PC) {
// 	p.Parent = &w.Object
// 	w.PCs = append(w.PCs, p)
// }

// func (w *World) AddWeapon(p *Weapon) {
// 	p.Parent = &w.Object
// 	w.Weapons = append(w.Weapons, p)
// }

// func (w *World) AddAmmo(a *Ammo) {
// 	a.Parent = &w.Object
// 	w.Ammo = append(w.Ammo, a)
// }

// func (w *World) GetPCs() []*PC {
// 	avoidID := UseObjects().PC().ID
// 	r := make([]*PC, 0, 3)
// 	for _, v := range w.PCs {
// 		if v.ID != avoidID {
// 			r = append(r, v)
// 		}
// 	}
// 	return r
// }

// func (w *World) deleteObjects() {
// 	w.Ammo = w.Ammo[:0]
// 	w.Elements = w.Elements[:0]
// 	w.Weapons = w.Weapons[:0]
// 	w.PCs = w.PCs[:0]
// }

// func (w *World) updatePCs(m []*api.PC) {
// 	pc := UseObjects().PC()
// 	for _, v := range m {
// 		if v.Object.Id == pc.ID.String() {
// 			pc.Parent = &w.Object
// 			pc.FromAPIMessage(v)
// 			w.PCs = append(w.PCs, pc)
// 		} else {
// 			np := NewPC()
// 			np.Parent = &w.Object
// 			np.FromAPIMessage(v)
// 			w.PCs = append(w.PCs, np)
// 		}
// 	}
// }

// func (w *World) UpdateObjects(m *api.WorldObjectsResponse) {
// 	w.deleteObjects()

// 	w.updatePCs(m.PCs)
// }

// func (w *World) GetWeaponByPC(p *PC) *Weapon {
// 	for _, v := range w.Weapons {
// 		if v.Parent.ID == p.ID {
// 			return v
// 		}
// 	}
// 	return nil
// }

// func (w *World) GetAmmoByWeapon(p *Weapon) *Ammo {
// 	if p == nil {
// 		return nil
// 	}
// 	for _, v := range w.Ammo {
// 		if v.Parent.ID == p.ID {
// 			return v
// 		}
// 	}
// 	return nil
// }
// func (w *World) SetID(i uuid.UUID) {
// 	w.ID = i
// }

// func (w *World) GetZoomedMapScale() (float64, float64) {
// 	m := w.GetMetadata().Modified
// 	sx, sy := w.GetMapScale()
// 	return ((sx + m.Scale.CoefficiantX) / 100 * m.Camera.Zoom * 3), ((sy + m.Scale.CoefficiantY) / 100 * m.Camera.Zoom * 3)
// }

// func (w *World) GetZoomedMapScaleManually(zoom float64)(float64, float64){
// 	m := w.GetMetadata().Modified
// 	sx, sy := w.GetMapScale()
// 	return ((sx + m.Scale.CoefficiantX) / 100 * zoom * 3), ((sy + m.Scale.CoefficiantY) / 100 * zoom * 3)
// }

// func (w *World) GetZoomedMaxMapScale() (float64, float64) {
// 	m := w.GetMetadata().Origin
// 	sx, sy := w.GetZoomedMapScale()
// 	return ((sx + m.Scale.CoefficiantX) / 100 * m.Camera.MaxZoom * 3), ((sy + m.Scale.CoefficiantY) / 100 * m.Camera.MaxZoom * 3)
// }

// func (w *World) GetWorldAxis() (float64, float64) {
// 	x, y := screen.UseScreen().GetScreen().Size()
// 	return float64(x) / 2, float64(y) / 2
// }

// func (w *World) IsAxisXCrossedBy(p *PC) bool {
// 	mm := p.GetMetadata().Modified
// 	mo := p.GetMetadata().Origin
// 	x, _ := p.GetZoomedRawPosForCamera(w.GetZoomedMapScale())
// 	// ax, _ := w.GetWorldAxis()
// 	ax := screen.UseScreen().GetAxisX()

// 	return (x-mm.Buffs.Speed.X - mo.Size.Width/2) <= ax && ax <= (x+mm.Buffs.Speed.X + mo.Size.Width/2)
// }

// func (w *World) IsAxisYCrossedBy(p *PC) bool {
// 	mm := p.GetMetadata().Modified
// 	mo := p.GetMetadata().Origin
// 	_, y := p.GetZoomedRawPosForCamera(w.GetZoomedMapScale())
// 	ay := screen.UseScreen().GetAxisY()

// 	return (y-mm.Buffs.Speed.Y - mo.Size.Height/2) <= ay && ay <= (y+mm.Buffs.Speed.Y + mo.Size.Height/2)
// }
func (w *Map) GetSpawns() []*image.Point {
	var r []*image.Point

	hudOffsetY := screen.UseScreen().GetHeight() / 12
	for _, v := range w.ModelCombination.Modified.Spawns {
		r = append(r, &image.Point{Y: v.Y + int(hudOffsetY), X: v.X})
	}

	return r
}

func (w *Map) ToAPIMessage() *server_external.Map {
	return &server_external.Map{
		Base:   w.Base.ToAPIMessage(),
		Spawns: *(*[]*server_external.PositionInt)(unsafe.Pointer(&w.Base.ModelCombination.Origin.Spawns)),
	}
}

func (w *Map) FromAPIMessage(m *server_external.Map) {
	w.Base.FromAPIMessage(m.Base)
}

func (m *Map) DebugInit() {
	m.Base.SetSkin("maps/helloween/background/background")
}

// func (w *World) GetScaleForSkin() (float64, float64) {
// 	m := w.GetMetadata().Modified
// 	sx, sy := w.GetMapScale()
// 	return ((sx + m.Scale.CoefficiantX) / 100 * m.Camera.Zoom * 3), ((sy + m.Scale.CoefficiantY/100) * m.Camera.Zoom * 3)
// }

// func (w *World) GetMaxScaleForSkin() (float64, float64) {
// 	m := w.GetMetadata().Modified
// 	sx, sy := w.GetMapScale()
// 	return ((sx + m.Scale.CoefficiantX) / 100 * m.Camera.MaxZoom * 3), ((sy + m.Scale.CoefficiantY/100) * m.Camera.MaxZoom * 3)
// }

// func (w *World) GetZoomedAttachedPos() (float64, float64) {
// 	return  w.AttachedPos.X * w.ModelCombination.Modified.Scale.X, w.AttachedPos.Y * w.ModelCombination.Modified.Scale.Y
// }

// func (w *World) GetScaledPos(){
// c.scaledMapTranslation.X = c.lastScaledMapTranslation.X * c.mapScale.X / c.lastMapScale.X
// c.scaledMapTranslation.Y = c.lastScaledMapTranslation.Y * c.mapScale.Y / c.lastMapScale.Y
// }

func NewMap() *Map {
	world := new(Map)
	id, err := uuid.NewUUID()
	if err != nil {
		logrus.Fatal("failed to create uuid for world:", err)
	}
	world.ID = id
	return world
}
