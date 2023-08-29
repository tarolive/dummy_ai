package template

import (
	"syscall/js"
)

var (
	masthead        = createMasthead()
	pageMainSection = createPageMainSection()
	pageMain        = createPageMain()
	page            = createPage()
)

func init() {

	js.Global().Get("document").Get("documentElement").Set("lang", language)
	js.Global().Get("document").Get("body").Call("appendChild", page)
}

func Masthead() js.Value {

	return masthead
}

func PageMainSection() js.Value {

	return pageMainSection
}

func PageMain() js.Value {

	return pageMain
}

func Page() js.Value {

	return page
}

func createMasthead() js.Value {

	masthead := js.Global().Get("document").Call("createElement", "header")
	masthead.Get("classList").Call("add", "pf-v5-c-masthead")

	return masthead
}

func createPageMainSection() js.Value {

	pageMainSection := js.Global().Get("document").Call("createElement", "section")
	pageMainSection.Get("classList").Call("add", "pf-v5-c-page__main-section")

	return pageMainSection
}

func createPageMain() js.Value {

	pageMain := js.Global().Get("document").Call("createElement", "main")
	pageMain.Get("classList").Call("add", "pf-v5-c-page__main")
	pageMain.Set("tabIndex", -1)
	pageMain.Call("appendChild", pageMainSection)

	return pageMain
}

func createPage() js.Value {

	page := js.Global().Get("document").Call("createElement", "div")
	page.Get("classList").Call("add", "pf-v5-c-page")
	page.Call("appendChild", masthead)
	page.Call("appendChild", pageMain)

	return page
}
