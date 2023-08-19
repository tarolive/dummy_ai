package template

import (
	"syscall/js"
)

var (
	language = Language()
	messages = map[string]map[string]string{
		English: {
			"navigationMenuButton.alt": "Menu",
		},
		Spanish: {
			"navigationMenuButton.alt": "Menu",
		},
		Portuguese: {
			"navigationMenuButton.alt": "Menu",
		},
	}[language]
)

var (
	navigation           = createNavigation()
	navigationMenuButton = createNavigationMenuButton()
	navigationTitleIcon  = createNavigationTitleIcon()
	navigationTitleText  = createNavigationTitleText()
	drawer               = createDrawer()
	page                 = createPage()
)

func init() {

	js.Global().Get("document").Get("documentElement").Set("lang", language)
	js.Global().Get("document").Get("body").Call("appendChild", navigation)
	js.Global().Get("document").Get("body").Call("appendChild", drawer)
	js.Global().Get("document").Get("body").Call("appendChild", page)
}

func Navigation() js.Value {

	return navigation
}

func NavigationMenuButton() js.Value {

	return navigationMenuButton
}

func NavigationTitleIcon() js.Value {

	return navigationTitleIcon
}

func NavigationTitleText() js.Value {

	return navigationTitleText
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
	navigation.Get("style").Set("padding" /*          */, "0 var(--rh-space-xl)")
	navigation.Get("style").Set("display" /*          */, "flex")
	navigation.Get("style").Set("align-items" /*      */, "center")
	navigation.Get("style").Set("background-color" /* */, "var(--rh-color-surface-darkest)")
	navigation.Get("style").Set("color" /*            */, "var(--rh-color-text-primary-on-dark)")
	navigation.Call("appendChild", navigationMenuButton)
	navigation.Call("appendChild", navigationTitleIcon)
	navigation.Call("appendChild", navigationTitleText)

	return navigation
}

func createNavigationMenuButton() js.Value {

	img := js.Global().Get("document").Call("createElement", "img")
	img.Get("style").Set("width" /*  */, "var(--rh-size-icon-02)")
	img.Get("style").Set("height" /* */, "var(--rh-size-icon-02)")
	img.Set("src", "/menu_open_white.svg")
	img.Set("alt", messages["navigationMenuButton.alt"])

	div := js.Global().Get("document").Call("createElement", "div")
	div.Get("style").Set("width" /*  */, "var(--rh-size-icon-02)")
	div.Get("style").Set("height" /* */, "var(--rh-size-icon-02)")
	div.Call("appendChild", img)

	navigationMenuButton := js.Global().Get("document").Call("createElement", "rh-button")
	navigationMenuButton.Set("variant", "link")
	navigationMenuButton.Call("appendChild", div)

	return navigationMenuButton
}

func createNavigationTitleIcon() js.Value {

	navigationTitleIcon := js.Global().Get("document").Call("createElement", "img")
	navigationTitleIcon.Get("style").Set("width" /*  */, "var(--rh-size-icon-03)")
	navigationTitleIcon.Get("style").Set("height" /* */, "var(--rh-size-icon-03)")
	navigationTitleIcon.Get("style").Set("margin" /* */, "var(--rh-space-md)")
	navigationTitleIcon.Set("src", "/logo.svg")
	navigationTitleIcon.Set("alt", "")

	return navigationTitleIcon
}

func createNavigationTitleText() js.Value {

	navigationTitleText := js.Global().Get("document").Call("createElement", "h2")
	navigationTitleText.Set("innerHTML", "DummyAI")

	return navigationTitleText
}

func createDrawer() js.Value {

	drawer := js.Global().Get("document").Call("createElement", "div")
	drawer.Get("style").Set("position" /*         */, "fixed")
	drawer.Get("style").Set("top" /*              */, "72px")
	drawer.Get("style").Set("bottom" /*           */, 0)
	drawer.Get("style").Set("left" /*             */, 0)
	drawer.Get("style").Set("width" /*            */, "100%")
	drawer.Get("style").Set("max-width" /*        */, "290px")
	drawer.Get("style").Set("overflow-x" /*       */, "hidden")
	drawer.Get("style").Set("overflow-y" /*       */, "auto")
	drawer.Get("style").Set("background-color" /* */, "#212427")
	drawer.Get("style").Set("color" /*            */, "var(--rh-color-text-primary-on-dark)")

	return drawer
}

func createPage() js.Value {

	page := js.Global().Get("document").Call("createElement", "div")
	page.Get("style").Set("position" /*         */, "fixed")
	page.Get("style").Set("top" /*              */, "72px")
	page.Get("style").Set("right" /*            */, 0)
	page.Get("style").Set("bottom" /*           */, 0)
	page.Get("style").Set("left" /*             */, "290px")
	page.Get("style").Set("padding" /*          */, "var(--rh-space-md)")
	page.Get("style").Set("overflow-x" /*       */, "hidden")
	page.Get("style").Set("overflow-y" /*       */, "auto")
	page.Get("style").Set("background-color" /* */, "var(--rh-color-surface-lightest)")
	page.Get("style").Set("color" /*            */, "var(--rh-color-text-primary-on-light)")

	return page
}
