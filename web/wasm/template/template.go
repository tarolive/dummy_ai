package template

import (
	"syscall/js"
)

var (
	navigator     = createNavigator()
	pageContainer = createPageContainer()
)

func init() {
}

func Navigator() js.Value {

	return navigator
}

func PageContainer() js.Value {

	return pageContainer
}

func createNavigator() js.Value {

	return js.Value{}
}

func createPageContainer() js.Value {

	return js.Value{}
}
