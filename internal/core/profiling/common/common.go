package common

type Profiler interface {
	//Starts monitoring for passed profiler
	StartMonitoring(profile ...string)

	//Stops monitoring for passed profiler
	StopMonitoring(profile ...string)

	Show() string
}
