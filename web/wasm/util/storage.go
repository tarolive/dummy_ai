package util

import (
	"syscall/js"
)

var (
	window         = js.Global()
	localStorage   = window.Get("localStorage")
	sessionStorage = window.Get("sessionStorage")
)

func LocalStorage(key string) string {

	return localStorage.Call("getItem", key).String()
}

func SetLocalStorage(key string, value string) {

	localStorage.Call("setItem", key, value)
}

func ClearLocalStorage() {

	localStorage.Call("clear")
}

func SessionStorage(key string) string {

	return sessionStorage.Call("getItem", key).String()
}

func SetSessionStorage(key string, value string) {

	sessionStorage.Call("setItem", key, value)
}

func ClearSessionStorage() {

	sessionStorage.Call("clear")
}
