package main

import (
	"dummy_ai/web/wasm/dom"
	"dummy_ai/web/wasm/layout"
)

func main() {

	h2 := dom.Document.Call("createElement", "h2")
	h2.Set("innerHTML", "Hello World, DummyAI!")

	layout.NewPage(h2)
}
