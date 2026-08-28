package telnom

type NumberType string

const (
	TypeUnknown   NumberType = "Unknown"
	TypeMobile    NumberType = "Mobile"
	TypeFixedLine NumberType = "FixedLine"
)

func TypeOf(input string) NumberType {
	n, err := Parse(input)
	if err != nil {
		return TypeUnknown
	}

	if isMobileNational(n.National) {
		return TypeMobile
	}
	if isFixedLineNational(n.National) {
		return TypeFixedLine
	}
	return TypeUnknown
}
