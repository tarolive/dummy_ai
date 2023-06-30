package layout

import (
	"syscall/js"
)

var (
	window = js.Global()
)

var (
	document = window.Get("document")
)

var (
	head = document.Get("head")
	body = document.Get("body")
)

func CreatePage(page js.Value) {

	pageContainer := document.Call("createElement", "div")
	pageContainer.Call("appendChild", page)

	body.Call("appendChild", pageContainer)
}
