package handler

import (
	"github.com/YarikRevich/HideSeek-Client/internal/cursor/blinking"
	"github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard/buffers/common"
	"github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard/collection"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type PipelineEntity struct {
	Keys     []ebiten.Key
	Callback func(common.IBuffer, rune)
}

func HandleKeyboardPress(b common.IBuffer, pe []PipelineEntity) {
	b.CleanBlinking()

	for _, pk := range inpututil.PressedKeys() {
		if inpututil.KeyPressDuration(pk) == 1 {
			for _, e := range pe {
				if collection.IsKeyInList(pk, e.Keys) {
					if collection.IsServiceKey(pk) {
						e.Callback(b, '0')
						break
					}
					for _, k := range pk.String() {
						e.Callback(b, k)
					}
					break
				}
			}
		}
	}

	blinking.SetCursorBlink(b)
}
