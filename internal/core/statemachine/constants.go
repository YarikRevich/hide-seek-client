package statemachine

const (
	AUDIO_DONE = iota
	AUDIO_UNDONE
)

const (
	INPUT_EMPTY = iota
	INPUT_SETTINGS_MENU_USERNAME
	INPUT_JOIN_MENU
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

	UI_JOIN_MENU

	UI_MAP_CHOOSE
	UI_HERO_CHOOSE

	UI_WAIT_ROOM_START
	UI_WAIT_ROOM_JOIN

	UI_CHOOSE_EQUIPMENT

	UI_GAME
)

const (
	UI_SETTINGS_MENU_CHECKBOX_ON = iota
	UI_SETTINGS_MENU_CHECKBOX_OFF
)

