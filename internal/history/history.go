package history

type Direction int

const (
	UP Direction = iota
	DOWN
	RIGHT
	LEFT
)

var (
	lastDirection Direction
)

func SetDirection(d Direction){
	lastDirection = d
}

func GetDirection()Direction{
	return lastDirection
}