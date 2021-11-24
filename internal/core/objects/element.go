package objects

type Element struct {
	Base
}

func NewElement() *Element{
	return new(Element)
}

