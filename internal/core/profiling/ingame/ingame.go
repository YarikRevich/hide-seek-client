package ingame

import (
	"fmt"
	"time"

	"github.com/YarikRevich/hide-seek-client/internal/core/profiling/common"
)

// type handlers []*monitoring

// type monitoring struct {
// 	handler     handler
// 	currentTime time.Time
// 	time        []float64
// 	avg         float64
// }

// type profiler struct {
// 	delay           *time.Ticker
// 	handlers        handlers
// 	monitoringQueue *list.List
// }

// func (p *profiler) getMonitoring(name handler) (*monitoring, error) {
// 	for _, v := range p.handlers {
// 		if v != nil && v.handler == name {
// 			return v, nil
// 		}
// 	}
// 	return nil, errors.New("monitoring handler was not found")
// }

// func (p *profiler) StartMonitoring(name handler) {
// 	p.monitoringQueue.PushBack(name)
// 	m, err := p.getMonitoring(name)
// 	if err != nil {
// 		logrus.Fatal(err)
// 	}
// 	m.currentTime = time.Now()
// }

// func (p *profiler) EndMonitoring() {
// 	b := p.monitoringQueue.Back()
// 	m, err := p.getMonitoring(b.Value.(handler))
// 	if err != nil {
// 		logrus.Fatal(err)
// 	}
// 	p.monitoringQueue.Remove(b)

// 	select {
// 	case <-p.delay.C:
// 		for _, v := range p.handlers {
// 			var sum float64
// 			for _, t := range v.time {
// 				sum += t
// 			}
// 			v.avg = sum / float64(len(v.time))
// 			v.time = v.time[:0]
// 		}
// 	default:
// 	}

// 	m.time = append(m.time, float64(time.Since(m.currentTime).Seconds()/float64(time.Millisecond)))
// }

// func (p *profiler) String() string {
// 	var r string
// 	for _, v := range p.handlers {
// 		r += fmt.Sprintf("\n%s: %g", v.handler, v.avg)
// 	}
// 	return r
// }

// func UseProfiler() common.Profiler {
// 	if instance == nil {
// 		instance = &profiler{
// 			handlers:        handlers{},
// 			delay:           time.NewTicker(time.Millisecond * 300),
// 			monitoringQueue: list.New(),
// 		}
// 	}
// 	return instance
// }

const (
	RENDER        string = "render"
	UI            string = "ui"
	UI_START_MENU string = "ui_start_menu"
	UI_GAME_MENU  string = "ui_game_menu"
	UI_ANIMATION  string = "ui_animation"

	MOUSE    string = "mouse"
	KEYBOARD string = "keyboard"
	AUDIO    string = "audio"
)

var instance common.Profiler

type Profiler struct {
	profiles []*struct {
		Name      string
		Timestamp float64
	}
	delay *time.Ticker
}

func (p *Profiler) StartMonitoring(profiler ...string) {
	for _, v := range profiler {
		for _, q := range p.profiles {
			if q.Name == v {
				(*q).Timestamp = float64(time.Now().Nanosecond())
			}
		}
	}
}
func (p *Profiler) StopMonitoring(profiler ...string) {
	for _, v := range profiler {
		for _, q := range p.profiles {
			if q.Name == v {
				(*q).Timestamp = float64(time.Now().Nanosecond()) - (*q).Timestamp
			}
		}
	}
}
func (p *Profiler) Show() string {
	var r string
	for _, v := range p.profiles {
		r += fmt.Sprintf("\n%s: %g", v.Name, v.Timestamp)
	}
	return r
}

func UseProfiler() common.Profiler {
	if instance == nil {
		instance = &Profiler{
			delay: time.NewTicker(time.Second),
			profiles: []*struct {
				Name      string
				Timestamp float64
			}{
				{Name: UI},
				{Name: MOUSE},
				{Name: AUDIO},
				{Name: RENDER},
				{Name: KEYBOARD},
				{Name: UI_GAME_MENU},
				{Name: UI_ANIMATION},
				{Name: UI_START_MENU},
			},
		}
	}
	return instance
}
