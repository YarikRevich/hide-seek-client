package latency

func Seq(callbacks ...func()) {
	for _, v := range callbacks {
		v()
	}
}
