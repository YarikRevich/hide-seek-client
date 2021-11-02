package settingsmenu

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	"github.com/YarikRevich/HideSeek-Client/internal/core/keycodes"
	"github.com/atotto/clipboard"
	"github.com/sirupsen/logrus"

	"github.com/hajimehoshi/ebiten/v2"
)

func Exec() {
	events.UseEvents().Keyboard().HandleKeyPress(events.UseEvents().Input().SettingsMenuNameBuffer, []events.KeyBoardEntity{
		{CombinedKeys: []ebiten.Key{ebiten.KeyMetaLeft, ebiten.KeyV}, Callback: func(b events.IBuffer, k rune) {
			r, err := clipboard.ReadAll()
			if err != nil {
				logrus.Fatal(err)
			}
			for _, v := range r {
				b.Push(v)
			}
		}},
		{SingleKeys: []ebiten.Key{ebiten.KeyBackspace}, Pressed: true, Callback: func(b events.IBuffer, k rune) {
			b.Pop()
		}},
		{SingleKeys: keycodes.UserKeys, Callback: func(b events.IBuffer, k rune) {
			b.Push(k)
		}},
	})
}
