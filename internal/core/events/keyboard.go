package events

import (
	"regexp"
	"strings"
	"time"

	"github.com/YarikRevich/hide-seek-client/internal/core/keycodes"
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
		AwaitKey   []ebiten.Key
		ControlKey ebiten.Key
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

func (k *KeyBoard) CleanPressedKey(key ebiten.Key) string {
	return regexp.MustCompile(strings.Join(keycodes.ServiceKeyPrefixes, "|")).ReplaceAllString(key.String(), "")
}

func (k *KeyBoard) HandleKeyPress(b IBuffer, ke []KeyBoardEntity) {
	b.CleanBlinking()

	for _, pk := range inpututil.PressedKeys() {
	entities:
		for _, e := range ke {
		awaitkey:
			for _, v := range e.Combination.AwaitKey {
				if v == pk {
					k.awaitKeyTimer = time.Now().Add(time.Millisecond * 700)
					break awaitkey
				}
			}
			if inpututil.KeyPressDuration(pk) == 1 && pk == e.Combination.ControlKey {
				if time.Since(k.awaitKeyTimer) <= 0 {
					e.Callback(b, '0')
					break entities
				}
			}

			if inpututil.KeyPressDuration(pk) == 1 || (e.Pressed && inpututil.KeyPressDuration(pk)%15 == 0) {
				if k.IsKeyInList(pk, e.SingleKeys) {
					for _, k := range k.CleanPressedKey(pk) {
						e.Callback(b, k)
					}
					break entities
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

func (k *KeyBoard) AreKeysCombinedInOrder(m, s ebiten.Key) bool {
	return ebiten.IsKeyPressed(m) == true && ebiten.IsKeyPressed(s) == true && inpututil.KeyPressDuration(m) > inpututil.KeyPressDuration(s)
}

func NewKeyBoard() *KeyBoard {
	return new(KeyBoard)
}
