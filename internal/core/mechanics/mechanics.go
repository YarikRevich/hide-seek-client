package mechanics

var instance Mechanics

type mechanics struct {}

type Mechanics interface {

}

func UseMechanics() Mechanics{
	if instance == nil{
		instance = new(mechanics)
	}
	return instance
}