package collection

import (
	"sync"

	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/ui"
)

var OnceCollection = map[int]*sync.Once{
	ui.WAIT_ROOM: new(sync.Once),
	ui.GAME: new(sync.Once),
}
