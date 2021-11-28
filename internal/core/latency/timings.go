package latency

import (
	"fmt"
	"time"
)

type timingIndex struct {
	state int
	time  time.Duration
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
		if k.state == s {
			v.close <- 1
			delete(t.timingsEach, k)
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

func (t *Timings) ExecEach(c func(), s int, d time.Duration) {
	i := timingIndex{s, d}
	if _, ok := t.timingsEach[i]; !ok {
		t.timingsEach[i] = &struct {
			ticker   *time.Ticker
			callback func()
			close    chan int
		}{
			ticker:   time.NewTicker(d),
			callback: c,
			close:    make(chan int, 1),
		}
	}
}

func (t *Timings) start() {
	go func() {
		for {
			for range t.loopDelay.C {
				for i, v := range t.timingsEach {
					fmt.Println(i.state)
					select {
					case <-v.close:
						fmt.Println("BEFORE CALL BACK CLOSE")
						v.callback()
						fmt.Println("AFTER CALL BACK CLOSE")
						// continue
					case <-v.ticker.C:
						v.callback()
						continue

					case <-t.loopDelay.C:
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
