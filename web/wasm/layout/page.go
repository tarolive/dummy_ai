package layout

import (
	"syscall/js"
)

var (
	Window   = js.Global()
	Document = Window.Get("document")
)

var (
	Head = Document.Get("head")
	Body = Document.Get("body")
)

func NewPage(page js.Value) {

	Body.Call("appendChild", page)
}
