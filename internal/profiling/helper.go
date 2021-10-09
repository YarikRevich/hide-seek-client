package profiling

func GetMonitoringByHandler(name handler, handlers handlers)*monitoring{
	for _, v := range handlers{
		if v != nil && v.handler == name{
			return v
		}
	}
	return nil
}