package template

import (
	"syscall/js"
)

const (
	ENGLISH    = "en"
	SPANISH    = "es"
	PORTUGUESE = "pt"
)

func Language() string {

	return ENGLISH
}

func SetLanguage(language string) {

	js.Global().Get("localStorage").Call("setItem", "language", language)
}
