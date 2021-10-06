package cursor

import (
	"time"

	"github.com/YarikRevich/HideSeek-Client/internal/buffers"
	"github.com/YarikRevich/HideSeek-Client/internal/buffers/text"
)

var ticker = time.NewTicker(time.Second)

const (
	on = '|'
	off = ' '
)

var blinkPosition rune

func SetCursorBlink(buffers.TextBuffer){
	select {
	case <- ticker.C:
		if blinkPosition == on{
			blinkPosition = off
		}else{
			blinkPosition = on
		}
	default:
	}

	b := text.UseBuffer()
	l := b.Last()
	if l == on || l == off{
		b.Pop()
		b.Push(blinkPosition)
	}else{
		b.Push(blinkPosition)
	}
}