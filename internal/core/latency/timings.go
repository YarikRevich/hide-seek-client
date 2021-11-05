package latency

import (
	"sync"
	"time"
)

type Timings struct {
	timings map[time.Duration]*struct {
		ticker    *time.Ticker
		once      *sync.Once
		callbacks []func()
	}
}

func (t *Timings) ExecEach(c func(), d time.Duration) {
	if _, ok := t.timings[d]; !ok {
		t.timings[d] = &struct {
			ticker    *time.Ticker
			once      *sync.Once
			callbacks []func()
		}{
			ticker:    time.NewTicker(d),
			once:      new(sync.Once),
			callbacks: []func(){c},
		}
	} else {
		l := t.timings[d]
		l.callbacks = append(l.callbacks, c)
	}
}

func (t *Timings) start() {
	go func() {
		for _, v := range t.timings {
			v.once.Do(func() {
				for range v.ticker.C {
					for _, c := range v.callbacks {
						c()
					}
				}
			})
		}
	}()
}

func NewTimings() *Timings {
	t := new(Timings)
	t.start()
	return t
}
