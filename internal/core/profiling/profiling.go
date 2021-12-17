package profiling

import (
	"container/list"
	"errors"
	"fmt"
	"os"
	"runtime/pprof"
	"time"

	"github.com/YarikRevich/HideSeek-Client/internal/core/paths"
	"github.com/sirupsen/logrus"
)

var instance IProfiler

type handler string

const (
	RENDER        handler = "render"
	UI            handler = "ui"
	UI_START_MENU handler = "ui_start_menu"
	UI_GAME_MENU  handler = "ui_game_menu"
	UI_ANIMATION  handler = "ui_animation"

	MOUSE    handler = "mouse"
	KEYBOARD handler = "keyboard"
	AUDIO    handler = "audio"
)

type handlers []*monitoring

type monitoring struct {
	handler     handler
	currentTime time.Time
	time        []float64
	avg         float64
}

type profiler struct {
	delay           *time.Ticker
	handlers        handlers
	monitoringQueue *list.List
}

type IProfiler interface {
	Init()
	StartMonitoring(handler)
	EndMonitoring()
	String() string
}

func (p *profiler) Init() {
	mem, err := os.OpenFile(paths.GAME_PPROF_DIR+"/mem.prof", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		logrus.Fatal(err)
	}

	if err := pprof.WriteHeapProfile(mem); err != nil {
		logrus.Fatal(err)
	}

	p.handlers = []*monitoring{
		{handler: UI},
		{handler: UI_START_MENU},
		{handler: UI_GAME_MENU},
		{handler: UI_ANIMATION},
		{handler: MOUSE},
		{handler: KEYBOARD},
		{handler: AUDIO},
		{handler: RENDER},
	}
}

func (p *profiler) getMonitoring(name handler) (*monitoring, error) {
	for _, v := range p.handlers {
		if v != nil && v.handler == name {
			return v, nil
		}
	}
	return nil, errors.New("monitoring handler was not found")
}

func (p *profiler) StartMonitoring(name handler) {
	p.monitoringQueue.PushBack(name)
	m, err := p.getMonitoring(name)
	if err != nil {
		logrus.Fatal(err)
	}
	m.currentTime = time.Now()
}

func (p *profiler) EndMonitoring() {
	b := p.monitoringQueue.Back()
	m, err := p.getMonitoring(b.Value.(handler))
	if err != nil {
		logrus.Fatal(err)
	}
	p.monitoringQueue.Remove(b)

	select {
	case <-p.delay.C:
		for _, v := range p.handlers {
			var sum float64
			for _, t := range v.time {
				sum += t
			}
			v.avg = sum / float64(len(v.time))
			v.time = v.time[:0]
		}
	default:
	}

	m.time = append(m.time, float64(time.Since(m.currentTime).Seconds()/float64(time.Millisecond)))
}

func (p *profiler) String() string {
	var r string
	for _, v := range p.handlers {
		r += fmt.Sprintf("\n%s: %g", v.handler, v.avg)
	}
	return r
}

func UseProfiler() IProfiler {
	if instance == nil {
		instance = &profiler{
			handlers:        handlers{},
			delay:           time.NewTicker(time.Millisecond * 300),
			monitoringQueue: list.New(),
		}
	}
	return instance
}
