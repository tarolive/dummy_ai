package layout

import (
	"syscall/js"
)

var (
	window = js.Global()
)

var (
	document = window.Get("document")
)

var (
	head = document.Get("head")
	body = document.Get("body")
)

func CreatePage(page js.Value) {

	body.Call("appendChild", createNavigation())
	body.Call("appendChild", createPageContainer(page))
}

func createNavigation() js.Value {

	navigation := document.Call("createElement", "div")

	navigationStyle := navigation.Get("style")
	navigationStyle.Set("position" /*         */, "fixed")
	navigationStyle.Set("top" /*              */, "0")
	navigationStyle.Set("right" /*            */, "0")
	navigationStyle.Set("left" /*             */, "0")
	navigationStyle.Set("height" /*           */, "72px")
	navigationStyle.Set("background-color" /* */, "var(--rh-color-surface-darkest)")
	navigationStyle.Set("color" /*            */, "var(--rh-color-text-primary-on-dark)")

	return navigation
}

func createPageContainer(page js.Value) js.Value {

	pageContainer := document.Call("createElement", "div")
	pageContainer.Call("appendChild", page)

	pageContainerStyle := pageContainer.Get("style")
	pageContainerStyle.Set("position" /*         */, "fixed")
	pageContainerStyle.Set("inset" /*            */, "72px 0 0 0")
	pageContainerStyle.Set("padding" /*          */, "10px")
	pageContainerStyle.Set("overflow-x" /*       */, "hidden")
	pageContainerStyle.Set("overflow-y" /*       */, "auto")
	pageContainerStyle.Set("background-color" /* */, "var(--rh-color-surface-lightest)")
	pageContainerStyle.Set("color" /*            */, "var(--rh-color-text-primary-on-light)")

	return pageContainer
}
