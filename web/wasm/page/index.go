package main

import (
	"syscall/js"
)

import (
	"dummy_ai/web/wasm/template"
)

var (
	_window   = js.Global()
	_document = _window.Get("document")
)

func main() {

	page := template.CreateTemplate()
	page.Call("appendChild", createSearchForm())
}

func createSearchForm() js.Value {

	searchForm := _document.Call("createElement", "form")
	return searchForm
}
