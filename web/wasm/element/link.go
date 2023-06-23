package element

import (
	"syscall/js"
)

import (
	"dummy_ai/web/wasm/dom"
)

func NewLink(rel string, href string) js.Value {

	link := dom.Document.Call("createElement", "link")
	link.Set("rel", rel)
	link.Set("href", href)

	return link
}

func NewLinkWithSizes(rel string, href string, sizes string) js.Value {

	link := NewLink(rel, href)
	link.Set("sizes", sizes)

	return link
}

func NewLinkWithType(rel string, href string, linkType string) js.Value {

	link := NewLink(rel, href)
	link.Set("type", linkType)

	return link
}
