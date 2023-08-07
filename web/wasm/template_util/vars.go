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
