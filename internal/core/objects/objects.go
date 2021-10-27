package objects

var instance ObjectsProvider

type provider struct{
	pc *PC
	world *World
}

type ObjectsProvider interface {
	PC() *PC
	World() *World
}

func (p *provider) PC() *PC {
	return p.pc
}

func (p *provider) World() *World {
	return p.world
}

func UseObjects() ObjectsProvider {
	if instance == nil{
		instance = &provider{
			pc: NewPC(),
			world: NewWorld(),
		}
	}
	return instance 
}
