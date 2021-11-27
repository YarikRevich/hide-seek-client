package latency

import "time"

type timingIndex struct {
	state int
	time  time.Duration
}

type Timings struct {
	loopDelay   *time.Ticker
	timingsEach map[timingIndex]*struct {
		ticker   *time.Ticker
		callback func()
	}

	timingsFor map[timingIndex]*struct {
		after    <-chan time.Time
		callback func()
	}
}

func (t *Timings) CleanEachTimings(s int) {
	for k := range t.timingsEach {
		if k.state == s {
			delete(t.timingsEach, k)
		}
	}
}

func (t *Timings) ExecFor(c func(), s int, d time.Duration) {
	i := timingIndex{s, d}
	if _, ok := t.timingsFor[i]; !ok {
		t.timingsFor[i] = &struct {
			after    <-chan time.Time
			callback func()
		}{
			after:    time.After(d),
			callback: c,
		}
	}
}

func (t *Timings) ExecEach(c func(), s int, d time.Duration) {
	i := timingIndex{s, d}
	if _, ok := t.timingsEach[i]; !ok {
		t.timingsEach[i] = &struct {
			ticker   *time.Ticker
			callback func()
		}{
			ticker:   time.NewTicker(d),
			callback: c,
		}
	}
}

func (t *Timings) start() {
	go func() {
		for {
			for range t.loopDelay.C {
				for _, v := range t.timingsEach {
					select {
					case <-v.ticker.C:
						v.callback()
						continue
					case <-t.loopDelay.C:
					}
				}
				for k, v := range t.timingsFor {
					select {
					case <-v.after:
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
		}),

		timingsFor: make(map[timingIndex]*struct {
			after    <-chan time.Time
			callback func()
		}),
		loopDelay: time.NewTicker(time.Millisecond * 200),
	}
	t.start()
	return t
}
