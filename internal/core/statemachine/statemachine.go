package statemachine


var instance IStateMachine

type StateMachine struct{
	audio IState
	ui IState
	input IState
	networking IState

}

type IStateMachine interface {
	Audio() IState
	UI() IState
	Input() IState
	Networking() IState
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

func UseStateMachine() IStateMachine {
	if instance == nil {
		instance = &StateMachine{
			audio: NewState(AUDIO_DONE),
			ui: NewState(UI_START_MENU),
			input: NewState(INPUT_EMPTY),
			networking: NewState(NETWORKING_OFFLINE),
		}
	}
	return instance
}
