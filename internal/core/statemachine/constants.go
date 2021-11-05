package statemachine

const (
	AUDIO_DONE = iota
	AUDIO_UNDONE
)

const (
	INPUT_EMPTY = iota
	INPUT_SETTINGS_MENU_USERNAME
	INPUT_GAME
)

const (
	NETWORKING_OFFLINE = iota
	NETWORKING_ONLINE
	NETWORKING_WAIT_CONNECTION
)

const (
	DIAL_LAN = iota
	DIAL_WAN
)

const (
	UI_START_MENU = iota
	UI_SETTINGS_MENU

	UI_MAP_CHOOSE
	UI_HERO_CHOOSE

	UI_WAIT_ROOM
	UI_JOIN_LOBBY_MENU

	UI_CHOOSE_EQUIPMENT

	UI_GAME
)

