package telnom_test

import (
	"testing"

	"github.com/turkmenos/telnom"
)

func TestValidate(t *testing.T) {

	tests := []struct {
		input string
		want  bool
	}{
		{"+99371123456", true},
		{"99371123456", true},
		{"+993 71 12 34 56", true},
		{"", false},
		{"+905551234567", false},
		{"+99371abc456", false},
	}
	for _, tt := range tests {
		got := telnom.Validate(tt.input)
		if got != tt.want {
			t.Errorf("Validate(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}

}

func TestFormat(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"+99371123456", "+993 71 12 34 56"},
		{"99371123456", "+993 71 12 34 56"},
		{"+993 71 12 34 56", "+993 71 12 34 56"},
	}

	for _, tt := range tests {
		got, err := telnom.Format(tt.input)
		if err != nil {
			t.Fatalf("Format(%q) returned error: %v", tt.input, err)
		}

		if got != tt.want {
			t.Errorf("Format(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestOperatorOf(t *testing.T) {
	tests := []struct {
		input string
		want  telnom.Operator
	}{
		{"+99371123456", telnom.OperatorTMCell},
		{"+99372123456", telnom.OperatorTMCell},
		{"invalid", telnom.OperatorUnknown},
	}

	for _, tt := range tests {
		got := telnom.OperatorOf(tt.input)

		if got != tt.want {
			t.Errorf("OperatorOf(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestParseInvalid(t *testing.T) {
	tests := []string{
		"",
		"+905551234567",
		"+99371abc456",
		"+99371123",
	}

	for _, input := range tests {
		_, err := telnom.Parse(input)
		if err == nil {
			t.Errorf("Parse(%q) expected error", input)
		}
	}
}

func TestIsMobile(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"+99371123456", true},
		{"+99372123456", true},
		{"+99312123456", false},
		{"invalid", false},
	}

	for _, tt := range tests {
		got := telnom.IsMobile(tt.input)
		if got != tt.want {
			t.Errorf("IsMobile(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsFixedLine(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"+99312123456", true},
		{"+99324312345", true},
		{"+99332212345", true},
		{"+99342212345", true},
		{"+99352212345", true},
		{"993 12 12 34 56", true},
		{"+99371123456", false},
		{"+99372123456", false},
		{"+99381234567", false},
		{"invalid", false},
	}

	for _, tt := range tests {
		got := telnom.IsFixedLine(tt.input)
		if got != tt.want {
			t.Errorf("IsFixedLine(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}
