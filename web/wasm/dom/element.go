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
