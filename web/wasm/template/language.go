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

	if language := js.Global().Get("localStorage").Call("getItem", "language").String(); isSupportedLanguage(language) {

		return language
	}

	if language := js.Global().Get("navigator").Get("language").String(); language != "" {

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

	js.Global().Get("localStorage").Call("setItem", "language", language)
}

func isSupportedLanguage(language string) bool {

	switch language {

	case ENGLISH, SPANISH, PORTUGUESE:

		return true

	default:

		return false

	}
}
