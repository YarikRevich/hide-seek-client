package ingame

import (
	"fmt"
	"time"

	"github.com/YarikRevich/hide-seek-client/internal/core/profiling/common"
)

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
