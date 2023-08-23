package template

import (
	"syscall/js"
)

var (
	navigation           = createNavigation()
	navigationMenuButton = createNavigationMenuButton()
	navigationTitleImage = createNavigationTitleImage()
	navigationTitle      = createNavigationTitle()
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

func NavigationTitleImage() js.Value {

	return navigationTitleImage
}

func NavigationTitle() js.Value {

	return navigationTitle
}

func Drawer() js.Value {

	return drawer
}

func Page() js.Value {

	return page
}

func createNavigation() js.Value {

	navigation := js.Global().Get("document").Call("createElement", "div")
	navigation.Call("appendChild", navigationTitleImage)
	navigation.Call("appendChild", navigationTitle)
	navigation.Get("style").Set("position" /*         */, "fixed")
	navigation.Get("style").Set("top" /*              */, 0)
	navigation.Get("style").Set("right" /*            */, 0)
	navigation.Get("style").Set("left" /*             */, 0)
	navigation.Get("style").Set("height" /*           */, "72px")
	navigation.Get("style").Set("display" /*          */, "flex")
	navigation.Get("style").Set("align-items" /*      */, "center")
	navigation.Get("style").Set("background-color" /* */, "var(--rh-color-surface-darkest)")
	navigation.Get("style").Set("color" /*            */, "var(--rh-color-text-primary-on-dark)")

	if isDesktop {

		navigation.Get("style").Set("padding-left" /* */, "var(--rh-space-xl)")
	}

	if isMobile {

		navigation.Call("insertBefore", navigationMenuButton, navigationTitleImage)
		navigation.Get("style").Set("padding-left" /* */, "var(--rh-space-md)")
	}

	return navigation
}

func createNavigationMenuButton() js.Value {

	img := js.Global().Get("document").Call("createElement", "img")
	img.Set("src", "/menu_open_white.svg")
	img.Set("alt", messages[Menu])
	img.Get("style").Set("width" /*  */, "var(--rh-size-icon-02)")
	img.Get("style").Set("height" /* */, "var(--rh-size-icon-02)")

	div := js.Global().Get("document").Call("createElement", "div")
	div.Call("appendChild", img)
	div.Get("style").Set("width" /*  */, "var(--rh-size-icon-02)")
	div.Get("style").Set("height" /* */, "var(--rh-size-icon-02)")
	div.Get("style").Set("margin" /* */, "var(--rh-space-sm) 0")

	navigationMenuButton := js.Global().Get("document").Call("createElement", "rh-button")
	navigationMenuButton.Set("variant", "link")
	navigationMenuButton.Call("appendChild", div)

	return navigationMenuButton
}

func createNavigationTitleImage() js.Value {

	navigationTitleImage := js.Global().Get("document").Call("createElement", "img")
	navigationTitleImage.Set("src", "/logo.svg")
	navigationTitleImage.Set("alt", "")
	navigationTitleImage.Get("style").Set("width" /*  */, "var(--rh-size-icon-02)")
	navigationTitleImage.Get("style").Set("height" /* */, "var(--rh-size-icon-02)")
	navigationTitleImage.Get("style").Set("margin" /* */, "var(--rh-space-md)")

	return navigationTitleImage
}

func createNavigationTitle() js.Value {

	navigationTitle := js.Global().Get("document").Call("createElement", "h3")
	navigationTitle.Set("innerHTML", messages[DummyAI])

	return navigationTitle
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
	drawer.Get("style").Set("z-index" /*          */, 10)

	if isMobile {

		drawer.Get("style").Set("transform" /* */, "translateX(-100%)")
	}

	return drawer
}

func createPage() js.Value {

	page := js.Global().Get("document").Call("createElement", "div")
	page.Get("style").Set("position" /*         */, "fixed")
	page.Get("style").Set("top" /*              */, "72px")
	page.Get("style").Set("right" /*            */, 0)
	page.Get("style").Set("bottom" /*           */, 0)
	page.Get("style").Set("padding" /*          */, "var(--rh-space-md)")
	page.Get("style").Set("overflow-x" /*       */, "hidden")
	page.Get("style").Set("overflow-y" /*       */, "auto")
	page.Get("style").Set("background-color" /* */, "var(--rh-color-surface-lightest)")
	page.Get("style").Set("color" /*            */, "var(--rh-color-text-primary-on-light)")

	if isDesktop {

		page.Get("style").Set("left" /* */, "290px")
	}

	if isMobile {

		page.Get("style").Set("left" /* */, 0)
	}

	return page
}
