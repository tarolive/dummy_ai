package template

const (
	DummyAI = iota
	Menu
)

var (
	allMessages = map[string]map[int]string{
		English: {
			DummyAI: "DummyAI",
			Menu:    "Menu",
		},
		Spanish: {
			DummyAI: "DummyAI",
			Menu:    "Menú",
		},
		Portuguese: {
			DummyAI: "DummyAI",
			Menu:    "Menu",
		},
	}
)

var (
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
