package system

var (
	system *System
)

type System struct {
	root string
}



func GetInstace()*System{
	if system == nil{
		return new(System)
	}
	return system
}