package text

import "github.com/YarikRevich/HideSeek-Client/internal/buffers"

var (
	buff buffers.TextBuffer
)

type Buff struct {
	value []rune
}

func (t *Buff) Write(value string) {
	for _, v := range value{
		t.value = append(t.value, v)
	}
}

func (t *Buff) Read()string{
	return string(t.value)
}

func (t *Buff) Clean(){
	t.value = t.value[:0]
}

func UseBuffer()buffers.TextBuffer{
	if buff == nil{
		buff = new(Buff)
	}
	return buff
}
