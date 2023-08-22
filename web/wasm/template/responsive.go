package template

import (
	"syscall/js"
)

var (
	media      = js.Global().Call("getComputedStyle", js.Global().Get("document").Get("documentElement")).Call("getPropertyValue", "--rh-media-sm")
	matchMedia = js.Global().Call("matchMedia", media)
)

var (
	isDesktop = matchMedia.Get("matches").Bool()
	isMobile  = !isDesktop
)

func IsDesktop() bool {

	return isDesktop
}

func IsMobile() bool {

	return isMobile
}
