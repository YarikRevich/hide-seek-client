package events

var instance *Events

type Events struct {
	mouse    *Mouse
	gamepad  *Gamepad
	keyboard *KeyBoard
	input    *Input
}

func (e *Events) Mouse() *Mouse {
	return e.mouse
}

func (e *Events) Gamepad() *Gamepad {
	return e.gamepad
}

func (e *Events) Keyboard() *KeyBoard {
	return e.keyboard
}

func (e *Events) Input() *Input {
	return e.input
}

func UseEvents() *Events {
	if instance == nil {
		instance = &Events{
			mouse:    NewMouse(),
			gamepad:  NewGamepad(),
			keyboard: NewKeyBoard(),
			input:    NewInput(),
		}
	}
	return instance
}
