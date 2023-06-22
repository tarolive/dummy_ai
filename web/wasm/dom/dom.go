package dom

import (
	"syscall/js"
)

var (
	Window = js.Global()
)

var (
	Document        = Window.Get("document")
	DocumentElement = Window.Get("documentElement")
)

var (
	Head = Document.Get("head")
	Body = Document.Get("body")
)
