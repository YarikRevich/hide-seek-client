package settingsmenu

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	"github.com/YarikRevich/HideSeek-Client/internal/core/keycodes"

	"github.com/hajimehoshi/ebiten/v2"
)

func Exec() {
	events.UseEvents().Keyboard().HandleKeyPress(events.UseEvents().Input().SettingsMenuNameBuffer, []events.KeyBoardEntity{
		{Keys: []ebiten.Key{ebiten.KeyBackspace}, Callback: func(b events.IBuffer, k rune) {
			b.Pop()
		}},
		{Keys: keycodes.UserKeys, Callback: func(b events.IBuffer, k rune) {
			b.Push(k)
		}},
	})
}
