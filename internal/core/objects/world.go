package objects

import (
	"fmt"
	// "image"
	// "strings"

	// "github.com/YarikRevich/HideSeek-Client/internal/gameplay/objects"
	// "github.com/YarikRevich/HideSeek-Client/internal/gameplay/pc"
	"github.com/YarikRevich/HideSeek-Client/internal/core/screen"
	"github.com/google/uuid"
	// "github.com/hajimehoshi/ebiten/v2"
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
	PCs []*PC
	Elements []*Object
	Weapons []*Weapon
	Ammo []*Ammo
	LootSet []*LootSet
}

func (w *World) GetWeaponByPC(p *PC)*Weapon{
	for _, v := range w.Weapons{
		if v.ParentID == p.ID{
			return v
		}
	}
	return nil
}

func (w *World) GetAmmoByWeapon(p *Weapon)*Ammo{
	if p == nil{
		return nil
	}
	for _, v := range w.Ammo{
		if v.ParentID == p.ID{
			return v
		}
	}
	return nil
}


//Resets the list of users on the map
func (w *World) ResetPCs() {
	w.PCs = w.PCs[:0]
}

//Formats users' username
func (w *World) String() string {
	var r string
	for _, v := range w.PCs {
		r += fmt.Sprintf("%s\n", v.Username)
	}
	return r
}

//Returns map scale in relating map image
//to current screen sizes
func (w *World) GetMaxMapScale() (float64, float64) {
	var sx, sy float64 
	screenW := float64(screen.GetMaxWidth())
	screenH := float64(screen.GetMaxHeight())
	
	if screenW > w.Metadata.RawSize.Width {
		sx = w.Metadata.RawSize.Width / screenW
	} else {
		sx = screenW / w.Metadata.RawSize.Width
	}

	if screenH > w.Metadata.RawSize.Height {
	sy = w.Metadata.RawSize.Height / screenH
	} else {
		sy = screenH / w.Metadata.RawSize.Height
	}
	return sx, sy
}


func (w *World) GetMapScale() (float64, float64) {
	var sx, sy float64 
	screenIW, screenIH := screen.GetScreen().Size()
	screenW := float64(screenIW)
	screenH := float64(screenIH)
	
	if screenW > w.Metadata.RawSize.Width {
		sx = w.Metadata.RawSize.Width / screenW
	} else {
		sx = screenW / w.Metadata.RawSize.Width
	}

	if screenH > w.Metadata.RawSize.Height {
	sy = w.Metadata.RawSize.Height / screenH
	} else {
		sy = screenH / w.Metadata.RawSize.Height
	}
	return sx, sy
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

func NewWorld() *World {
	world := new(World)
	id, err := uuid.NewUUID()
	if err != nil {
		logrus.Fatal("failed to create uuid for world:", err)
	}
	world.ID = id
	return world
}