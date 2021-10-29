package objects

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/keycodes"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Ammo struct {
	Object

	Direction keycodes.Direction
}

func NewAmmoByObject(o Object)*Ammo{
	a := new(Ammo)
	id, err := uuid.NewUUID()
	if err != nil {
		logrus.Fatal("failed to create uuid for world:", err)
	}
	a.ID = id
	a.ParentID = o.ID
	return a
}


func NewAmmo()*Ammo{
	a := new(Ammo)
	id, err := uuid.NewUUID()
	if err != nil {
		logrus.Fatal("failed to create uuid for world:", err)
	}
	a.ID = id
	return a
}
