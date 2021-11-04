package animation

var instance *Animation

type Animation struct {
	pc *PCAnimation
}

func (a *Animation) PC()*PCAnimation{
	return a.pc
}

func UseAnimation() *Animation{
	if instance == nil{
		instance = &Animation{
			pc: NewPCAnimation(),
		}
	}
	return instance
}