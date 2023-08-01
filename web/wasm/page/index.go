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
			"searchInput.placeholder": "What are you looking for?",
		},
		template.SPANISH: {
			"searchButton.innerHTML":  "Buscar",
			"searchInput.placeholder": "¿Qué estás buscando?",
		},
		template.PORTUGUESE: {
			"searchButton.innerHTML":  "Buscar",
			"searchInput.placeholder": "O que você está buscando?",
		},
	}[template.Language]
)

func main() {

	page := template.CreateTemplate()
	page.Call("appendChild", createSearchBar())
}

func createSearchBar() js.Value {

	searchForm := _document.Call("createElement", "div")
	searchForm.Call("appendChild", createSearchBarInput())
	searchForm.Call("appendChild", createSearchBarButton())

	searchFormStyle := searchForm.Get("style")
	searchFormStyle.Set("display" /* */, "flex")

	return searchForm
}

func createSearchBarInput() js.Value {

	searchInput := _document.Call("createElement", "input")
	searchInput.Set("type", "search")
	searchInput.Set("placeholder", _messages["searchInput.placeholder"])

	searchInputStyle := searchInput.Get("style")
	searchInputStyle.Set("flex" /*         */, 1)
	searchInputStyle.Set("margin-right" /* */, "var(--rh-space-md)")
	searchInputStyle.Set("padding" /*      */, "var(--rh-space-sm) var(--rh-space-md)")
	searchInputStyle.Set("font-family" /*  */, "var(--rh-font-family-body-text)")
	searchInputStyle.Set("font-size" /*    */, "var(--rh-font-size-body-text-md)")
	searchInputStyle.Set("border" /*       */, "var(--rh-border-width-sm) solid")
	searchInputStyle.Set("border-color" /* */, "#f0f0f0 #f0f0f0 #8a8d90 #f0f0f0")

	return searchInput
}

func createSearchBarButton() js.Value {

	searchButton := _document.Call("createElement", "rh-button")
	searchButton.Set("innerHTML", _messages["searchButton.innerHTML"])

	return searchButton
}
