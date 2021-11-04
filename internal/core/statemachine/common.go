package statemachine

type State struct {
	state int
}

type IState interface {
	SetState(int)
	GetState() int
}

func (s *State) SetState(st int) {
	s.state = st

}

func (s *State) GetState() int {
	return s.state
}

func NewState(defaultState int) IState {
	return &State{state: defaultState}
}
