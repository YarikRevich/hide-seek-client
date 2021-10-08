package profiling


const (
	UI = "ui"
	MOUSE_HANDLER = "mouse_handler"
	KEYBOARD_HANDLER = "keyboard_handler"
	AUDIO_HANDLER = "audio_handler"
)

type Profiler map[string]string

func (p *Profiler) Add(handler string, )
