package statemachine

type SingleState struct {
	state int
}

type ISingleState interface {
	SetState(int)
	GetState() int

	//Checks if current state is equal to passed
	Check(int) bool
}

func (s *SingleState) SetState(st int) {
	s.state = st

}

func (s *SingleState) GetState() int {
	return s.state
}

func (s *SingleState) Check(v int) bool {
	return s.state == v
}

func NewSingleState(defaultState int) ISingleState {
	return &SingleState{state: defaultState}
}
