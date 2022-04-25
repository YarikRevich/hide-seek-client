package objects

import (
	"github.com/YarikRevich/caching/pkg/zeroshifter"
	"github.com/YarikRevich/hide-seek-client/internal/core/networking/api/server_external"
	"github.com/YarikRevich/hide-seek-client/internal/core/storage"
	"github.com/google/uuid"
)

const DEFAULT_HEALTH = 10

type Buffs struct{ SpeedX, SpeedY float64 }

type PC struct {
	Base

	Opts PCOpts

	//TODO: set a temporary killer name
	// Killer uuid.UUID

	// LobbyNumber int

	//TODO: set a temporary kicked state
	// IsKicked bool

	// LastActivity int64
}

type PCOpts struct {
	Username string
	Health   uint64
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
	p.Opts.Username = storage.UseStorage().User().GetUsername()
}

// func (p *PC) DebugInit() {
// 	p.Base.SetTilemap("heroes/pumpkin")
// }

// //Returns movement rotation related to the last
// //movement direction
func (p *PC) GetMovementRotation() float64 {
	if p.IsDirectionLEFT() && p.RawPos.X != 0 {
		return -1
	}
	return 1
}

// func (p *PC) UpdateLastActivity() {
// 	p.LastActivity = time.Now().Unix()
// }

// func (p *PC) SetKicked(s bool) {
// 	p.Opts.IsKicked = s
// }

func (p *PC) ToAPIMessage() *server_external.PC {
	return &server_external.PC{
		Base:     p.Base.ToAPIMessage(),
		Username: p.Opts.Username,
		Health:   p.Opts.Health,
	}
}

func (p *PC) FromAPIMessage(m *server_external.PC) {
	p.Base.FromAPIMessage(m.Base)
	p.Opts.Username = m.Username
	p.Opts.Health = m.Health
}

func (p *PC) String() string {
	return p.Opts.Username
}

func NewPC(opts PCOpts) *PC {
	return &PC{
		Base: Base{
			ID:                      uuid.New(),
			Type:                    PLAYER,
			PositionHistorySequence: zeroshifter.New(2),
		},
		Opts: PCOpts{Health: 10}}
}
