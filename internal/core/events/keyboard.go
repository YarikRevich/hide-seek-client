package events

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// type KeyBoard struct {
// 	movementType direction
// }

// import "github.com/hajimehoshi/ebiten/v2"

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

//Checks if any keyboard key pressed
func (k *KeyBoard) IsAnyKeyPressed()bool{
	return len(inpututil.PressedKeys()) != 0
}

func NewKeyBoard() *KeyBoard{
	return new(KeyBoard)
}