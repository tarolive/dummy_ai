package template

import (
	"syscall/js"
)

var (
	isDesktop = matchesBreakpoint()
	isMobile  = !isDesktop
)

func IsDesktop() bool {

	return isDesktop
}

func IsMobile() bool {

	return isMobile
}

func matchesBreakpoint() bool {

	breakpoint := js.Global().Call("getComputedStyle", js.Global().Get("document").Get("documentElement")).Call("getPropertyValue", "--pf-v5-global--breakpoint--xl").String()
	mediaQuery := "(min-width: " + breakpoint + ")"

	return js.Global().Call("matchMedia", mediaQuery).Get("matches").Bool()
}
