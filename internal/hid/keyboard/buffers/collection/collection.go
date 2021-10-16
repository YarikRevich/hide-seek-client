package collection

import "github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard/buffers/common"

var (
	SettingsMenuNameBuffer = common.NewBuffer()
)

func CleanBlinking() {
	SettingsMenuNameBuffer.CleanBlinkingUnfocus()
}
