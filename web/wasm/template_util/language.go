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
	_navigator    = _window.Get("navigator")
)

func Language() string {

	if language := _localStorage.Call("getItem", "language").String(); isSupportedLanguage(language) {

		return language
	}

	if language := _navigator.Get("language").String(); language != "" {

		if len(language) > 2 {

			language = language[:2]
		}

		if isSupportedLanguage(language) {

			SetLanguage(language)
			return language
		}
	}

	language := ENGLISH

	SetLanguage(language)
	return language
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
