package profiling

import (
	"fmt"
	"time"
)

var instance IProfiler

const (
	RENDER           handler = "render"
	UI               handler = "ui"
	MOUSE_HANDLER    handler = "mouse_handler"
	KEYBOARD_HANDLER handler = "keyboard_handler"
	AUDIO_HANDLER    handler = "audio_handler"
)

type handler string

type handlers []*monitoring

type monitoring struct {
	handler     handler
	currentTime time.Time
	time        []float64
	avg         float64
}

type profiler struct {
	current handler

	delay    *time.Ticker
	handlers handlers
}

type IProfiler interface {
	StartMonitoring(handler)
	EndMonitoring()
	SumUpMonitoring()
	GetFormattedMonitoringLog() string
}

func (p *profiler) StartMonitoring(name handler) {
	p.current = name
	if GetMonitoringByHandler(p.current, p.handlers) == nil {
		p.handlers = append(p.handlers, &monitoring{handler: p.current})
	}
	GetMonitoringByHandler(p.current, p.handlers).currentTime = time.Now()
}

func (p *profiler) EndMonitoring() {
	h := GetMonitoringByHandler(p.current, p.handlers)
	h.time = append(h.time, float64(time.Since(h.currentTime).Seconds()))
}

func (p *profiler) SumUpMonitoring() {
	select {
	case <-p.delay.C:
		for _, h := range p.handlers {
			var sum float64
			for _, t := range h.time {
				sum += t
			}
			h.avg = sum / float64(len(h.time))
			h.time = h.time[:0]
		}
	default:
	}
}

func (p *profiler) GetFormattedMonitoringLog() string {
	var r string
	for _, v := range p.handlers {
		r += fmt.Sprintf("\n%s: %g", v.handler, v.avg)
	}
	return r
}

func UseProfiler() IProfiler {
	if instance == nil {
		instance = &profiler{
			handlers: handlers{},
			delay:    time.NewTicker(time.Millisecond * 400)}
	}
	return instance
}
