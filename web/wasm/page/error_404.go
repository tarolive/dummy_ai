package main

import (
	"syscall/js"
)

import (
	"dummy_ai/web/wasm/layout"
)

func main() {

	document := js.Global().Get("document")

	h2 := document.Call("createElement", "h2")
	h2.Set("innerHTML", "Error 404 - Page Not Found")

	layout.NewPage(h2)
}
