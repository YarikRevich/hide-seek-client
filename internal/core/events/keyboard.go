package events

type Direction int

const (
	UP Direction = iota
	DOWN
	RIGHT
	LEFT
)

type KeyBoard struct {

}

func (k *KeyBoard) 