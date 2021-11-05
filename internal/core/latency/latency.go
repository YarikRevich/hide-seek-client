package latency

var instance *Latency

type Latency struct {
	timings *Timings
	once    *Once
}

func (l *Latency) Timings() *Timings {
	return l.timings
}

func (l *Latency) Once() *Once {
	return l.once
}

func UseLatency() *Latency {
	if instance == nil {
		instance = &Latency{
			timings: NewTimings(),
			once:    NewOnce(),
		}
	}
	return instance
}
