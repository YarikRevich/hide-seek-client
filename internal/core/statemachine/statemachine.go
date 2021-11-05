package statemachine


var instance IStateMachine

type StateMachine struct{
	audio IState
	ui IState
	input IState
	networking IState
	dial IState
}

type IStateMachine interface {
	Audio() IState
	UI() IState
	Input() IState
	Networking() IState
	Dial() IState
}

func (s *StateMachine) Audio() IState {
	return s.audio
}

func (s *StateMachine) UI() IState {
	return s.ui
}

func (s *StateMachine) Input() IState {
	return s.input
}

func (s *StateMachine) Networking() IState {
	return s.networking
}

func (s *StateMachine) Dial() IState{
	return s.dial
}

func UseStateMachine() IStateMachine {
	if instance == nil {
		instance = &StateMachine{
			audio: NewState(AUDIO_DONE),
			ui: NewState(UI_START_MENU),
			input: NewState(INPUT_EMPTY),
			networking: NewState(NETWORKING_OFFLINE),
			dial: NewState(DIAL_WAN),
		}
	}
	return instance
}
