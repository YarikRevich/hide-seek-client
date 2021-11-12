package statemachine

var instance IStateMachine

type StateMachine struct {
	audio                  IState
	ui                     IState
	input                  IState
	networking             IState
	dial                   IState
	settings_menu_checkbox IState
	game                   IState
}

type IStateMachine interface {
	Audio() IState
	Game() IState
	UI() IState
	Input() IState
	Networking() IState
	Dial() IState
	SettingsMenuCheckbox() IState
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

func (s *StateMachine) Dial() IState {
	return s.dial
}

func (s *StateMachine) SettingsMenuCheckbox() IState {
	return s.settings_menu_checkbox
}

func (s *StateMachine) Game() IState {
	return s.game
}

func UseStateMachine() IStateMachine {
	if instance == nil {
		instance = &StateMachine{
			audio:                  NewState(AUDIO_DONE),
			ui:                     NewState(UI_START_MENU),
			input:                  NewState(INPUT_EMPTY),
			networking:             NewState(NETWORKING_ONLINE),
			dial:                   NewState(DIAL_WAN),
			settings_menu_checkbox: NewState(UI_SETTINGS_MENU_CHECKBOX_OFF),
			game:                   NewState(GAME_START),
		}
	}
	return instance
}
