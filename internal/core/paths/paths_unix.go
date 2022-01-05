//go:build linux || darwin
// +build linux darwin

package paths

const (
	GAME_LOG_DIR     = "/usr/local/share/games/HideSeek/log"
	GAME_STORAGE_DIR = "/usr/local/share/games/HideSeek/db"
	GAME_PPROF_DIR   = "/usr/local/share/games/HideSeek/pprof"
)
