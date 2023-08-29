package template

import (
	"syscall/js"
)

var (
	page = createPage()
)

func init() {

	js.Global().Get("document").Get("documentElement").Set("lang", language)
	js.Global().Get("document").Get("body").Call("appendChild", page)
}

func Page() js.Value {

	return page
}

func createPage() js.Value {

	page := js.Global().Get("document").Call("createElement", "section")
	page.Get("classList").Call("add", "pf-v5-c-page__main-section")

	return page
}
