package objects

import (
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

func NewWeapon()*Weapon{
	w := new(Weapon)
	id, err := uuid.NewUUID()
	if err != nil {
		logrus.Fatal("failed to create uuid for world:", err)
	}
	w.ID = id
	return w
}
