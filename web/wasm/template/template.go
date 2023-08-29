package template

import (
	"syscall/js"
)

var (
	pageMainSection = createPageMainSection()
	pageMain        = createPageMain()
)

func init() {

	js.Global().Get("document").Get("documentElement").Set("lang", language)
	js.Global().Get("document").Get("body").Call("appendChild", pageMain)
}

func PageMainSection() js.Value {

	return pageMainSection
}

func PageMain() js.Value {

	return pageMain
}

func createPageMainSection() js.Value {

	pageMainSection := js.Global().Get("document").Call("createElement", "section")
	pageMainSection.Get("classList").Call("add", "pf-v5-c-page__main-section")

	return pageMainSection
}

func createPageMain() js.Value {

	pageMain := js.Global().Get("document").Call("createElement", "main")
	pageMain.Get("classList").Call("add", "pf-v5-c-page__main")
	pageMain.Set("tabindex", -1)
	pageMain.Call("appendChild", pageMainSection)

	return pageMain
}
