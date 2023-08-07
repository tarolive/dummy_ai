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

	template.Page.Call("appendChild", _createSearchBar())
}

func _createSearchBar() js.Value {

	searchBar := _document.Call("createElement", "div")
	searchBar.Call("appendChild", _createSearchBarInput())
	searchBar.Call("appendChild", _createSearchBarButton())

	searchBarStyle := searchBar.Get("style")
	searchBarStyle.Set("display" /* */, "flex")
	searchBarStyle.Set("gap" /*     */, "var(--rh-space-md)")

	return searchBar
}

func _createSearchBarInput() js.Value {

	searchInput := _document.Call("createElement", "input")
	searchInput.Set("type", "search")
	searchInput.Set("placeholder", _messages["searchInput.placeholder"])

	searchInputStyle := searchInput.Get("style")
	searchInputStyle.Set("flex" /*         */, 1)
	searchInputStyle.Set("padding" /*      */, "var(--rh-space-sm) var(--rh-space-md)")
	searchInputStyle.Set("font-family" /*  */, "var(--rh-font-family-body-text)")
	searchInputStyle.Set("font-size" /*    */, "var(--rh-font-size-body-text-md)")
	searchInputStyle.Set("border-style" /* */, "solid")
	searchInputStyle.Set("border-width" /* */, "var(--rh-border-width-sm)")
	searchInputStyle.Set("border-color" /* */, "#f0f0f0 #f0f0f0 #8a8d90 #f0f0f0")

	return searchInput
}

func _createSearchBarButton() js.Value {

	searchButton := _document.Call("createElement", "rh-button")
	searchButton.Set("innerHTML", _messages["searchButton.innerHTML"])

	return searchButton
}
