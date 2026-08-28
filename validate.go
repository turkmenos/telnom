package telnom

func Validate(input string) bool {
	_, err := Parse(input)
	return err == nil
}
