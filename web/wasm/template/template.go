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
