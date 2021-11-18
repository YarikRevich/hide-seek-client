package objects

import (
	"fmt"
	"github.com/YarikRevich/HideSeek-Client/internal/core/networking/api"
	"github.com/YarikRevich/HideSeek-Client/internal/core/screen"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type regime int

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

type World struct {
	Object

	Regime regime

	//Describes the objects are on the map
	PCs      []*PC
	Elements []*Object
	Weapons  []*Weapon
	Ammo     []*Ammo
	// LootSet []*LootSet
}

func (w *World) AddPC(p *PC) {
	p.WorldID = w.ID
	w.PCs = append(w.PCs, p)
}

func (w *World) AddWeapon(p *Weapon) {
	p.WorldID = w.ID
	w.Weapons = append(w.Weapons, p)
}

func (w *World) AddAmmo(a *Ammo) {
	a.WorldID = w.ID
	w.Ammo = append(w.Ammo, a)
}

func (w *World) GetPCs() []*PC {
	avoidID := UseObjects().PC().ID
	r := make([]*PC, 0, 3)
	for _, v := range w.PCs {
		if v.ID != avoidID {
			r = append(r, v)
		}
	}
	return r
}

func (w *World) deleteObjects() {
	w.Ammo = w.Ammo[:0]
	w.Elements = w.Elements[:0]
	w.Weapons = w.Weapons[:0]
	w.PCs = w.PCs[:0]
}

func (w *World) updatePCs(m []*api.PC) {
	pc := UseObjects().PC()
	for _, v := range m {
		if v.Object.Id == pc.ID.String() {
			pc.FromAPIMessage(v)
			w.PCs = append(w.PCs, pc)
		} else {
			np := NewPC()
			np.FromAPIMessage(v)
			w.PCs = append(w.PCs, np)
		}
	}
}

func (w *World) UpdateObjects(m *api.WorldObjectsResponse) {
	w.deleteObjects()

	w.updatePCs(m.PCs)
}

func (w *World) GetWeaponByPC(p *PC) *Weapon {
	for _, v := range w.Weapons {
		if v.ParentID == p.ID {
			return v
		}
	}
	return nil
}

func (w *World) GetAmmoByWeapon(p *Weapon) *Ammo {
	if p == nil {
		return nil
	}
	for _, v := range w.Ammo {
		if v.ParentID == p.ID {
			return v
		}
	}
	return nil
}

//Resets the list of users on the map
func (w *World) ResetPCs() {
	w.PCs = w.PCs[:0]
}

// //Formats users' username
func (w *World) PCsToString() string {
	var r string
	for _, v := range w.PCs {
		r += fmt.Sprintf("%s\n", v.String())
	}
	return r
}

func (w *World) SetID(i uuid.UUID) {
	w.ID = i
}

func (w *World) GetMapScale() (float64, float64) {
	var sx, sy float64
	s := screen.UseScreen()
	screenW := s.GetWidth()
	screenH := s.GetHeight()

	m := w.GetMetadata().Origin

	if screenW > m.Size.Width {
		sx = m.Size.Width / screenW
	} else {
		sx = screenW / m.Size.Width
	}

	if screenH > m.Size.Height {
		sy = m.Size.Height / screenH
	} else {
		sy = screenH / m.Size.Height
	}
	return sx, sy
}

func (w *World) GetZoomedMapScale() (float64, float64) {
	m := w.GetMetadata().Modified
	sx, sy := w.GetMapScale()
	return ((sx + m.Scale.CoefficiantX) / 100 * m.Camera.Zoom * 3), ((sy + m.Scale.CoefficiantY) / 100 * m.Camera.Zoom * 3)
}

// func (w *World) GetZoomedMapScaleManually(zoom float64)(float64, float64){
// 	m := w.GetMetadata().Modified
// 	sx, sy := w.GetMapScale()
// 	return ((sx + m.Scale.CoefficiantX) / 100 * zoom * 3), ((sy + m.Scale.CoefficiantY) / 100 * zoom * 3)
// }

func (w *World) GetZoomedMaxMapScale() (float64, float64) {
	m := w.GetMetadata().Origin
	sx, sy := w.GetZoomedMapScale()
	return ((sx + m.Scale.CoefficiantX) / 100 * m.Camera.MaxZoom * 3), ((sy + m.Scale.CoefficiantY) / 100 * m.Camera.MaxZoom * 3)
}

// func (w *World) GetWorldAxis() (float64, float64) {
// 	x, y := screen.UseScreen().GetScreen().Size()
// 	return float64(x) / 2, float64(y) / 2
// }

func (w *World) IsAxisXCrossedBy(p *PC) bool {
	mm := p.GetMetadata().Modified
	mo := p.GetMetadata().Origin
	x, _ := p.GetZoomedRawPosForCamera(w.GetZoomedMapScale())
	// ax, _ := w.GetWorldAxis()
	ax := screen.UseScreen().GetAxisX()

	return (x-mm.Buffs.Speed.X - mo.Size.Width/2) <= ax && ax <= (x+mm.Buffs.Speed.X + mo.Size.Width/2)
}

func (w *World) IsAxisYCrossedBy(p *PC) bool {
	mm := p.GetMetadata().Modified
	mo := p.GetMetadata().Origin
	_, y := p.GetZoomedRawPosForCamera(w.GetZoomedMapScale())
	ay := screen.UseScreen().GetAxisY()

	return (y-mm.Buffs.Speed.Y - mo.Size.Height/2) <= ay && ay <= (y+mm.Buffs.Speed.Y + mo.Size.Height/2)
}

//Swaps spawns of the teams
func (w *World) SwapSpawns() {
	// map[pc.Team]map[uuid.UUID]image.Point{}
	// newTeam1Swaps := map[uuid.UUID]image.Point{}
	// newTeam2Swaps := map[uuid.UUID]image.Point{}
	for _, u := range w.PCs {
		// switch u.Team {
		// case pc.Team1:
		// 	newTeam2Swaps[u.ID] = u.Spawn
		// case pc.Team2:
		// 	newTeam1Swaps[u.ID] = u.Spawn
		// }
		fmt.Println(u)
	}

	for _, u := range w.PCs {
		// switch u.Team {
		// case pc.Team1:
		// 	u.Spawn = newTeam2Swaps[u.ID]
		// case pc.Team2:
		// 	u.Spawn = newTeam1Swaps[u.ID]
		// }
		fmt.Println(u)
	}
}

func (w *World) ToAPIMessage() *api.World {
	return &api.World{
		Object: w.Object.ToAPIMessage(),
		Regime: int64(w.Regime),
	}
}

func (w *World) FromAPIMessage(m *api.World) {
	w.Object.FromAPIMessage(m.Object)
	w.Regime = regime(m.Regime)
}

func (w *World) GetScaleForSkin() (float64, float64) {
	m := w.GetMetadata().Modified
	sx, sy := w.GetMapScale()
	return ((sx + m.Scale.CoefficiantX) / 100 * m.Camera.Zoom * 3), ((sy + m.Scale.CoefficiantY/100) * m.Camera.Zoom * 3)
}

func (w *World) GetMaxScaleForSkin() (float64, float64) {
	m := w.GetMetadata().Modified
	sx, sy := w.GetMapScale()
	return ((sx + m.Scale.CoefficiantX) / 100 * m.Camera.MaxZoom * 3), ((sy + m.Scale.CoefficiantY/100) * m.Camera.MaxZoom * 3)
}

func (w *World) GetZoomedAttachedPos() (float64, float64) {
	mapScaleX, mapScaleY := UseObjects().World().GetZoomedMapScale()
	return  w.AttachedPos.X * mapScaleX, w.AttachedPos.Y * mapScaleY
}

// func (w *World) GetScaledPos(){
// c.scaledMapTranslation.X = c.lastScaledMapTranslation.X * c.mapScale.X / c.lastMapScale.X
// c.scaledMapTranslation.Y = c.lastScaledMapTranslation.Y * c.mapScale.Y / c.lastMapScale.Y
// }

func NewWorld() *World {
	world := new(World)
	id, err := uuid.NewUUID()
	if err != nil {
		logrus.Fatal("failed to create uuid for world:", err)
	}
	world.ID = id
	return world
}
