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

func (element Element) Get(property string) js.Value {

	return element.value.Get(property)
}

func (element Element) GetElement(property string) Element {

	return New(element.Get(property))
}

func (element Element) GetBool(property string) bool {

	return element.Get(property).Bool()
}

func (element Element) GetInt(property string) int {

	return element.Get(property).Int()
}

func (element Element) GetFloat(property string) float64 {

	return element.Get(property).Float()
}

func (element Element) GetString(property string) string {

	return element.Get(property).String()
}
