package collection

import "github.com/hajimehoshi/ebiten/v2"

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

func IsKeyInList(k ebiten.Key, l []ebiten.Key) bool {
	for _, v := range l {
		if v == k {
			return true
		}
	}
	return false
}

func IsServiceKey(k ebiten.Key) bool {
	return len(k.String()) != 1
}
