package objects

import (
	"fmt"
	// "image"
	// "strings"

	// "github.com/YarikRevich/HideSeek-Client/internal/gameplay/objects"
	// "github.com/YarikRevich/HideSeek-Client/internal/gameplay/pc"
	"github.com/google/uuid"
	// "github.com/hajimehoshi/ebiten/v2"
	"github.com/sirupsen/logrus"
)

type regime int

const (
	//Players are separated on two teams
	//If player dies it will respawn
	//If the biggest part of the players is
	//on the oposite teritory, spawns will be swapped
	deathmatch regime = iota

	//Players are separated on two teams
	//The game will end if all members of the team
	//will dir
	teamToTeam
)

type World struct {
	Object

	Regime regime
	PCs []*PC
	// Elements []
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
func (w *World) GetMapScale(screenW, screenH int) (float64, float64) {
	var sx, sy float64
	if screenW > int(w.Metadata.RawSize.Width) {
		sx = w.Metadata.RawSize.Width / float64(screenW)
	} else {
		sx = float64(screenW) / w.Metadata.RawSize.Width
	}

	if screenH > int(w.Metadata.RawSize.Height) {
		sy = w.Metadata.RawSize.Height / float64(screenH)
	} else {
		sy = float64(screenH) / w.Metadata.RawSize.Height
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
