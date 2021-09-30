package buffers

type Control interface {
	Clean()
}

type TextBuffer interface {
	Control

	Write(string)
	Read()string
}