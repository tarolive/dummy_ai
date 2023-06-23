package element

import (
	"syscall/js"
)

import (
	"dummy_ai/web/wasm/dom"
)

func NewMeta(name string, content string) js.Value {

	meta := dom.Document.Call("createElement", "meta")
	meta.Set("name", name)
	meta.Set("content", content)

	return meta
}

func NewMetaCharset(charset string) js.Value {

	meta := dom.Document.Call("createElement", "meta")
	meta.Set("charset", charset)

	return meta
}
