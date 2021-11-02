package events

import (
	"fmt"
	"time"

	"github.com/YarikRevich/caching/pkg/zeroshifter"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

//Keyboard entity which is used for
//key handler with further callback
type KeyBoardEntity struct {

	//Describes single keys like "T", "A"...
	SingleKeys []ebiten.Key

	//Describes combined keys like: "CTRL+V" or "CTRL+C"
	CombinedKeys []ebiten.Key
	Pressed      bool
	Callback     func(IBuffer, rune)
}

type KeyBoard struct {
	queue       zeroshifter.IZeroShifter
	combinationTimer time.Time
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
			if len(k.queue.Get()) == 2 {
				e.Callback(b, '0')
				k.queue.Clean()
			}

			if inpututil.KeyPressDuration(pk) == 1 && len(e.CombinedKeys) != 0 {

			queueCheck:
				for i, v := range k.queue.Get() {
					if v != e.CombinedKeys[i] {
						k.queue.Clean()
						break queueCheck
					}
				}

				for _, v := range e.CombinedKeys {
					if v == pk {
						// if len(k.queue.Get()) > 0 {
						// 	if k.queue.Get()[0] == pk {
						// 		k.queue.Clean()
						// 	}
						// }
						fmt.Println(pk, k.queue.Get())
						k.queue.Add(pk)
						// k.queueAwaits++
						break entities

					}
				}

			}

			if inpututil.KeyPressDuration(pk) == 1 || (e.Pressed && inpututil.KeyPressDuration(pk)%12 == 0) {
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
	return &KeyBoard{queue: zeroshifter.New(2)}
}
