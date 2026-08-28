package telnom

import (
	"errors"
	"strings"
)

const countryCode = "993"

type Number struct {
	CountryCode string
	National    string
}

func Parse(input string) (Number, error) {
	s := strings.TrimSpace(input)

	s = strings.NewReplacer(
		" ", "",
		"-", "",
		"(", "",
		")", "",
	).Replace(s)

	switch {
	case strings.HasPrefix(s, "+993"):
		s = strings.TrimPrefix(s, "+993")
	case strings.HasPrefix(s, "993"):
		s = strings.TrimPrefix(s, "993")
	default:
		return Number{}, errors.New("invalid country code")
	}

	if len(s) != 8 {
		return Number{}, errors.New("invalid number length")
	}

	for _, r := range s {
		if r < '0' || r > '9' {
			return Number{}, errors.New("invalid character")
		}
	}

	return Number{
		CountryCode: countryCode,
		National:    s,
	}, nil
}
