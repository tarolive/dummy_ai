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

	if breakpoint == "" {

		breakpoint = "1200px"
	}

	return js.Global().Call("matchMedia", "(min-width: "+breakpoint+")").Get("matches").Bool()
}
