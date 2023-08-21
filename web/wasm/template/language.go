package template

import (
	"syscall/js"
)

const (
	English    = "en"
	Spanish    = "es"
	Portuguese = "pt"
)

const (
	DefaultLanguage = English
)

var (
	language = Language()
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

	language := DefaultLanguage

	SetLanguage(language)
	return language
}

func SetLanguage(language string) {

	js.Global().Get("localStorage").Call("setItem", "language", language)
}

func isSupportedLanguage(language string) bool {

	switch language {

	case English, Spanish, Portuguese:

		return true

	default:

		return false

	}
}
