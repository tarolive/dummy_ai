package template

const (
	App = iota
	Error404Title
	Error404Text
	ButtonMenuAriaLabel
)

var (
	allMessages = map[string]map[int]string{
		English: {
			App:                 "DummyAI",
			Error404Title:       "",
			Error404Text:        "",
			ButtonMenuAriaLabel: "Menu",
		},
		Spanish: {
			App:                 "DummyAI",
			Error404Title:       "",
			Error404Text:        "",
			ButtonMenuAriaLabel: "Menú",
		},
		Portuguese: {
			App:                 "DummyAI",
			Error404Title:       "",
			Error404Text:        "",
			ButtonMenuAriaLabel: "Menu",
		},
	}
	messages = allMessages[language]
)

func AllMessages() map[string]map[int]string {

	return allMessages
}

func Messages() map[int]string {

	return messages
}

func Message(key int) string {

	return messages[key]
}
