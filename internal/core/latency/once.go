package latency

import "sync"

type Once struct {
	collection map[int]*sync.Once
}

func (o *Once) ExecOnce(s int, c func()) {
	if _, ok := o.collection[s]; !ok {
		o.collection[s] = new(sync.Once)
	}
	o.collection[s].Do(c)
}

func (o *Once) Reset() {
	for k := range o.collection {
		o.collection[k] = new(sync.Once)
	}
}

func NewOnce() *Once {
	return &Once{
		collection: make(map[int]*sync.Once),
	}
}
