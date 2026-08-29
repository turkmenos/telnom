package telnom

// Region identifies the geographic origin of a fixed-line number.
type Region string

const (
	RegionUnknown  Region = ""
	RegionAshgabat Region = "Ashgabat"
	RegionAhal     Region = "Ahal"
	RegionBalkan   Region = "Balkan"
	RegionDashoguz Region = "Dashoguz"
	RegionLebap    Region = "Lebap"
	RegionMary     Region = "Mary"
)

// RegionOf returns the region associated with a Turkmen fixed-line number.
// It returns RegionUnknown when input is invalid, non-geographic, or its
// prefix is not recognized.
func RegionOf(input string) Region {
	n, err := Parse(input)
	if err != nil {
		return RegionUnknown
	}

	if n.National[:2] == "12" {
		return RegionAshgabat
	}

	switch n.National[0] {
	case '1':
		return RegionAhal
	case '2':
		return RegionBalkan
	case '3':
		return RegionDashoguz
	case '4':
		return RegionLebap
	case '5':
		return RegionMary
	default:
		return RegionUnknown
	}
}
