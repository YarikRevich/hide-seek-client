package cli

import "flag"

var (
	debug        = flag.Bool("debug", false, "Enables debug mode")
	withoutSound = flag.Bool("without-sound", false, "Disables sound in game")
)

func IsDebug() bool {
	return *debug
}

func IsWithoutSound() bool {
	return *withoutSound
}
