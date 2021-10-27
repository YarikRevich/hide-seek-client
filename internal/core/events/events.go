package events

var instance EventsProvider

type provider struct {
	mouse    *Mouse
	gamepad  *Gamepad
	keyboard *KeyBoard
	input    *Input
}

type EventsProvider interface {
	Mouse() *Mouse
	Gamepad() *Gamepad
	Keyboard() *KeyBoard
	Input() *Input
}

func (p *provider) Mouse() *Mouse {
	return p.mouse
}

func (p *provider) Gamepad() *Gamepad {
	return p.gamepad
}

func (p *provider) Keyboard() *KeyBoard {
	return p.keyboard
}

func (p *provider) Input() *Input {
	return p.input
}

func UseEvents() EventsProvider {
	if instance == nil {
		instance = &provider{
			mouse:    NewMouse(),
			gamepad:  NewGamepad(),
			keyboard: NewKeyBoard(),
			input:    NewInput(),
		}
	}
	return instance
}
