package template

import (
	"syscall/js"
)

var (
	window = js.Global()
)

var (
	document        = window.Get("document")
	documentElement = document.Get("documentElement")
	head            = document.Get("head")
	body            = document.Get("body")
)

var (
	sessionStorage = window.Get("sessionStorage")
	localStorage   = window.Get("localStorage")
)

func Window() js.Value {

	return window
}

func Document() js.Value {

	return document
}

func DocumentElement() js.Value {

	return documentElement
}

func Head() js.Value {

	return head
}

func Body() js.Value {

	return body
}

func SessionStorage() js.Value {

	return sessionStorage
}

func LocalStorage() js.Value {

	return localStorage
}
