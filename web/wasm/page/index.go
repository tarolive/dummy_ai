package main

import (
	"syscall/js"
)

import (
	"dummy_ai/web/wasm/layout"
)

func main() {

	h2 := js.Global().Get("document").Call("createElement", "h2")
	h2.Set("innerHTML", "Hello World, DummyAI! Language: "+layout.Language())

	layout.CreatePage(h2)
}
