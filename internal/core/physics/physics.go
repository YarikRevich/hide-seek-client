package physics

var instance *Physics

type Physics struct {
	jump *Jump
}

func (p *Physics) Jump() *Jump{
	return p.jump
}

func UsePhysics()  *Physics{
	if instance == nil{
		instance = &Physics{
			jump: NewJump(),
		}
	}
	return instance
} 