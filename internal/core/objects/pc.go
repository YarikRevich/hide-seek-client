package objects

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/keycodes"
	"github.com/YarikRevich/HideSeek-Client/internal/core/networking/api"
	"github.com/YarikRevich/HideSeek-Client/internal/core/storage"

	// strorageprovider "github.com/YarikRevich/HideSeek-Client/internal/storage/provider"
	"github.com/YarikRevich/caching/pkg/zeroshifter"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const DEFAULT_HEALTH = 10

type Buffs struct{ SpeedX, SpeedY float64 }

// type Team int

// const (
// 	Team1 Team = iota
// 	Team2
// )

type PC struct {
	Object

	Username string
	Health   uint64

	
	// Team Tea
}

// func (p *PC) SetSkin(path string){
// 	p.Object.SetSkin(path)
// 	// p.SetSpeed(p.Metadata.Buffs.Speed)
// 	fmt.Println(p.Metadata, "METEDATA")
// }

//Initializes pc username by requesting storage
func (p *PC) SetUsername() {
	p.Username = storage.UseStorage().User().GetUsername()
}

// func (p *PC) SetSpeed(speedX float64) {
// 	p.Buffs.SpeedX = speedX
// 	p.Buffs.SpeedY = spe
// 	// wx, wy := screen.GetMaxWidth(), screen.GetMaxHeight()
// 	// p.Buffs.SpeedY = speedX * float64(wy) / float64(wx) * 2
// }

// //Returns movement rotation related to the last
// //movement direction
func (p *PC) GetMovementRotation() float64 {
	if p.Direction == keycodes.LEFT && p.RawPos.X != 0 {
		return -1
	}
	return 1
}

func (p *PC) ToAPIMessage() *api.PC{
	return &api.PC{
		Object: p.Object.ToAPIMessage(),
		Username: p.Username,
		Health: p.Health,
	}
}

func (p *PC) FromAPIMessage(m *api.PC){
	p.Object.FromAPIMessage(m.Object)
	p.Username = m.Username
	p.Health = m.Health
}

func (p *PC) String()string{
	return p.Username
}

func NewPC() *PC {
	pc := new(PC)
	id, err := uuid.NewUUID()
	if err != nil {
		logrus.Fatal("failed to create uuid for world:", err)
	}
	pc.ID = id
	pc.PositionHistory = zeroshifter.New(2)
	pc.Health = 10
	// pc.Spawn = image.Point{X: 1500, Y: 0}
	// instance.Animation.FrameDelay = 5
	// instance.Animation.FrameDelayCounter = 1

	// instance.Image = imagecollection.GetImage("assets/images/heroes/pumpkin")
	// instance.Metadata = metadatacollection.GetMetadata("assets/images/heroes/pumpkin")
	return pc
}
