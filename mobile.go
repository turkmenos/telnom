package telnom

func IsMobile(input string) bool {
	n, err := Parse(input)
	if err != nil {
		return false
	}

	return isMobileNational(n.National)
}

func isMobileNational(national string) bool {
	switch national[:2] {
	case "71", "72":
		return true
	default:
		return false
	}
}
