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

	createTitle()
	createPage(page)
}

func createTitle() {

	title := document.Call("createElement", "title")
	title.Set("innerHTML", "DummyAI")

	head.Call("appendChild", title)
}

func createPage(page js.Value) {

	body.Call("appendChild", page)
}
