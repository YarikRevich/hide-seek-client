package objects

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/networking/api/server_external"
	"github.com/google/uuid"
)

type Weapon struct {
	Base
	Opts WeaponOpts
}

type WeaponOpts struct {
	Name, Radius string
}

func (w *Weapon) ToAPIMessage() *server_external.Weapon {
	return &server_external.Weapon{
		Base:   w.Base.ToAPIMessage(),
		Name:   w.Opts.Name,
		Radius: w.Opts.Radius,
	}
}

func NewWeapon() *Weapon {
	return &Weapon{
		Base: Base{
			ID:   uuid.New(),
			Type: WEAPON,
		}}
}
