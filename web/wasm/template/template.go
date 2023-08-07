package template

import (
	"syscall/js"
)

var (
	Window = js.Global()
)

var (
	Document        = Window.Get("document")
	DocumentElement = Document.Get("documentElement")
	Head            = Document.Get("head")
	Body            = Document.Get("body")
)

var (
	SessionStorage = Window.Get("sessionStorage")
	LocalStorage   = Window.Get("localStorage")
)

var (
	Navigator = Window.Get("navigator")
)

var (
	Page = createPage()
)

const (
	ENGLISH    = "en"
	SPANISH    = "es"
	PORTUGUESE = "pt"
)

var (
	Language = ENGLISH
)

func init() {

	DocumentElement.Set("lang", Language)

	Body.Call("appendChild", createNavigation())
	Body.Call("appendChild", Page)
}

func createPage() js.Value {

	page := Document.Call("createElement", "div")

	pageStyle := page.Get("style")
	pageStyle.Set("position" /*         */, "fixed")
	pageStyle.Set("top" /*              */, "72px")
	pageStyle.Set("right" /*            */, 0)
	pageStyle.Set("bottom" /*           */, 0)
	pageStyle.Set("left" /*             */, 0)
	pageStyle.Set("padding" /*          */, "var(--rh-space-md)")
	pageStyle.Set("overflow-x" /*       */, "hidden")
	pageStyle.Set("overflow-y" /*       */, "auto")
	pageStyle.Set("background-color" /* */, "var(--rh-color-surface-lightest)")
	pageStyle.Set("color" /*            */, "var(--rh-color-text-primary-on-light)")

	return page
}

func createNavigation() js.Value {

	navigation := Document.Call("createElement", "div")
	navigation.Call("appendChild", createNavigationImage())
	navigation.Call("appendChild", createNavigationText())

	navigationStyle := navigation.Get("style")
	navigationStyle.Set("position" /*         */, "fixed")
	navigationStyle.Set("top" /*              */, 0)
	navigationStyle.Set("right" /*            */, 0)
	navigationStyle.Set("left" /*             */, 0)
	navigationStyle.Set("height" /*           */, "72px")
	navigationStyle.Set("padding" /*          */, "0 var(--rh-space-xl)")
	navigationStyle.Set("display" /*          */, "flex")
	navigationStyle.Set("align-items" /*      */, "center")
	navigationStyle.Set("background-color" /* */, "var(--rh-color-surface-darkest)")
	navigationStyle.Set("color" /*            */, "var(--rh-color-text-primary-on-dark)")

	return navigation
}

func createNavigationImage() js.Value {

	navigationImage := Document.Call("createElement", "img")
	navigationImage.Set("src", "/logo.svg")
	navigationImage.Set("alt", "")

	navigationImageStyle := navigationImage.Get("style")
	navigationImageStyle.Set("width" /*  */, "var(--rh-space-2xl)")
	navigationImageStyle.Set("height" /* */, "var(--rh-space-2xl)")

	return navigationImage
}

func createNavigationText() js.Value {

	navigationText := Document.Call("createElement", "h2")
	navigationText.Set("innerHTML", "DummyAI")

	navigationTextStyle := navigationText.Get("style")
	navigationTextStyle.Set("margin" /* */, "0 var(--rh-space-xl)")

	return navigationText
}
