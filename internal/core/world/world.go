package world

import (
	"fmt"

	"github.com/YarikRevich/HideSeek-Client/internal/core/gamesettings"
	"github.com/YarikRevich/HideSeek-Client/internal/core/networking/api"
	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	"github.com/google/uuid"
)

var instance *World

type World struct {
	ID uuid.UUID

	gamesettings gamesettings.GameSettings
	worldMap     *objects.Map
	pc           *objects.PC
	camera       *objects.Camera

	pcs      []*objects.PC
	elements []*objects.Element
	weapons  []*objects.Weapon
	ammos    []*objects.Ammo
}

func (w *World) AddPCs(pc *objects.PC) {
	pc.Parent = &w.worldMap.Base
	w.pcs = append(w.pcs, pc)
}

func (w *World) DeletePCs() {
	w.pcs = w.pcs[:0]
}

func (w *World) UpdatePCs(m []*api.PC) {
	w.DeletePCs()

	for _, pc := range m {
		if pc.Base.Id == w.pc.ID.String() {
			w.pc.FromAPIMessage(pc)
			w.pcs = append(w.pcs, w.pc)
		} else {
			npc := objects.NewPC()
			npc.FromAPIMessage(pc)
			w.pcs = append(w.pcs, npc)
		}
	}
}

func (w *World) AddElements(el *objects.Element) {
	el.Parent = &w.worldMap.Base
	w.elements = append(w.elements, el)
}

func (w *World) DeleteElements() {
	w.elements = w.elements[:0]
}

func (w *World) UpdateElements(m []*api.Element) {
	w.DeleteElements()
}

func (w *World) AddWeapons(we *objects.Weapon) {
	we.Parent = &w.worldMap.Base
	w.weapons = append(w.weapons, we)
}

func (w *World) DeleteWeapons() {
	w.weapons = w.weapons[:0]
}

func (w *World) UpdateWeapons(m []*api.Weapon) {
	w.DeleteWeapons()
}

func (w *World) AddAmmos(am *objects.Ammo) {
	am.Parent = &w.worldMap.Base
	w.ammos = append(w.ammos, am)
}

func (w *World) DeleteAmmos() {
	w.ammos = w.ammos[:0]
}

func (w *World) UpdateAmmos(m []*api.Ammo) {
	w.DeleteAmmos()
}

func (w *World) UpdateProperty(m *api.WorldObjectsResponse) {
	w.UpdatePCs(m.PCs)
	w.UpdateElements(m.Elements)
	w.UpdateWeapons(m.Weapons)
	w.UpdateAmmos(m.Ammos)
}

func (w *World) SwapPCsSpawns() {

	//Swaps spawns of the teams
	// func (w *Map) SwapSpawns() {
	// map[pc.Team]map[uuid.UUID]image.Point{}
	// newTeam1Swaps := map[uuid.UUID]image.Point{}
	// newTeam2Swaps := map[uuid.UUID]image.Point{}
	// for _, u := range w.PCs {
	// 	// switch u.Team {
	// 	// case pc.Team1:
	// 	// 	newTeam2Swaps[u.ID] = u.Spawn
	// 	// case pc.Team2:
	// 	// 	newTeam1Swaps[u.ID] = u.Spawn
	// 	// }
	// 	fmt.Println(u)
	// }

	// for _, u := range w.PCs {
	// 	// switch u.Team {
	// 	// case pc.Team1:
	// 	// 	u.Spawn = newTeam2Swaps[u.ID]
	// 	// case pc.Team2:
	// 	// 	u.Spawn = newTeam1Swaps[u.ID]
	// 	// }
	// 	fmt.Println(u)
	// }
	// }
}

func (w *World) GetWorldMap() *objects.Map {
	return w.worldMap
}

func (w *World) GetPC() *objects.PC {
	return w.pc
}

func (w *World) GetPCs() []*objects.PC {
	return w.pcs
}

func (w *World) GetCamera() *objects.Camera {
	return w.camera
}

func (w *World) GetWeapons() []*objects.Weapon {
	return w.weapons
}

func (w *World) GetAmmos() []*objects.Ammo {
	return w.ammos
}

func (w *World) ToAPIMessage() *api.World{
	return &api.World{
		Id: w.ID.String(), 
		GameSettings: w.gamesettings.ToAPIMessage(),
	}
}

func (w *World) FromAPIMessage(m *api.World){
	w.gamesettings.Regime = m.GameSettings.Regime
}

func (w *World) SetID(id uuid.UUID){
	w.ID = id
}

// //Formats users' username
func (w *World) String() string {
	var r string
	for _, v := range w.pcs {
		r += fmt.Sprintf("%s\n", v.String())
	}
	return r
}

func UseWorld() *World {
	if instance == nil {
		instance = &World{
			pc:       objects.NewPC(),
			worldMap: objects.NewMap(),
			camera:   objects.NewCamera(),
		}

		instance.pc.Parent = &objects.Base{ID: instance.ID}
		instance.worldMap.Parent = &objects.Base{ID: instance.ID}
		instance.camera.Parent = &objects.Base{ID: instance.ID}

		instance.AddPCs(instance.pc)
	}
	return instance
}
