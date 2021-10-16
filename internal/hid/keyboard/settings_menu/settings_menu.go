package settingsmenu

import (
	"fmt"

	buffercollection "github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard/buffers/collection"
	"github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard/buffers/common"
	"github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard/collection"
	"github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard/handler"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func Exec() {
	fmt.Println(inpututil.PressedKeys())
	handler.HandleKeyboardPress(buffercollection.SettingsMenuNameBuffer, []handler.PipelineEntity{
		{Keys: []ebiten.Key{ebiten.KeyBackspace}, Callback: func(b common.IBuffer, k rune) {
			b.Pop()
		}},
		{Keys: collection.UserKeys, Callback: func(b common.IBuffer, k rune) {
			b.Push(k)
		}},
	})
}
