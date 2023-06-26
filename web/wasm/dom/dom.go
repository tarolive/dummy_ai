package dom

var (
	Window = Global()
)

var (
	Document        = Window.GetElement("document")
	DocumentElement = Window.GetElement("documentElement")
)

var (
	Head = Document.GetElement("head")
	Body = Document.GetElement("body")
)
