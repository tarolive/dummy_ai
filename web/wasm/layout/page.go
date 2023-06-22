package layout

import (
	"syscall/js"
)

import (
	"dummy_ai/web/wasm/dom"
)

func NewPage(page js.Value) {

	dom.Body.Call("appendChild", page)
}
