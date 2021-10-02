package pc

import (
	"crypto/sha256"

	"github.com/YarikRevich/HideSeek-Client/internal/direction"
	metadataloader "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/loader/metadata_loader"
)

const (
	EMPTY = ""
	DEFAULT_HEALTH = 10
)

var (
	pc *PC
)

type Animation struct {
	FrameCount uint32
	FrameDelay uint32
	FrameDelayCounter uint32
	// CurrentFrame uint32
	CurrentFrameMatrix []float64
}

type GameCredentials struct {
	LobbyID string
}

type Skin struct {
	ImageHash [sha256.Size]byte
	Animation Animation
}

type Weapon struct {
	Name string
	Radius int
	Animation Animation
}

type Equipment struct {
	Skin Skin
	Weapon string
}

type Buffs struct {
	Speed float64 
}

type Physics struct {
	Jump []direction.Direction
}

type PC struct{
	Username string

	Health uint64

	X float64
	Y float64

	Buffs Buffs

	Physics  Physics
	Equipment Equipment

	Metadata *metadataloader.Metadata

	GameCredentials GameCredentials
}

type PCs []PC

// userConfig := Users.User{
// 	// Conn: conn,
// 	// Pos: &Users.Pos{
// 		X: int(randomSpawn.X),
// 		Y: int(randomSpawn.Y),
// 	// },
// 	// GameInfo: &Users.GameInfo{
// 	// 	Health: 10,
// 	// 	WeaponName:  Utils.GetRandomWeaponImage(winConf.Components.AvailableWeaponImages),
// 	// },
// 	// PersonalInfo: &Users.PersonalInfo{
// 	// 	Username:    username,
// 	// 	HeroPicture: Utils.GetRandomHeroImage(winConf.Components.AvailableHeroImages),
// 	// },
// 	Animation:  &Users.Animation{CurrentFrameMatrix: []float64{0, 0, 0, 0}},
// 	// Networking: new(Users.Networking),
// 	// Context:    new(Users.Context),
// }


func GetPC()*PC{
	if pc == nil{
		pc = new(PC)
		pc.Username = EMPTY
		pc.Health = DEFAULT_HEALTH
		pc.Buffs.Speed = 2.5
		pc.Equipment.Skin.Animation.FrameDelay = 5
		pc.Equipment.Skin.Animation.FrameDelayCounter = 1
		pc.Metadata = metadataloader.GetMetadata("assets/images/heroes/pumpkinhero")
	}
	return pc
}


