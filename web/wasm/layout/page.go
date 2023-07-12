package layout

import (
	"syscall/js"
)

var (
	window   = js.Global()
	document = window.Get("document")
)

func CreatePage() js.Value {

	page := createPage()

	body := document.Get("body")
	body.Call("appendChild", createNavigation())
	body.Call("appendChild", createPageContainer(page))

	return page
}

func createPage() js.Value {

	page := document.Call("createElement", "div")
	return page
}

func createNavigation() js.Value {

	navigation := document.Call("createElement", "div")
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

	navigationImage := document.Call("createElement", "img")
	navigationImage.Set("src", "/logo.svg")
	navigationImage.Set("alt", "")

	navigationImageStyle := navigationImage.Get("style")
	navigationImageStyle.Set("width" /*        */, "var(--rh-space-2xl)")
	navigationImageStyle.Set("height" /*       */, "var(--rh-space-2xl)")
	navigationImageStyle.Set("margin-right" /* */, "var(--rh-space-xl)")

	return navigationImage
}

func createNavigationText() js.Value {

	navigationText := document.Call("createElement", "h2")
	navigationText.Set("innerHTML", "DummyAI")

	return navigationText
}

func createPageContainer(page js.Value) js.Value {

	pageContainer := document.Call("createElement", "div")
	pageContainer.Call("appendChild", page)

	pageContainerStyle := pageContainer.Get("style")
	pageContainerStyle.Set("position" /*         */, "fixed")
	pageContainerStyle.Set("top" /*              */, "72px")
	pageContainerStyle.Set("right" /*            */, 0)
	pageContainerStyle.Set("bottom" /*           */, 0)
	pageContainerStyle.Set("left" /*             */, 0)
	pageContainerStyle.Set("padding" /*          */, "var(--rh-space-md)")
	pageContainerStyle.Set("overflow-x" /*       */, "hidden")
	pageContainerStyle.Set("overflow-y" /*       */, "auto")
	pageContainerStyle.Set("background-color" /* */, "var(--rh-color-surface-lightest)")
	pageContainerStyle.Set("color" /*            */, "var(--rh-color-text-primary-on-light)")

	return pageContainer
}
