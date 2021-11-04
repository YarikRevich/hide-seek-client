package middlewares

var instance *Middlewares

type Middlewares struct {
	render *Render
	prepare *Prepare
	ui *UI
	audio *Audio
}

func (m *Middlewares)Render()*Render {
	return m.render
}

func (m *Middlewares) Prepare() *Prepare{
	return m.prepare
}

func (m *Middlewares) Audio() *Audio{
	return m.audio
}

func (m *Middlewares) UI() *UI{
	return m.ui
}


func UseMiddlewares()*Middlewares{
	if instance == nil{
		instance = &Middlewares{
			render: NewRender(),
			prepare: NewPrepare(),
			ui: NewUI(),
			audio: NewAudio(),
		}
	}
	return instance
}