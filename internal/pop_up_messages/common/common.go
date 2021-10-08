package common

import (
	"fmt"
	"sync"
	"time"

	"github.com/YarikRevich/caching/pkg/zeroshifter"
)

var instance IPopUpMessageBuff

type PopUpEntity struct {
	Timestamp int64
	Message   string
}

type IPopUpMessageBuff interface {
	write(string)

	WriteError(string)
	WriteWarning(string)
	WriteDebug(string)

	Read() []PopUpEntity

	Filter(func(*PopUpEntity) bool)
}

type popUpMessageBuff struct {
	sync.Mutex
	queue zeroshifter.IZeroShifter
}

func (p *popUpMessageBuff) write(m string) {
	p.Lock()
	timestamp := time.Now().Add(time.Second * 5).Unix()
	if r, ok := p.queue.IsExist(func(i interface{}) bool {
		return i.(*PopUpEntity).Message == m
	}); ok {
		r.(*PopUpEntity).Timestamp = timestamp
	} else {
		p.queue.Add(&PopUpEntity{Timestamp: timestamp, Message: m})
	}
	p.Unlock()
}

func (p *popUpMessageBuff) WriteError(m string) {
	p.write(fmt.Sprintf("Error: %s", m))
}
func (p *popUpMessageBuff) WriteWarning(m string) {
	p.write(fmt.Sprintf("Warning: %s", m))
}
func (p *popUpMessageBuff) WriteDebug(m string) {
	p.write(fmt.Sprintf("Debug: %s", m))
}

func (p *popUpMessageBuff) Read() []PopUpEntity {
	m := p.queue.Get()
	r := make([]PopUpEntity, len(m))
	for _, v := range m{
		r = append(r, *v.(*PopUpEntity))
	}
	return r
}

func (p *popUpMessageBuff) Filter(f func(e *PopUpEntity) bool) {
	p.queue.Filter(func(i interface{}) bool {
		return f(i.(*PopUpEntity))
	})
}

func NewPopUpMessageBuff() IPopUpMessageBuff {
	if instance == nil {
		instance = &popUpMessageBuff{
			queue: zeroshifter.New(3),
		}
	}
	return instance
}
