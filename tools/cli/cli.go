package cli

import (
	"flag"
)

var debug = flag.Bool("debug", false, "Enables debug mode")

func GetDebug() bool {
	return *debug
}
