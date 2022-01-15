package notifications

import (
	"fmt"
	"sync"
	"time"

	"github.com/YarikRevich/caching/pkg/zeroshifter"
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"
)

const DEFAULT_POPUP_MESSAGE_TIME = 7

type NotificatorEntity struct {
	Timestamp int64
	Message   string
}

type Notificator interface {
	WriteError(string)
	WriteWarning(string)
	WriteDebug(string)
	WriteInfo(string)
	WriteInfoWithPopUpTime(string, int)

	Read() []NotificatorEntity

	Filter(func(*NotificatorEntity) bool)
}

type notificator struct {
	sync.Mutex
	queue zeroshifter.IZeroShifter
}

//Writes popup message
//popTime param means a time in seconds
//the popup message will be shown
func (p *notificator) write(m string, popTime int) {
	p.Lock()
	timestamp := time.Now().Add(time.Second * time.Duration(popTime)).Unix()
	if r, ok := p.queue.IsExist(func(i interface{}) bool {
		return i.(*NotificatorEntity).Message == m
	}); ok {
		r.(*NotificatorEntity).Timestamp = timestamp
	} else {
		statemachine.UseStateMachine().Notification().SetState(statemachine.NOTIFICATION_NEW)
		p.queue.Add(&NotificatorEntity{Timestamp: timestamp, Message: m})
	}
	p.Unlock()
}

func (p *notificator) WriteError(m string) {
	p.write(fmt.Sprintf("Error: %s", m), DEFAULT_POPUP_MESSAGE_TIME)
}

func (p *notificator) WriteWarning(m string) {
	p.write(fmt.Sprintf("Warning: %s", m), DEFAULT_POPUP_MESSAGE_TIME)
}

func (p *notificator) WriteDebug(m string) {
	p.write(fmt.Sprintf("Debug: %s", m), DEFAULT_POPUP_MESSAGE_TIME)
}

func (p *notificator) WriteInfo(m string) {
	p.write(fmt.Sprintf("Info: %s", m), DEFAULT_POPUP_MESSAGE_TIME)
}

func (p *notificator) WriteInfoWithPopUpTime(m string, popTime int) {
	p.write(fmt.Sprintf("Info: %s", m), popTime)
}

func (p *notificator) Read() []NotificatorEntity {
	m := p.queue.Get()
	var r []NotificatorEntity
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
