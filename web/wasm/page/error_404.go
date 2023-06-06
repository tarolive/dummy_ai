package main

import (
	"syscall/js"
)

func main() {

	document := js.Global().Get("document")

	h2 := document.Call("createElement", "h2")
	h2.Set("innerHTML", "Error 404 - Page Not Found")

	body := document.Get("body")
	body.Call("appendChild", h2)
}
