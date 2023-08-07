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
	window       = js.Global()
	localStorage = window.Get("localStorage")
	navigator    = window.Get("navigator")
)

func Language() string {

	if language := localStorage.Call("getItem", "language").String(); isSupportedLanguage(language) {

		return language
	}

	if language := navigator.Get("language").String(); language != "" {

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

	localStorage.Call("setItem", "language", language)
}

func isSupportedLanguage(language string) bool {

	switch language {

	case ENGLISH, SPANISH, PORTUGUESE:

		return true
	}

	return false
}
