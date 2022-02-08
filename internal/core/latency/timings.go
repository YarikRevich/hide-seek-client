package latency

import (
	"time"
)

const (
	Connectivity = iota
)

type timingIndex struct {
	// state int
	timingType int
	time       time.Duration
}

type Timings struct {
	loopDelay   *time.Ticker
	timingsEach map[timingIndex]*struct {
		ticker   *time.Ticker
		callback func()
		close    chan int
	}

	timingsFor map[timingIndex]*struct {
		after    <-chan time.Time
		callback func()
		end      func()
	}
}

func (t *Timings) CleanEachTimings(s int) {
	for k, v := range t.timingsEach {
		if k.timingType == s {
			go func() {
				v.close <- 1
			}()
		}
	}
}

func (t *Timings) ExecFor(c, e func(), s int, d time.Duration) {
	i := timingIndex{s, d}
	if _, ok := t.timingsFor[i]; !ok {
		t.timingsFor[i] = &struct {
			after         <-chan time.Time
			callback, end func()
		}{
			after:    time.After(d),
			callback: c,
			end:      e,
		}
	}
}

func (t *Timings) ExecEach(s int, d time.Duration, c func()) {
	i := timingIndex{s, d}
	if _, ok := t.timingsEach[i]; !ok {
		t.timingsEach[i] = &struct {
			ticker   *time.Ticker
			callback func()
			close    chan int
		}{
			ticker:   time.NewTicker(d),
			callback: c,
			close:    make(chan int),
		}
	}
}

func (t *Timings) start() {
	go func() {
		for {
			for range t.loopDelay.C {
				for i, v := range t.timingsEach {
					select {
					case <-v.close:
						delete(t.timingsEach, i)
						v.callback()
					case <-v.ticker.C:
						v.callback()
						continue
					}
				}
				for k, v := range t.timingsFor {
					select {
					case <-v.after:
						v.end()
						delete(t.timingsFor, k)
					default:
						v.callback()
					}
				}
			}
		}
	}()
}

func NewTimings() *Timings {
	t := &Timings{
		timingsEach: make(map[timingIndex]*struct {
			ticker   *time.Ticker
			callback func()
			close    chan int
		}),

		timingsFor: make(map[timingIndex]*struct {
			after         <-chan time.Time
			callback, end func()
		}),
		loopDelay: time.NewTicker(time.Millisecond * 200),
	}
	t.start()
	return t
}
