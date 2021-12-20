package statemachine

type SingleState struct {
	state int
}

type ISingleState interface {
	SetState(int)
	GetState() int
}

func (s *SingleState) SetState(st int) {
	s.state = st

}

func (s *SingleState) GetState() int {
	return s.state
}

func NewSingleState(defaultState int) ISingleState {
	return &SingleState{state: defaultState}
}
