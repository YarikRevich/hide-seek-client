package events

const (
	ELEMENT ObjectType = iota
	AMMO
	WEAPON
	PLAYER
	MAP
)

type Collisions struct {
	collection map[struct{ X, Y float64 }]int
}

func (c *Collisions) IsCollide(X, Y float64) bool {
	return false
}

func (c *Collisions) RegisterCollisions() {

}

func NewCollisions() *Collisions {
	return new(Collisions)
}
