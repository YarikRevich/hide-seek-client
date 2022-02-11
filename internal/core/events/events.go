package events

var KeyboardPress = NewKeyboardPressEventManager()
var MousePress = NewMousePressEventManager()
var MouseScroll = NewMouseScrollEventManager()
var Gamepad = NewGamepadEventManager()
var GamepadPress = NewGamepadPressEventManager()
var GamepadScroll = NewGamepadScrollEventManager()

// func NewEventsManager() EventsManager {
// 	return &provider{
// 		mouse:    NewMouse(),
// 		gamepad:  NewGamepad(),
// 		keyboard: NewKeyBoard(),
// 		input:    NewInput(),
// 		window:   NewWindow(),
// 	}
// }
