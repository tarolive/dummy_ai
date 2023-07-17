package template_util

import (
	"syscall/js"
)

const (
	ENGLISH    = "en"
	SPANISH    = "es"
	PORTUGUESE = "pt"
)

var (
	_window       = js.Global()
	_localStorage = _window.Get("localStorage")
)

func Language() string {

	if language := _localStorage.Call("getItem", "language").String(); isSupportedLanguage(language) {

		return language
	}

	return ENGLISH
}

func SetLanguage(language string) {

	_localStorage.Call("setItem", "language", language)
}

func isSupportedLanguage(language string) bool {

	switch language {

	case ENGLISH:
	case SPANISH:
	case PORTUGUESE:

		return true
	}

	return false
}
