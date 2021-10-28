package objects

import "github.com/YarikRevich/HideSeek-Client/internal/core/keycodes"

type Ammo struct {
	Object

	Direction keycodes.Direction
}

func NewAmmo(){}