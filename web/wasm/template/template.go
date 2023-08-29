package template

import (
	"syscall/js"
)

var (
	pageMainSection = createPageMainSection()
)

func init() {

	js.Global().Get("document").Get("documentElement").Set("lang", language)
	js.Global().Get("document").Get("body").Call("appendChild", pageMainSection)
}

func PageMainSection() js.Value {

	return pageMainSection
}

func createPageMainSection() js.Value {

	page := js.Global().Get("document").Call("createElement", "section")
	page.Get("classList").Call("add", "pf-v5-c-page__main-section")

	return page
}
