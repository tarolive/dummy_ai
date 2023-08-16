package template

import (
	"syscall/js"
)

var (
	navigation = createNavigation()
	drawer     = createDrawer()
	page       = createPage()
)

func init() {

	js.Global().Get("document").Get("documentElement").Set("lang", Language())
	js.Global().Get("document").Get("body").Call("appendChild", navigation)
	js.Global().Get("document").Get("body").Call("appendChild", drawer)
	js.Global().Get("document").Get("body").Call("appendChild", page)
}

func Navigation() js.Value {

	return navigation
}

func Drawer() js.Value {

	return drawer
}

func Page() js.Value {

	return page
}

func createNavigation() js.Value {

	navigation := js.Global().Get("document").Call("createElement", "div")
	navigation.Get("style").Set("position" /*         */, "fixed")
	navigation.Get("style").Set("top" /*              */, 0)
	navigation.Get("style").Set("right" /*            */, 0)
	navigation.Get("style").Set("left" /*             */, 0)
	navigation.Get("style").Set("height" /*           */, "72px")
	navigation.Get("style").Set("background-color" /* */, "var(--rh-color-surface-darkest)")
	navigation.Get("style").Set("color" /*            */, "var(--rh-color-text-primary-on-dark)")

	return navigation
}

func createDrawer() js.Value {

	drawer := js.Global().Get("document").Call("createElement", "div")
	return drawer
}

func createPage() js.Value {

	page := js.Global().Get("document").Call("createElement", "div")
	return page
}
