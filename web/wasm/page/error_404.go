package main

import (
	"dummy_ai/web/wasm/dom"
	"dummy_ai/web/wasm/layout"
)

func main() {

	h2 := dom.Document.Call("createElement", "h2")
	h2.Set("innerHTML", "Error 404 - Page Not Found")

	layout.NewPage(h2)
}
