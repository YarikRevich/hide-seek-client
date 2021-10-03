package audio

import "sync"

const (
	DONE statusEntry = iota
	UNDONE
)

var (
	stateMachine *Status
)

type statusEntry int

type Status struct {
	sync.Mutex
	status statusEntry
}

func (s *Status) SetState(st statusEntry) {
	s.Lock()
	defer s.Unlock()
	s.status = st
}

func (s *Status) GetState() statusEntry {
	s.Lock()
	defer s.Unlock()
	return s.status
}

func UseStatus() *Status {
	if stateMachine == nil {
		stateMachine = &Status{status: DONE}
	}
	return stateMachine
}
