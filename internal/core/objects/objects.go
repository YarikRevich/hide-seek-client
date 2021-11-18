package objects

var instance ObjectsProvider

type provider struct{
	pc *PC
	world *World
	camera *Camera
}

type ObjectsProvider interface {
	PC() *PC
	World() *World
	Camera() *Camera
}

func (p *provider) PC() *PC {
	return p.pc
}

func (p *provider) World() *World {
	return p.world
}

func (p *provider) Camera() *Camera{
	return p.camera
}

func UseObjects() ObjectsProvider {
	if instance == nil{
		instance = &provider{
			pc: NewPC(),
			world: NewWorld(),
			camera: NewCamera(),
		}
	}
	return instance 
}
