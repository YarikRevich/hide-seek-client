package latency

import "sync"

//Asynchronyse sequence
func Seq(callbacks ...func(*sync.WaitGroup)) {
	var s sync.WaitGroup
	for _, v := range callbacks {
		s.Add(1)
		go v(&s)
		s.Wait()
	}
}
