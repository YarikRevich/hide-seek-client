//go:build (linux && ignore) || (bsd && ignore) || darwin
// +build linux,ignore bsd,ignore darwin

package paths

const (
	GAME_LOG_DIR     = "/usr/local/share/games/HideSeek/log"
	GAME_STORAGE_DIR = "/usr/local/share/games/HideSeek/db"
)
