package template

import (
	"syscall/js"
)

var (
	mastheadToggle  = createMastheadToggle()
	mastheadBrand   = createMastheadBrand()
	mastheadMain    = createMastheadMain()
	mastheadContent = createMastheadContent()
	masthead        = createMasthead()
	sidebarBody     = createSidebarBody()
	sidebar         = createSidebar()
	pageMainSection = createPageMainSection()
	pageMain        = createPageMain()
	page            = createPage()
)

func init() {

	js.Global().Get("document").Get("documentElement").Set("lang", language)
	js.Global().Get("document").Get("body").Call("appendChild", page)
}

func MastheadToggle() js.Value {

	return mastheadToggle
}

func MastheadBrand() js.Value {

	return mastheadBrand
}

func MastheadMain() js.Value {

	return mastheadMain
}

func MastheadContent() js.Value {

	return mastheadContent
}

func Masthead() js.Value {

	return masthead
}

func SidebarBody() js.Value {

	return sidebarBody
}

func Sidebar() js.Value {

	return sidebar
}

func PageMainSection() js.Value {

	return pageMainSection
}

func PageMain() js.Value {

	return pageMain
}

func Page() js.Value {

	return page
}

func createMastheadToggle() js.Value {

	i := js.Global().Get("document").Call("createElement", "i")
	i.Get("classList").Call("add", "fas")
	i.Get("classList").Call("add", "fa-bars")
	i.Set("ariaHidden", true)

	button := js.Global().Get("document").Call("createElement", "button")
	button.Get("classList").Call("add", "pf-v5-c-button")
	button.Get("classList").Call("add", "pf-m-plain")
	button.Set("type", "button")
	button.Set("ariaLabel", messages[Menu])
	button.Call("appendChild", i)

	mastheadToggle := js.Global().Get("document").Call("createElement", "span")
	mastheadToggle.Get("classList").Call("add", "pf-v5-c-masthead__toggle")
	mastheadToggle.Call("appendChild", button)

	return mastheadToggle
}

func createMastheadBrand() js.Value {

	img := js.Global().Get("document").Call("createElement", "img")
	img.Set("src", "/logo.svg")
	img.Set("alt", "")
	img.Set("width", 24)
	img.Set("height", 24)

	mastheadBrand := js.Global().Get("document").Call("createElement", "a")
	mastheadBrand.Get("classList").Call("add", "pf-v5-c-masthead__brand")
	mastheadBrand.Set("href", "/")
	mastheadBrand.Call("appendChild", img)

	return mastheadBrand
}

func createMastheadMain() js.Value {

	mastheadMain := js.Global().Get("document").Call("createElement", "div")
	mastheadMain.Get("classList").Call("add", "pf-v5-c-masthead__main")
	mastheadMain.Call("appendChild", mastheadBrand)

	return mastheadMain
}

func createMastheadContent() js.Value {

	title := js.Global().Get("document").Call("createElement", "h1")
	title.Get("classList").Call("add", "pf-v5-c-title")
	title.Get("classList").Call("add", "pf-m-xl")
	title.Set("innerHTML", messages[DummyAI])

	mastheadContent := js.Global().Get("document").Call("createElement", "div")
	mastheadContent.Get("classList").Call("add", "pf-v5-c-masthead__content")
	mastheadContent.Get("classList").Call("add", "pf-v5-u-ml-0")
	mastheadContent.Call("appendChild", title)

	return mastheadContent
}

func createMasthead() js.Value {

	masthead := js.Global().Get("document").Call("createElement", "header")
	masthead.Get("classList").Call("add", "pf-v5-c-masthead")
	masthead.Get("classList").Call("add", "pf-m-inset-lg")
	masthead.Get("classList").Call("add", "pf-m-display-inline")
	masthead.Call("appendChild", mastheadToggle)
	masthead.Call("appendChild", mastheadMain)
	masthead.Call("appendChild", mastheadContent)

	return masthead
}

func createSidebarBody() js.Value {

	sidebarBody := js.Global().Get("document").Call("createElement", "div")
	sidebarBody.Get("classList").Call("add", "pf-v5-c-page__sidebar-body")

	return sidebarBody
}

func createSidebar() js.Value {

	sidebar := js.Global().Get("document").Call("createElement", "div")
	sidebar.Get("classList").Call("add", "pf-v5-c-page__sidebar")
	sidebar.Call("appendChild", sidebarBody)

	return sidebar
}

func createPageMainSection() js.Value {

	pageMainSection := js.Global().Get("document").Call("createElement", "section")
	pageMainSection.Get("classList").Call("add", "pf-v5-c-page__main-section")
	pageMainSection.Get("classList").Call("add", "pf-v5-u-p-md")

	return pageMainSection
}

func createPageMain() js.Value {

	pageMain := js.Global().Get("document").Call("createElement", "main")
	pageMain.Get("classList").Call("add", "pf-v5-c-page__main")
	pageMain.Set("tabIndex", -1)
	pageMain.Call("appendChild", pageMainSection)

	return pageMain
}

func createPage() js.Value {

	page := js.Global().Get("document").Call("createElement", "div")
	page.Get("classList").Call("add", "pf-v5-c-page")
	page.Call("appendChild", masthead)
	page.Call("appendChild", sidebar)
	page.Call("appendChild", pageMain)

	return page
}
