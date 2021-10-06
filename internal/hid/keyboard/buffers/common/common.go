package common

import "github.com/YarikRevich/HideSeek-Client/internal/cursor/collection"

type IBuffer interface {
	Clean()
	CleanBlinking()
	CleanBlinkingUnfocus()
	Push(rune)
	Pop()
	ReadClear()string
	Last() rune
	Read()string
}

type buff struct {
	value []rune
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

func (t *buff) CleanBlinking(){
	if l := t.Last(); l == collection.BlinkingOn || l == collection.BlinkingOff{
		t.Pop()
	}
}

func (t *buff) CleanBlinkingUnfocus(){
	t.CleanBlinking()
	t.Push(' ')
}

func (t *buff) 	ReadClear()string{
	if len(t.value) != 0{
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

func NewBuffer() IBuffer {
	return new(buff)
}
