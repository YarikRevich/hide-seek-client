package objects

type Camera struct {
	Object
}


func NewCamera() *Camera{
	return new(Camera)
}