package buffers

type Control interface {
	Clean()
}

type TextBuffer interface {
	Control

	Push(rune)
	Pop()
	Last() rune
	Read()string
}