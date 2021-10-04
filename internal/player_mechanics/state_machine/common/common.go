package common

type IState interface {
	SetState(int) func()
	GetState() int
}