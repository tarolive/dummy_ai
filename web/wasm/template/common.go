package template

import (
	"syscall/js"
)

var (
	window = js.Global()
)

func Window() js.Value {

	return window
}
