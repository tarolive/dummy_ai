package template

import (
	"syscall/js"
)

var (
	navigator = createNavigator()
	page      = createPage()
)

func init() {
}

func Navigator() js.Value {

	return navigator
}

func Page() js.Value {

	return page
}

func createNavigator() js.Value {

	return js.Value{}
}

func createPage() js.Value {

	return js.Value{}
}
