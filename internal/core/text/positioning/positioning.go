package positioning

var instance *Positioning

type Positioning struct {
	input  *Input
	button *Button
}

func (p *Positioning) Input() *Input {
	return p.input
}

func (p *Positioning) Button() *Button {
	return p.button
}

//TODO: refactor interface for text positioning
func UsePositioning() *Positioning {
	if instance == nil {
		instance = &Positioning{
			input:  NewInput(),
			button: NewButton(),
		}
	}
	return instance
}
