package template

var (
	isDesktop = true
	isMobile  = !isDesktop
)

func IsDesktop() bool {

	return isDesktop
}

func IsMobile() bool {

	return isMobile
}
