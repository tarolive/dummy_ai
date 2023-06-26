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

func (element Element) Set(property string, value any) {

	element.value.Set(property, value)
}

func (element Element) SetElement(property string, value Element) {

	element.Set(property, value)
}

func (element Element) SetBool(property string, value bool) {

	element.Set(property, value)
}

func (element Element) SetInt(property string, value int) {

	element.Set(property, value)
}

func (element Element) SetFloat(property string, value float64) {

	element.Set(property, value)
}

func (element Element) SetString(property string, value string) {

	element.Set(property, value)
}

func (element Element) Call(method string, args ...any) js.Value {

	return element.value.Call(method, args)
}
