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

	createMetaCharset()
	createFavicons()
	createManifest()
	createTitle()
	createPage(page)
}

func createMetaCharset() {

	metaCharset := document.Call("createElement", "meta")
	metaCharset.Set("charset", "UTF-8")
	head.Call("appendChild", metaCharset)
}

func createFavicons() {

	faviconICO := document.Call("createElement", "link")
	faviconICO.Set("rel", "icon")
	faviconICO.Set("href", "/favicon.ico")
	faviconICO.Set("sizes", "any")
	head.Call("appendChild", faviconICO)

	faviconSVG := document.Call("createElement", "link")
	faviconSVG.Set("rel", "icon")
	faviconSVG.Set("href", "/favicon.svg")
	faviconSVG.Set("type", "image/svg+xml")
	head.Call("appendChild", faviconSVG)

	appleTouchIcon := document.Call("createElement", "link")
	appleTouchIcon.Set("rel", "apple-touch-icon")
	appleTouchIcon.Set("href", "/apple_touch_icon.png")
	head.Call("appendChild", appleTouchIcon)
}

func createManifest() {

	manifest := document.Call("createElement", "link")
	manifest.Set("rel", "manifest")
	manifest.Set("href", "/manifest.json")
	head.Call("appendChild", manifest)
}

func createTitle() {

	title := document.Call("createElement", "title")
	title.Set("innerHTML", "DummyAI")
	head.Call("appendChild", title)
}

func createPage(page js.Value) {

	body.Call("appendChild", page)
}
