package world

import (
	"fmt"

	"github.com/YarikRevich/hide-seek-client/internal/core/camera"
	"github.com/YarikRevich/hide-seek-client/internal/core/events"
	"github.com/YarikRevich/hide-seek-client/internal/core/gamesettings"
	"github.com/YarikRevich/hide-seek-client/internal/core/keycodes"
	"github.com/YarikRevich/hide-seek-client/internal/core/networking/api/server_external"
	"github.com/YarikRevich/hide-seek-client/internal/core/objects"
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"
	"github.com/YarikRevich/hide-seek-client/internal/core/statistics"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
)

var instance *World

type World struct {
	ID uuid.UUID

	gamesettings gamesettings.GameSettings
	worldMap     *objects.Map
	pc           *objects.PC

	pcs      []*objects.PC
	elements []*objects.Element
	weapons  []*objects.Weapon
	ammos    []*objects.Ammo

	statistics *statistics.Statistics
}

func (w *World) AddPCs(pc *objects.PC) {
	w.pcs = append(w.pcs, pc)
}

func (w *World) DeletePCs() {
	w.pcs = w.pcs[:0]
}

func (w *World) UpdateWorld(m *server_external.World) {
	w.FromAPIMessage(m)
}

func (w *World) UpdateWorldMap(m *server_external.Map) {
	w.worldMap.FromAPIMessage(m)
}

func (w *World) UpdatePCs(m []*server_external.PC) {
	w.DeletePCs()

	var foundPC bool
	for _, pc := range m {
		if pc.Base.Id == w.pc.ID.String() {
			pc.Base.Parent = w.worldMap.Base.ToAPIMessage()
			w.pc.FromAPIMessage(pc)
			foundPC = true
			w.AddPCs(w.pc)
		} else {
			npc := objects.NewPC(objects.PCOpts{Health: 10})
			npc.FromAPIMessage(pc)
			w.AddPCs(npc)
		}
	}
	if !foundPC {
		// middlewares.UseMiddlewares().UI().UseAfter(func() {
		statemachine.Layers.SetState(statemachine.LAYERS_START_MENU)
		// })

		// notifications.PopUp.WriteError("You were kicked from the session")
	}
}

func (w *World) AddElements(el *objects.Element) {
	el.Parent = &w.worldMap.Base
	w.elements = append(w.elements, el)
}

func (w *World) DeleteElements() {
	w.elements = w.elements[:0]
}

func (w *World) UpdateElements(m []*server_external.Element) {
	w.DeleteElements()
}

func (w *World) AddWeapons(we *objects.Weapon) {
	we.Parent = &w.worldMap.Base
	w.weapons = append(w.weapons, we)
}

func (w *World) DeleteWeapons() {
	w.weapons = w.weapons[:0]
}

func (w *World) UpdateWeapons(m []*server_external.Weapon) {
	w.DeleteWeapons()
}

func (w *World) AddAmmos(am *objects.Ammo) {
	am.Parent = &w.worldMap.Base
	w.ammos = append(w.ammos, am)
}

func (w *World) DeleteAmmos() {
	w.ammos = w.ammos[:0]
}

func (w *World) UpdateAmmos(m []*server_external.Ammo) {
	w.DeleteAmmos()
}

func (w *World) Update(m *server_external.FindWorldObjectsResponse) {
	w.UpdateWorld(m.World)
	w.UpdateWorldMap(m.WorldMap)
	// fmt.Println(m.WorldMap, "\n\n", w.worldMap, "\n\n")
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

func (w *World) GetGameSettings() *gamesettings.GameSettings {
	return &w.gamesettings
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

func (w *World) GetWeapons() []*objects.Weapon {
	return w.weapons
}

func (w *World) GetAmmos() []*objects.Ammo {
	return w.ammos
}

func (w *World) GetStatistics() *statistics.Statistics {
	return w.statistics
}

func (w *World) ToAPIMessage() *server_external.World {
	return &server_external.World{
		Id:           w.ID.String(),
		GameSettings: w.gamesettings.ToAPIMessage(),
	}
}

func (w *World) FromAPIMessage(m *server_external.World) {
	// w.gamesettings.Regime = m.GameSettings.Regime
	// fmt.Println(m)
	// fmt.Println(m.GameSettings)
	w.gamesettings.IsGameStarted = m.GameSettings.IsGameStarted
	w.gamesettings.IsWorldExist = m.GameSettings.IsWorldExist
}

func (w *World) SetID(id uuid.UUID) {
	w.ID = id
	w.pc.Parent.ID = id
	w.worldMap.Parent.ID = id
	for _, pc := range w.pcs {
		pc.Parent.ID = id
	}

	for _, el := range w.elements {
		el.Parent.ID = id
	}

	for _, we := range w.weapons {
		we.Parent.ID = id
	}

	for _, am := range w.ammos {
		am.Parent.ID = id
	}
}

// //Formats users' username
func (w *World) String() string {
	var r string
	for _, v := range w.pcs {
		r += fmt.Sprintf("%s\n", v.String())
	}
	return r
}

// //Inits basic world for debug purpose
// func (w *World) DebugInit() {
// 	w.pc.DebugInit()
// 	w.worldMap.DebugInit()
// }

//Checks if x and y pos collide with objects in the world
// func (w *World) IsCollision(x, y float64) (bool, int) {
// for _, v := range w.pcs {
// 	if v.RawPos.X == x && v.RawPos.Y == y {
// 		return true, objects.PLAYER
// 	}
// }

// for _, v := range w.weapons {
// 		if v.RawPos.X == x && v.RawPos.Y == y {
// 			return true, objects.WEAPON
// 		}
// 	}

// 	for _, v := range w.ammos {
// 		if v.RawPos.X == x && v.RawPos.Y == y {
// 			return true, objects.AMMO
// 		}
// 	}

// 	for _, v := range w.elements {
// 		if v.RawPos.X == x && v.RawPos.Y == y {
// 			return true, objects.ELEMENT
// 		}
// 	}
// 	return false, 0
// }

func UseWorld() *World {
	if instance == nil {
		instance = &World{
			ID:       uuid.New(),
			pc:       objects.NewPC(objects.PCOpts{Health: 10}),
			worldMap: objects.NewMap(),
		}

		instance.worldMap.Parent = &objects.Base{ID: instance.ID}
		instance.pc.Parent = &instance.worldMap.Base

		instance.AddPCs(instance.pc)

		instance.gamesettings.SetWorldExist(true)
	}
	return instance
}

type WorldManager struct {
	objects []*objects.Base

	ID uuid.UUID

	gamesettings gamesettings.GameSettings
	worldMap     *objects.Map
	pc           *objects.PC

	pcs      []*objects.PC
	elements []*objects.Element
	weapons  []*objects.Weapon
	ammos    []*objects.Ammo

	statistics *statistics.Statistics

	Camera *camera.Camera
}

//Creates snapshot of world and sends it to the server
func (wm *WorldManager) ToAPIMessage() *server_external.World {
	return &server_external.World{
		Id:           wm.ID.String(),
		GameSettings: wm.gamesettings.ToAPIMessage(),
	}
}

func (wm *WorldManager) FromAPIMessage(m *server_external.World) {
	//TODO: implementing importing of map regime
	// w.gamesettings.Regime = m.GameSettings.Regime
	wm.gamesettings.IsGameStarted = m.GameSettings.IsGameStarted
	wm.gamesettings.IsWorldExist = m.GameSettings.IsWorldExist
}

func (wm *WorldManager) Update() {
	if events.KeyboardPress.IsAnyKeyPressed() || events.GamepadPress.IsAnyButtonPressed() {
		if events.GamepadPress.AreGamepadButtonsCombined(keycodes.GamepadUPButton, keycodes.GamepadLEFTUPPERCLICKERButton) || ebiten.IsKeyPressed(ebiten.KeyF1) {
			wm.Camera.ZoomIn(-0.4)
		} else if events.GamepadPress.AreGamepadButtonsCombined(keycodes.GamepadDOWNButton, keycodes.GamepadLEFTUPPERCLICKERButton) || ebiten.IsKeyPressed(ebiten.KeyF2) {
			wm.Camera.ZoomIn(0.4)
		}

		// if ebiten.IsKeyPressed(ebiten.Key8) {
		// 	wm.Camera.MoveAngle(0.04)
		// }

		// if ebiten.IsKeyPressed(ebiten.Key7) {
		// 	wm.Camera.MoveAngle(-0.04)
		// }

		// if ebiten.IsKeyPressed(ebiten.Key6) {
		// 	wm.Camera.MovePitch(0.04)
		// }

		// if ebiten.IsKeyPressed(ebiten.Key5) {
		// 	wm.Camera.MovePitch(-0.04)
		// }

		if ebiten.IsKeyPressed(ebiten.KeyA) {
			wm.Camera.MovePositionX(-2)
		}

		if ebiten.IsKeyPressed(ebiten.KeyD) {
			wm.Camera.MovePositionX(2)
		}

		if ebiten.IsKeyPressed(ebiten.KeyW) {
			wm.Camera.MovePositionY(-2)
		}

		if ebiten.IsKeyPressed(ebiten.KeyS) {
			wm.Camera.MovePositionY(2)
		}
	}
}

func (wm *WorldManager) Render(sm *screen.ScreenManager) {
	// wm.camera.Opts.Angle
	// sm.GetImage().DrawImage()
}

func NewWorldManager() *WorldManager {
	return &WorldManager{Camera: &camera.Camera{Pitch: 0.5, Angle: 5.48, Zoom: 1, Position: types.Vec3{10, 10, 10}}}
}
