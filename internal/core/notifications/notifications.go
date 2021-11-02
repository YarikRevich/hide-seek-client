package notifications

import (
	"fmt"
	"sync"
	"time"

	"github.com/YarikRevich/caching/pkg/zeroshifter"
)

type NotificatorEntity struct {
	Timestamp int64
	Message   string
}

type Notificator interface {
	WriteError(string)
	WriteWarning(string)
	WriteDebug(string)

	Read() []NotificatorEntity

	Filter(func(*NotificatorEntity) bool)
}

type notificator struct {
	sync.Mutex
	queue zeroshifter.IZeroShifter
}

func (p *notificator) write(m string) {
	p.Lock()
	timestamp := time.Now().Add(time.Second * 5).Unix()
	if r, ok := p.queue.IsExist(func(i interface{}) bool {
		return i.(*NotificatorEntity).Message == m
	}); ok {
		r.(*NotificatorEntity).Timestamp = timestamp
	} else {
		p.queue.Add(&NotificatorEntity{Timestamp: timestamp, Message: m})
	}
	p.Unlock()
}

func (p *notificator) WriteError(m string) {
	p.write(fmt.Sprintf("Error: %s", m))
}
func (p *notificator) WriteWarning(m string) {
	p.write(fmt.Sprintf("Warning: %s", m))
}
func (p *notificator) WriteDebug(m string) {
	p.write(fmt.Sprintf("Debug: %s", m))
}

func (p *notificator) Read() []NotificatorEntity {
	m := p.queue.Get()
	r := make([]NotificatorEntity, len(m))
	for _, v := range m {
		r = append(r, *v.(*NotificatorEntity))
	}
	return r
}

func (p *notificator) Filter(f func(*NotificatorEntity) bool) {
	p.queue.Filter(func(i interface{}) bool {
		return f(i.(*NotificatorEntity))
	})
}

func NewNotificator() Notificator {
	return &notificator{queue: zeroshifter.New(3)}
}

var PopUp = NewNotificator()
