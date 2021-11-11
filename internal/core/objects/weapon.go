package objects

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/networking/api"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Weapon struct {
	Object

	Name, Radius string
}

func NewWeaponByObject(o Object)*Weapon{
	w := new(Weapon)
	id, err := uuid.NewUUID()
	if err != nil {
		logrus.Fatal("failed to create uuid for world:", err)
	}
	w.ID = id

	w.ParentID = o.ID
	return w
}

func (w *Weapon) ToAPIMessage()*api.Weapon{
	return &api.Weapon{
		Object: w.Object.ToAPIMessage(),
		Name: w.Name,
		Radius: w.Radius,
	}
}

func NewWeapon()*Weapon{
	w := new(Weapon)
	id, err := uuid.NewUUID()
	if err != nil {
		logrus.Fatal("failed to create uuid for world:", err)
	}
	w.ID = id
	return w
}
