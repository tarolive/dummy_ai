package util

var (
	language = ""
)

func Language() string {

	if language != "" {

		return language
	}

	if language = LocalStorage("language"); language != "" {

		return language
	}

	return "en"
}
