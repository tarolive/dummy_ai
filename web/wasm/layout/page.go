package layout

import (
	"syscall/js"
)

func NewPage(page js.Value) {

	body := js.Global().Get("document").Get("body")
	body.Call("appendChild", page)
}
