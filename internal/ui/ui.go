package ui

import (
	"github.com/YarikRevich/Hide-Seek-with-Guns/internal/status"
	"github.com/YarikRevich/Hide-Seek-with-Guns/internal/ui/start_menu"
)

// 	//Inits stage with corrisponding params
// 	Init(*Window.WindowConfig, *States.States, *Users.User, Map.MapConf)

// 	//Method which works for networking getting the newest comming info
// 	ProcessNetworking()

// 	//Method which works to process pressed buttons
// 	ProcessKeyboard()

// 	//Method which works to process input from keyboard
// 	ProcessTextInput()

// 	//Method which works to draw all the announcements in scene
// 	DrawAnnouncements()

// 	//Method which works to draw all the other elements
// 	DrawElements()

// 	//Method which runs chosen scene
// 	Run()
// }

func Process() {
	switch status.GetInstance().GetState() {
	case status.START_MENU:
		start_menu.Exec()
	case status.SETTINGS_MENU:
	case status.CREATE_LOBBY_MENU:
	case status.JOIN_LOBBY_MENU:
		
	case status.CHOOSE_EQUIPMENT:
	case status.WAIT_ROOM:
	case status.GAME:

	}
}
