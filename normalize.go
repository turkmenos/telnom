package telnom

func Normalize(input string) (string, error) {
	number, err := Parse(input)
	if err != nil {
		return "", err
	}
	return "+" + number.CountryCode + number.National, nil
}
