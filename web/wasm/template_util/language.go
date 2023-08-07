package template_util

const (
	ENGLISH    = "en"
	SPANISH    = "es"
	PORTUGUESE = "pt"
)

func Language() string {

	if language := LocalStorage.Call("getItem", "language").String(); isSupportedLanguage(language) {

		return language
	}

	if language := Navigator.Get("language").String(); language != "" {

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

	LocalStorage.Call("setItem", "language", language)
}

func isSupportedLanguage(language string) bool {

	switch language {

	case ENGLISH, SPANISH, PORTUGUESE:

		return true
	}

	return false
}
