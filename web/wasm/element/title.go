package element

import (
	"syscall/js"
)

import (
	"dummy_ai/web/wasm/dom"
)

func NewTitle(title string) js.Value {

	titleElement := dom.Document.Call("createElement", "title")
	titleElement.Set("innerHTML", title)

	return titleElement
}
