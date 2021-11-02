package events

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

//Keyboard entity which is used for
//key handler with further callback
type KeyBoardEntity struct {

	//Describes single keys like "T", "A"...
	SingleKeys []ebiten.Key

	//Describes combined keys like: "CTRL+V" or "CTRL+C"
	Combination struct {
		AwaitKey, ControlKey ebiten.Key
	}
	Pressed  bool
	Callback func(IBuffer, rune)
}

type KeyBoard struct {
	awaitKeyTimer time.Time
}

func (b *KeyBoard) IsKeyInList(k ebiten.Key, l []ebiten.Key) bool {
	for _, v := range l {
		if v == k {
			return true
		}
	}
	return false
}

//Checks if pressed key is service
func (b *KeyBoard) IsServiceKey(k ebiten.Key) bool {
	return len(k.String()) != 1
}

func (k *KeyBoard) HandleKeyPress(b IBuffer, ke []KeyBoardEntity) {
	b.CleanBlinking()

	for _, pk := range inpututil.PressedKeys() {
	entities:
		for _, e := range ke {
			if pk == e.Combination.AwaitKey {
				k.awaitKeyTimer = time.Now().Add(time.Millisecond * 700)
			}
			if inpututil.KeyPressDuration(pk) == 1 && pk == e.Combination.ControlKey {
				if time.Since(k.awaitKeyTimer) <= 0 {
					e.Callback(b, '0')
					break entities
				}
			}

			if inpututil.KeyPressDuration(pk) == 1 || (e.Pressed && inpututil.KeyPressDuration(pk)%11 == 0) {
				if k.IsKeyInList(pk, e.SingleKeys) {
					if k.IsServiceKey(pk) {
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

	b.UpdateCursorBlink()
}

//Checks if any keyboard key pressed
func (k *KeyBoard) IsAnyKeyPressed() bool {
	return len(inpututil.PressedKeys()) != 0
}

func NewKeyBoard() *KeyBoard {
	return new(KeyBoard)
}
