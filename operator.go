package telnom

type Operator string

const (
	OperatorUnknown Operator = ""
	OperatorTMCell  Operator = "TM Cell"
)

func OperatorOf(input string) Operator {
	n, err := Parse(input)
	if err != nil {
		return OperatorUnknown
	}

	switch n.National[:2] {
	case "71", "72":
		return OperatorTMCell
	default:
		return OperatorUnknown
	}
}
