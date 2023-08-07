package template

import (
	"syscall/js"
)

var (
	Window = js.Global()
)

var (
	Document        = Window.Get("document")
	DocumentElement = Document.Get("documentElement")
	Head            = Document.Get("head")
	Body            = Document.Get("body")
)

var (
	SessionStorage = Window.Get("sessionStorage")
	LocalStorage   = Window.Get("localStorage")
)

var (
	Navigator = Window.Get("navigator")
)
