package middlewares

var instance *Middlewares

type Middlewares struct {
	render *Render
	prepare *Prepare
}

func (m *Middlewares)Render()*Render {
	return m.render
}

func (m *Middlewares) Prepare() *Prepare{
	return m.prepare
}

func UseMiddlewares()*Middlewares{
	if instance == nil{
		instance = &Middlewares{
			render: NewRender(),
			prepare: NewPrepare(),
		}
	}
	return instance
}