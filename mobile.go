package telnom

func IsMobile(input string) bool {
	n, err := Parse(input)
	if err != nil {
		return false
	}

	prefix := n.National[:2]

	switch prefix {
	case "71", "72":
		return true
	default:
		return false
	}
}
