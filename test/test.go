package test

import (
	parser "jsonparser/main"
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	input := "package main import fmt"
	expected := []string{"package", "main", "import", "fmt"}
	result := parser.Tokenize(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Tokenize(%q) = %v; want %v", input, result, expected)
	}
}

func TestExtractKeywords(t *testing.T) {
	tokens := []string{"package", "main", "import", "fmt", "func", "main"}
	expected := []string{"package", "import", "func"}
	result := parser.ExtractKeywords(tokens)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ExtractKeywords(%v) = %v; want %v", tokens, result, expected)
	}
}

func TestValidateSyntax(t *testing.T) {
	tests := []struct {
		tokens   []string
		expected bool
	}{
		{[]string{"package", "main"}, true},
		{[]string{"import", "fmt"}, false},
		{[]string{"func", "main"}, false},
		{[]string{"package", "main", "func", "main"}, true},
	}

	for _, tt := range tests {
		result := parser.ValidateSyntax(tt.tokens)
		if result != tt.expected {
			t.Errorf("ValidateSyntax(%v) = %v; want %v", tt.tokens, result, tt.expected)
		}
	}
}
