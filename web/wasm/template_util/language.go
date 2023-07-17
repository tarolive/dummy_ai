package template_util

const (
	ENGLISH    = "en"
	SPANISH    = "es"
	PORTUGUESE = "pt"
)

func Language() string {

	return ENGLISH
}

func isSupported(language string) bool {

	switch language {

	case ENGLISH:
	case SPANISH:
	case PORTUGUESE:

		return true
	}

	return false
}
