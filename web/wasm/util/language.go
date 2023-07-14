package util

var (
	languages = map[string]string{
		"en": "English",
		"es": "Español",
		"pt": "Português",
	}
	language = ""
)

func Language() string {

	if language != "" {

		return language
	}

	if language = LocalStorage("language"); languages[language] != "" {

		return language
	}

	return "en"
}
