package text

import "github.com/YarikRevich/HideSeek-Client/internal/buffers"

var (
	buff buffers.TextBuffer
)

type Buff struct {
	value []rune
}

// func (t *Buff) Write(value string) {
// 	for _, v := range value{
// 		t.value = append(t.value, v)
// 	}
// }

func (t *Buff) Read() string {
	return string(t.value)
}

func (t *Buff) Last() rune {
	if len(t.value) != 0 {
		return t.value[0]
	}
	return '0'
}

func (t *Buff) Clean() {
	t.value = t.value[:0]
}

func (t *Buff) Push(v rune) {
	t.value = append(t.value, v)
}

func (t *Buff) Pop() {
	if len(t.value) != 0 {
		t.value = t.value[:len(t.value)-1]
	}
}

func UseBuffer() buffers.TextBuffer {
	if buff == nil {
		buff = new(Buff)
	}
	return buff
}
