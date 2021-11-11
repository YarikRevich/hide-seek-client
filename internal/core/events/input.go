package events

import "time"

const (
	blinkingOn  = '|'
	blinkingOff = ' '
)

type IBuffer interface {
	Clean()

	UpdateCursorBlink()
	CleanBlinking()
	CleanBlinkingUnfocus()

	Push(rune)
	Pop()
	ReadClean() string
	Last() rune
	Read() string
}

type buff struct {
	value         []rune
	blinkTicker   *time.Ticker
	blinkPosition rune
}

func (t *buff) Read() string {
	return string(t.value)
}

func (t *buff) Last() rune {
	if len(t.value) != 0 {
		return t.value[len(t.value)-1]
	}
	return '0'
}

func (t *buff) Clean() {
	t.value = t.value[:0]
}

func (t *buff) CleanBlinking() {
	if l := t.Last(); l == blinkingOn || l == blinkingOff {
		t.Pop()
	}
}

func (t *buff) CleanBlinkingUnfocus() {
	t.CleanBlinking()
	t.Push(' ')
}

func (t *buff) ReadClean() string {
	if len(t.value) != 0 {
		v := t.Read()
		return v[:len(v)-1]
	}
	return ""
}

func (t *buff) Push(v rune) {
	t.value = append(t.value, v)
}

func (t *buff) Pop() {
	if len(t.value) != 0 {
		t.value = t.value[:len(t.value)-1]
	}
}

//Updates cursor blink of the input
func (t *buff) UpdateCursorBlink() {
	select {
	case <-t.blinkTicker.C:
		if t.blinkPosition == blinkingOn {
			t.blinkPosition = blinkingOff
		} else {
			t.blinkPosition = blinkingOn
		}
	default:
	}

	l := t.Last()
	if l == blinkingOn || l == blinkingOff {
		t.Pop()
		t.Push(t.blinkPosition)
	} else {
		t.Push(t.blinkPosition)
	}
}

func NewBuffer() IBuffer {
	return &buff{
		blinkTicker: time.NewTicker(time.Second),
	}
}

type Input struct {
	InputBuffers
}

type InputBuffers struct {
	SettingsMenuNameBuffer, JoinGameBuffer IBuffer
}

func (i *InputBuffers) CleanBuffersBlinking(){
	i.SettingsMenuNameBuffer.CleanBlinkingUnfocus()
}

func NewInput() *Input{
	return &Input{
		InputBuffers: InputBuffers{
			SettingsMenuNameBuffer: NewBuffer(),
			JoinGameBuffer: NewBuffer(),
		},
	}
}