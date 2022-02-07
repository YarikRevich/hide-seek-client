package notifications

import (
	"math"
	"time"

	"github.com/YarikRevich/caching/pkg/zeroshifter"
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"
)

const defaultTTL = 7

const (
	Fatal = iota
	Error
	Debug
	Info
)

type Notification struct {
	TTL     int64
	Level   int
	Message string
}

type NotificationManager struct {
	queue zeroshifter.IZeroShifter
}

func (nm *NotificationManager) Write(mess string, ttl, level int) {
	if ttl == -1 {
		ttl = defaultTTL
	}
	extendedTTL := time.Now().Add(time.Second * time.Duration(ttl)).Unix()
	r, ok := nm.queue.IsExist(func(i interface{}) bool {
		return i.(*Notification).Message == mess
	})
	if ok {
		r.(*Notification).TTL = extendedTTL
	} else {
		statemachine.Notification.SetState(statemachine.NOTIFICATION_NEW)
		nm.queue.Add(&Notification{
			TTL:     extendedTTL,
			Message: mess,
			Level:   level})
	}
}

func (nm *NotificationManager) Read() []Notification {
	m := nm.queue.Get()
	var r []Notification
	for _, v := range m {
		r = append(r, (*v.(*Notification)))
	}
	return r
}

func (nm *NotificationManager) DeleteOld(f func(*Notification) bool) {
	nm.queue.Filter(func(i interface{}) bool {
		return func(e *Notification) bool {
			return math.Signbit(float64(time.Now().Unix() - e.TTL))
		}(i.(*Notification))
	})
}

func NewNotificationManager() *NotificationManager {
	return &NotificationManager{queue: zeroshifter.New(3)}
}
