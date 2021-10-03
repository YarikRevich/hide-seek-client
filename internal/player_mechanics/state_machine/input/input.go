package input

const (
	EMPTY statusEntry = iota
	SETTINGS_MENU_USERNAME
	GAME
)

var (
	stateMachine *Status
)

type statusEntry int

type Status struct {
	status statusEntry
}

func (s *Status) SetState(st statusEntry) {
	s.status = st
}

func (s *Status) GetState() statusEntry {
	return s.status
}

func UseStatus() *Status {
	if stateMachine == nil {
		stateMachine = &Status{status: EMPTY}
	}
	return stateMachine
}
