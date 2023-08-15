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

	return js.Value{}
}

func createDrawer() js.Value {

	return js.Value{}
}

func createPage() js.Value {

	return js.Value{}
}
