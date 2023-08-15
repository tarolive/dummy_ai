package template

import (
	"syscall/js"
)

var (
	navigation = createNavigation()
	page       = createPage()
)

func init() {
}

func Navigation() js.Value {

	return navigation
}

func Page() js.Value {

	return page
}

func createNavigation() js.Value {

	return js.Value{}
}

func createPage() js.Value {

	return js.Value{}
}
