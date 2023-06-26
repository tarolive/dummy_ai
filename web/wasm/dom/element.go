package dom

import (
	"syscall/js"
)

type Element struct {
	value js.Value
}

func New(value js.Value) Element {

	return Element{
		value: value,
	}
}

func Global() Element {

	return New(js.Global())
}

func (element Element) Get(property string) Element {

	return New(element.value.Get(property))
}
