package events

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// type KeyBoard struct {
// 	movementType direction
// }

// import "github.com/hajimehoshi/ebiten/v2"

var UserKeys = []ebiten.Key{
	ebiten.KeyA,
	ebiten.KeyB,
	ebiten.KeyC,
	ebiten.KeyD,
	ebiten.KeyE,
	ebiten.KeyF,
	ebiten.KeyG,
	ebiten.KeyH,
	ebiten.KeyI,
	ebiten.KeyJ,
	ebiten.KeyK,
	ebiten.KeyL,
	ebiten.KeyM,
	ebiten.KeyN,
	ebiten.KeyO,
	ebiten.KeyP,
	ebiten.KeyQ,
	ebiten.KeyR,
	ebiten.KeyS,
	ebiten.KeyT,
	ebiten.KeyU,
	ebiten.KeyV,
	ebiten.KeyW,
	ebiten.KeyX,
	ebiten.KeyY,
	ebiten.KeyZ,

	ebiten.Key0,
	ebiten.Key1,
	ebiten.Key2,
	ebiten.Key3,
	ebiten.Key4,
	ebiten.Key5,
	ebiten.Key6,
	ebiten.Key7,
	ebiten.Key8,
	ebiten.Key9,
}

//Keyboard entity which is used for 
//key handler with further callback
type KeyBoardEntity struct {
	Keys     []ebiten.Key
	Callback func(IBuffer, rune)
}

type KeyBoard struct{}

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

func (k *KeyBoard) HandleKeyPress(b IBuffer, ke []KeyBoardEntity){
	b.CleanBlinking()

	for _, pk := range inpututil.PressedKeys() {
		if inpututil.KeyPressDuration(pk) == 1 {
			for _, e := range ke {
				if k.IsKeyInList(pk, e.Keys) {
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

func NewKeyBoard() *KeyBoard{
	return new(KeyBoard)
}