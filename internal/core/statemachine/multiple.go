package statemachine

import "github.com/google/uuid"

type MultipleState struct {
	states       map[uuid.UUID]int
	defaultState int
}

type IMultipleState interface {
	SetState(uuid.UUID, int)
	GetState(uuid.UUID) int
	Reset()
}

func (s *MultipleState) SetState(i uuid.UUID, st int) {
	s.states[i] = st
}

func (s *MultipleState) GetState(i uuid.UUID) int {
	v, ok := s.states[i]
	if !ok {
		s.states[i] = s.defaultState
		return s.defaultState
	}
	return v
}

func (s *MultipleState) Reset() {
	s.states = make(map[uuid.UUID]int)
}

func NewMultipleState(defaultState int) IMultipleState {
	return &MultipleState{
		defaultState: defaultState,
		states:       make(map[uuid.UUID]int)}
}
