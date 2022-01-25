package params

import (
	"flag"
)

var (
	debug                 = flag.Bool("debug", false, "Enables debug mode")
	withoutSound          = flag.Bool("without-sound", false, "Disables sound in game")
	disableConfigAutoSave = flag.Bool("disable-config-autosave", false, "Disables auto save of user set configs!")
	profilecpu            = flag.Bool("profilecpu", false, "Enables or disables cpu profiler")
)

func IsDebug() bool {
	return *debug
}

func IsWithoutSound() bool {
	return *withoutSound
}

func SetWithoutSoundManually(s bool) {
	*withoutSound = s
}

func IsDisableConfigAutoSave() bool {
	return *disableConfigAutoSave
}

func IsProfileCPU() bool {
	return *profilecpu
}
