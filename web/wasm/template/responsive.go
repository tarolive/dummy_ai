package template

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

	return true
}
