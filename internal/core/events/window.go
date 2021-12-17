package events

type Window struct{}

func (w *Window) OnStartUp() {

}

func (w *Window) OnClose() {

}

func NewWindow() *Window {
	return new(Window)
}
