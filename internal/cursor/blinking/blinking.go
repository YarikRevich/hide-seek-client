package blinking

import (
	"time"

	"github.com/YarikRevich/HideSeek-Client/internal/cursor/collection"
	"github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard/buffers/common"
)

var ticker = time.NewTicker(time.Second)

var blinkPosition rune

func SetCursorBlink(b common.IBuffer){
	select {
	case <- ticker.C:
		if blinkPosition == collection.BlinkingOn{
			blinkPosition = collection.BlinkingOff
		}else{
			blinkPosition = collection.BlinkingOn
		}
	default:
	}

	l := b.Last()
	if l == collection.BlinkingOn || l == collection.BlinkingOff{
		b.Pop()
		b.Push(blinkPosition)
	}else{
		b.Push(blinkPosition)
	}
}