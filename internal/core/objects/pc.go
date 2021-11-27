package objects

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/networking/api"
	"github.com/YarikRevich/HideSeek-Client/internal/core/storage"
	"github.com/YarikRevich/caching/pkg/zeroshifter"
	"github.com/google/uuid"
)

const DEFAULT_HEALTH = 10

type Buffs struct{ SpeedX, SpeedY float64 }

// type Team int

// const (
// 	Team1 Team = iota
// 	Team2
// )

type Status int

const (
	ALIVE Status = iota
	DEAD
	DIEING
)

type PC struct {
	Base

	Status Status
	Killer uuid.UUID

	Username string
	Health   uint64

	LobbyNumber int
	// Team Tea
}

// func (p *PC) GetScaledOffsetX() float64 {
// 	return (p.RawOffset.X * p.Parent.Modified.RuntimeDefined.ZoomedScale.X) - p.Modified.Offset.X
// }

// func (p *PC) GetScaledOffsetY() float64 {
// 	// fmt.Println(o.Parent.Modified.RuntimeDefined.ZoomedScale.Y)
// 	return ((p.RawOffset.Y - ) * p.Parent.Modified.RuntimeDefined.ZoomedScale.Y) - p.Modified.Offset.Y
// }

//Initializes pc username by requesting storage
func (p *PC) LoadUsername() {
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

func (p *PC) ToAPIMessage() *api.PC {
	return &api.PC{
		Base:     p.Base.ToAPIMessage(),
		Username: p.Username,
		Health:   p.Health,
	}
}

func (p *PC) FromAPIMessage(m *api.PC) {
	p.Base.FromAPIMessage(m.Base)
	p.Username = m.Username
	p.Health = m.Health
}

func (p *PC) String() string {
	return p.Username
}

func NewPC() *PC {
	pc := new(PC)
	pc.ID = uuid.New()
	pc.PositionHistorySequence = zeroshifter.New(2)
	pc.Health = 10
	return pc
}
