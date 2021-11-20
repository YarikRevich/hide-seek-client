package objects

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/networking/api"
	"github.com/YarikRevich/HideSeek-Client/internal/core/storage"
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


//Initializes pc username by requesting storage
func (p *PC) SetUsername() {
	p.Username = storage.UseStorage().User().GetUsername()
}

// //Returns movement rotation related to the last
// //movement direction
func (p *PC) GetMovementRotation() float64 {
	if p.IsDirectionLEFT() && p.RawPos.X != 0 {
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
	return pc
}
