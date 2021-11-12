package cli

import "flag"

var (
	debug        = flag.Bool("debug", false, "Enables debug mode")
	withoutSound = flag.Bool("without-sound", false, "Disables sound in game")
	disableConfigAutoSave = flag.Bool("disable-config-autosave", false, "Disables auto save of user set configs!")
)

func IsDebug() bool {
	return *debug
}

func IsWithoutSound() bool {
	return *withoutSound
}

func IsDisableConfigAutoSave()bool{
	return *disableConfigAutoSave
}
