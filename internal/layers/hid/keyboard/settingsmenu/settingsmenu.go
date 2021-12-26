package settingsmenu

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/events"
	"github.com/YarikRevich/hide-seek-client/internal/core/keycodes"
	"github.com/atotto/clipboard"
	"github.com/sirupsen/logrus"

	"github.com/hajimehoshi/ebiten/v2"
)

func Exec() {
	events.UseEvents().Keyboard().HandleKeyPress(events.UseEvents().Input().SettingsMenuNameBuffer, []events.KeyBoardEntity{
		{Combination: struct{ AwaitKey, ControlKey ebiten.Key }{AwaitKey: ebiten.KeyMetaLeft, ControlKey: ebiten.KeyV}, Callback: func(b events.IBuffer, k rune) {
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
