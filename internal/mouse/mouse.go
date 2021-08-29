package mouse

import (
	creationlobbymenu "github.com/YarikRevich/Hide-Seek-with-Guns/internal/mouse/creation_lobby_menu"
	"github.com/YarikRevich/Hide-Seek-with-Guns/internal/status"
)

func Process() {
	switch status.GetInstance().GetState(){
	case status.CREATE_LOBBY_MENU:
		creationlobbymenu.Exec()
	}
}