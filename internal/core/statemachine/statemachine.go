package statemachine

var Layers = NewSingleState(LAYERS_START_MENU)
var Audio = NewSingleState(AUDIO_DONE)
var Input = NewSingleState(INPUT_EMPTY)
var Networking = NewSingleState(NETWORKING_ONLINE)
var Dial = NewSingleState(DIAL_WAN)

// var UI = NewSingleState(UI_SETTINGS_MENU_CHECKBOX_OFF)
var Session = NewSingleState(GAME_START)
var Mouse = NewSingleState(MOUSE_NONE)
var Notification = NewSingleState(NOTIFICATION_NONE)
var Minimap = NewSingleState(MINIMAP_OFF)

var PCs = NewMultipleState(PC_ALIVE)

// type StateMachine struct {
// 	audio                  ISingleState
// 	ui                     ISingleState
// 	input                  ISingleState
// 	networking             ISingleState
// 	dial                   ISingleState
// 	settings_menu_checkbox ISingleState
// 	game                   ISingleState
// 	pcs                    IMultipleState
// 	mouse                  ISingleState
// 	notification           ISingleState
// 	minimap                ISingleState
// }

// type IStateMachine interface {
// 	Audio() ISingleState
// 	Game() ISingleState
// 	UI() ISingleState
// 	Input() ISingleState
// 	Networking() ISingleState
// 	Dial() ISingleState
// 	SettingsMenuCheckbox() ISingleState
// 	PCs() IMultipleState
// 	Mouse() ISingleState
// 	Notification() ISingleState
// 	Minimap() ISingleState
// }

// func (s *StateMachine) Audio() ISingleState {
// 	return s.audio
// }

// func (s *StateMachine) UI() ISingleState {
// 	return s.ui
// }

// func (s *StateMachine) Input() ISingleState {
// 	return s.input
// }

// func (s *StateMachine) Networking() ISingleState {
// 	return s.networking
// }

// func (s *StateMachine) Dial() ISingleState {
// 	return s.dial
// }

// func (s *StateMachine) SettingsMenuCheckbox() ISingleState {
// 	return s.settings_menu_checkbox
// }

// func (s *StateMachine) Game() ISingleState {
// 	return s.game
// }

// func (s *StateMachine) PCs() IMultipleState {
// 	return s.pcs
// }

// func (s *StateMachine) Mouse() ISingleState {
// 	return s.mouse
// }

// func (s *StateMachine) Notification() ISingleState {
// 	return s.notification
// }

// func (s *StateMachine) Minimap() ISingleState {
// 	return s.minimap
// }

// func UseStateMachine() IStateMachine {
// 	if instance == nil {
// 		instance = &StateMachine{
// 			audio:                  NewSingleState(AUDIO_DONE),
// 			ui:                     NewSingleState(UI_START_MENU),
// 			input:                  NewSingleState(INPUT_EMPTY),
// 			networking:             NewSingleState(NETWORKING_ONLINE),
// 			dial:                   NewSingleState(DIAL_WAN),
// 			settings_menu_checkbox: NewSingleState(UI_SETTINGS_MENU_CHECKBOX_OFF),
// 			game:                   NewSingleState(GAME_START),
// 			pcs:                    NewMultipleState(PC_ALIVE),
// 			mouse:                  NewSingleState(MOUSE_NONE),
// 			notification:           NewSingleState(NOTIFICATION_NONE),
// 			minimap:                NewSingleState(MINIMAP_OFF),
// 		}
// 	}
// 	return instance
// }
