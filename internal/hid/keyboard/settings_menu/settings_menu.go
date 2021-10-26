package settingsmenu

import (
	"fmt"

	buffercollection "github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard/buffers/collection"
	"github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard/buffers/common"
	// "github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard/collection"
	"github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard/handler"

	"github.com/hajimehoshi/ebiten/v2"
)

func Exec() {
	handler.HandleKeyboardPress(buffercollection.SettingsMenuNameBuffer, []handler.PipelineEntity{
		{Keys: []ebiten.Key{ebiten.KeyBackspace}, Callback: func(b common.IBuffer, k rune) {
			b.Pop()
		}},
		{Keys: collection.UserKeys, Callback: func(b common.IBuffer, k rune) {
			b.Push(k)
		}},
	})
}
