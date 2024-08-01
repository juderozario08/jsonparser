package parser

import (
	"testing"
)

func TestValidateSyntax(t *testing.T) {
	tests := []struct {
		tokens   []string
		expected bool
	}{
		{[]string{"{", "(", ")", "}"}, true},
		{[]string{"{", "(", "]", "}"}, false},
	}

	for _, tt := range tests {
		result := ValidateSyntax(tt.tokens)
		if result != tt.expected {
			t.Errorf("ValidateSyntax(%v) = %v; want %v", tt.tokens, result, tt.expected)
		}
	}
}
