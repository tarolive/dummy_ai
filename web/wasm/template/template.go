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
	return page
}
