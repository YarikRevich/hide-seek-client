package cursor

import "time"

var ticker = time.NewTicker(time.Second)

func GetCursorBlink()rune{
	select {
	case <- ticker.C:
		return 'j'
	default:
		return 'b'
	}
}