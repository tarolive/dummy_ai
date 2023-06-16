package layout

import (
	"syscall/js"
)

var (
	window   = js.Global()
	document = window.Get("document")
)

var (
	html = document.Get("documentElement")
	head = document.Get("head")
	body = document.Get("body")
)

func NewPage(page js.Value) {

	body.Call("appendChild", page)
}
