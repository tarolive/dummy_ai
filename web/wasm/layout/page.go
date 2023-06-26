package layout

import (
	"syscall/js"
)

func NewPage(page js.Value) {

	js.Global().Get("document").Get("body").Call("appendChild", page)
}
