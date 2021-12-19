package objects

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/networking/api/server_external"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Weapon struct {
	Base

	Name, Radius string
}

func (w *Weapon) ToAPIMessage() *server_external.Weapon {
	return &server_external.Weapon{
		Base:   w.Base.ToAPIMessage(),
		Name:   w.Name,
		Radius: w.Radius,
	}
}

func NewWeapon() *Weapon {
	w := new(Weapon)
	id, err := uuid.NewUUID()
	if err != nil {
		logrus.Fatal("failed to create uuid for world:", err)
	}
	w.ID = id
	return w
}
