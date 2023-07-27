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
	searchForm.Call("appendChild", createSearchInput())
	searchForm.Call("appendChild", createSearchButton())

	return searchForm
}

func createSearchInput() js.Value {

	searchInput := _document.Call("createElement", "input")
	return searchInput
}

func createSearchButton() js.Value {

	searchButton := _document.Call("createElement", "rh-button")
	return searchButton
}
