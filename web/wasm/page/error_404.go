package main

import (
	"syscall/js"
)

import (
	"dummy_ai/web/wasm/template"
)

func main() {

	h2 := js.Global().Get("document").Call("createElement", "h2")
	h2.Set("innerHTML", "Error 404 - Page Not Found")

	page := template.CreatePage()
	page.Call("appendChild", h2)
}
