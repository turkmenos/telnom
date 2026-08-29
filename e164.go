package telnom

// E164 returns the number in E.164 format.
func (n Number) E164() string {
	return "+" + n.CountryCode + n.National
}
