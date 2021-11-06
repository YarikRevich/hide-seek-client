package runtime

var instance *Runtime

type Runtime struct {
	prepared bool
}

func (r *Runtime) IsPrepared()bool{
	return r.prepared
}

func (r *Runtime) SetPrepared(){
	r.prepared = true
}

func UseRuntime() *Runtime{
	if instance == nil{
		instance = new(Runtime)
	}
	return instance
}