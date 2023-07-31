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

var (
	_messages = map[string]map[string]string{
		template.ENGLISH: {
			"searchButton.innerHTML":  "Search",
			"searchInput.placeholder": "",
		},
		template.SPANISH: {
			"searchButton.innerHTML":  "Buscar",
			"searchInput.placeholder": "",
		},
		template.PORTUGUESE: {
			"searchButton.innerHTML":  "Buscar",
			"searchInput.placeholder": "",
		},
	}[template.Language]
)

func main() {

	page := template.CreateTemplate()
	page.Call("appendChild", createSearchForm())
}

func createSearchForm() js.Value {

	searchForm := _document.Call("createElement", "form")
	searchForm.Call("appendChild", createSearchInput())
	searchForm.Call("appendChild", createSearchButton())

	searchFormStyle := searchForm.Get("style")
	searchFormStyle.Set("display" /* */, "flex")

	return searchForm
}

func createSearchInput() js.Value {

	searchInput := _document.Call("createElement", "input")
	searchInput.Set("type", "search")
	searchInput.Set("placeholder", _messages["searchInput.placeholder"])

	searchInputStyle := searchInput.Get("style")
	searchInputStyle.Set("flex" /*         */, 1)
	searchInputStyle.Set("margin-right" /* */, "var(--rh-space-md)")
	searchInputStyle.Set("font-family" /*  */, "var(--rh-font-family-body-text)")
	searchInputStyle.Set("font-size" /*    */, "var(--rh-font-size-body-text-md)")

	return searchInput
}

func createSearchButton() js.Value {

	searchButton := _document.Call("createElement", "rh-button")
	searchButton.Set("innerHTML", _messages["searchButton.innerHTML"])

	return searchButton
}
