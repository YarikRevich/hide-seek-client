package objects

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/keycodes"
	"github.com/YarikRevich/HideSeek-Client/internal/core/networking/api"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Ammo struct {
	Base

	Direction keycodes.Direction
}

func (a *Ammo) ToAPIMessage() *api.Ammo {
	return &api.Ammo{
		Base:    a.Base.ToAPIMessage(),
		Direction: int64(a.Direction),
	}
}

func (a *Ammo) FromAPIMessage(m *api.Ammo) {
	a.Base.FromAPIMessage(m.Base)
	a.Direction = keycodes.Direction(m.Direction)
}

func NewAmmo() *Ammo {
	a := new(Ammo)
	id, err := uuid.NewUUID()
	if err != nil {
		logrus.Fatal("failed to create uuid for world:", err)
	}
	a.ID = id
	return a
}
