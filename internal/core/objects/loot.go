package objects

type LootSet struct {
	Base

	IsHidden bool
	Weapon   []*Weapon
	Ammo     []*Ammo
}

func NewLootSet(w []*Weapon, a []*Ammo) *LootSet {
	l := new(LootSet)
	l.Weapon = w
	l.Ammo = a
	l.Type = "lootset"
	return l
}
