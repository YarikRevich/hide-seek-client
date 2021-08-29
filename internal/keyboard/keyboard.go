package keyboard

import (
	"github.com/YarikRevich/Hide-Seek-with-Guns/internal/status"
)

func Process() {
	switch status.GetInstance().GetState() {
	case status.START_MENU:
	case status.SETTINGS_MENU:
	case status.CREATE_LOBBY_MENU:
	case status.JOIN_LOBBY_MENU:
	case status.CHOOSE_EQUIPMENT:
	case status.WAIT_ROOM:
	case status.GAME:
	}
}
