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

func (p *provider) UsedManager() *UsedManager{
	
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
