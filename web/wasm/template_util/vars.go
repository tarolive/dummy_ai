package template_util

import (
	"syscall/js"
)

var (
	Window       = js.Global()
	Document     = Window.Get("document")
	LocalStorage = Window.Get("localStorage")
	Navigator    = Window.Get("navigator")
)

var (
	DocumentElement = Document.Get("documentElement")
	Head            = Document.Get("head")
	Body            = Document.Get("body")
)
