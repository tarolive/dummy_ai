package template

import (
	"syscall/js"
)

import (
	"dummy_ai/web/wasm/template_util"
)

var (
	Language = template_util.Language()
)

var (
	_window   = js.Global()
	_document = _window.Get("document")
)

func CreatePage() js.Value {

	html := _window.Get("documentElement")
	html.Set("lang", Language)

	page := _document.Call("createElement", "div")

	body := _document.Get("body")
	body.Call("appendChild", createNavigation())
	body.Call("appendChild", page)

	return page
}

func createNavigation() js.Value {

	navigation := _document.Call("createElement", "div")

	navigationStyle := navigation.Get("style")
	navigationStyle.Set("position" /* */, "fixed")
	navigationStyle.Set("top" /*      */, 0)
	navigationStyle.Set("right" /*    */, 0)
	navigationStyle.Set("left" /*     */, 0)
	navigationStyle.Set("height" /*   */, "72px")
	navigationStyle.Set("padding" /*  */, "0 var(--rh-space-xl)")

	return navigation
}
