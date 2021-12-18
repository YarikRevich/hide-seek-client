package events

type Collision struct {
	collection map[struct{ X, Y float64 }]func()
}

// func (c *Collision)

func NewCollision() *Collision {
	return new(Collision)
}
