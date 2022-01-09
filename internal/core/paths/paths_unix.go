//go:build linux || darwin
// +build linux darwin

package paths

import (
	"fmt"
	"os/user"
)

var GAME_LOG_DIR, GAME_STORAGE_DIR, GAME_PPROF_DIR string

func InitSystemPaths() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	GAME_LOG_DIR = fmt.Sprintf("/home/%s/games/HideSeek/log", user.Username)
	GAME_STORAGE_DIR = fmt.Sprintf("/home/%s/games/HideSeek/db", user.Username)
	GAME_PPROF_DIR = fmt.Sprintf("/home/%s/games/HideSeek/pprof", user.Username)
}
