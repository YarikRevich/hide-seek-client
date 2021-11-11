package latency

import (
	"time"
)

type timingIndex struct {
	state int
	time time.Duration 
}

type Timings struct {
	loopDelay  *time.Ticker
	timings map[timingIndex]*struct {
		ticker    *time.Ticker
		callback func()
	}
}

func (t *Timings) ExecEach(c func(), s int, d time.Duration) {
	i := timingIndex{s, d}
	if _, ok := t.timings[i]; !ok {
		t.timings[i] = &struct {
			ticker    *time.Ticker
			callback func()
		}{
			ticker:    time.NewTicker(d),
			callback: c,
		}
	}
}

func (t *Timings) start() {
	go func() {
		for {
			for range t.loopDelay.C {
				for _, v := range t.timings {
					for range v.ticker.C {
						v.callback()
					}
				}
			}
		}
	}()
}

func NewTimings() *Timings {
	t := &Timings{
		timings: make(map[timingIndex]*struct {
			ticker    *time.Ticker
			callback func()
		}),
		loopDelay: time.NewTicker(time.Millisecond * 200),
	}
	t.start()
	return t
}
