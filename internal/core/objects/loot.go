package objects

type LootSet struct {
	Base
	Opts LootSetOpts
}

type LootSetOpts struct {
	IsHidden bool
	Weapon   []*Weapon
	Ammo     []*Ammo
}

func (ls *LootSet) ToAPIMessage()   {}
func (ls *LootSet) FromAPIMessage() {}

func NewLootSet(opts LootSetOpts) *LootSet {
	return &LootSet{
		Base: Base{
			Type: LOOT,
		}}
}
