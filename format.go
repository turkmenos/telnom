package telnom

func Format(input string) (string, error) {
	n, err := Parse(input)
	if err != nil {
		return "", err
	}

	return "+" + n.CountryCode + " " +
		n.National[:2] + " " +
		n.National[2:4] + " " +
		n.National[4:6] + " " +
		n.National[6:8], nil
}
