package telnom

// IsFixedLine reports whether input is a valid Turkmen fixed-line number.
func IsFixedLine(input string) bool {
	n, err := Parse(input)
	if err != nil {
		return false
	}
	return isFixedLineNational(n.National)
}

func isFixedLineNational(national string) bool {
	switch national[0] {
	case '1', '2', '3', '4', '5':
		return true
	default:
		return false
	}
}
