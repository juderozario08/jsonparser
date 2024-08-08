package tokenizer

import (
	"testing"
)

type TokenizerTest struct {
	input    string
	expected []Token
}

func TestTokenizer(t *testing.T) {
	tests := []TokenizerTest{
		{
			input: `{"Key":"Value","Key2":"Value2"}`,
			expected: []Token{
				{Type: TokenBraceOpen, Value: "{"},
				{Type: TokenString, Value: "Key"},
				{Type: TokenColon, Value: ":"},
				{Type: TokenString, Value: "Value"},
				{Type: TokenComma, Value: ","},
				{Type: TokenString, Value: "Key2"},
				{Type: TokenColon, Value: ":"},
				{Type: TokenString, Value: "Value2"},
				{Type: TokenBraceClose, Value: "}"},
			},
		},
		{
			input: `
					{
						"id":"120391",
						"name": "Some Name",
						"age": "20",
						"something": [],
						"boolean": "true",
						"nullValue": "null"
					}
				`,
			expected: []Token{
				{Type: TokenBraceOpen, Value: "{"},
				{Type: TokenString, Value: "id"},
				{Type: TokenColon, Value: ":"},
				{Type: TokenNumber, Value: "120391"},
				{Type: TokenComma, Value: ","},
				{Type: TokenString, Value: "name"},
				{Type: TokenColon, Value: ":"},
				{Type: TokenString, Value: "Some Name"},
				{Type: TokenComma, Value: ","},
				{Type: TokenString, Value: "age"},
				{Type: TokenColon, Value: ":"},
				{Type: TokenNumber, Value: "20"},
				{Type: TokenComma, Value: ","},
				{Type: TokenString, Value: "something"},
				{Type: TokenColon, Value: ":"},
				{Type: TokenSquareOpen, Value: "["},
				{Type: TokenSquareClose, Value: "]"},
				{Type: TokenComma, Value: ","},
				{Type: TokenString, Value: "boolean"},
				{Type: TokenColon, Value: ":"},
				{Type: TokenBool, Value: "true"},
				{Type: TokenComma, Value: ","},
				{Type: TokenString, Value: "nullValue"},
				{Type: TokenColon, Value: ":"},
				{Type: TokenNull, Value: "null"},
				{Type: TokenBraceClose, Value: "}"},
			},
		},
	}

	for _, test := range tests {
		result := Tokenizer(test.input)
		if !result.IsEqual(test.expected) {
			t.Errorf("Expected %v, got %v", test.expected, result)
		}
	}
}

func (tokens *Tokens) IsEqual(otherTokens Tokens) bool {
	if len(*tokens) != len(otherTokens) {
		return false
	}
	for i, token := range *tokens {
		if token != otherTokens[i] {
			return false
		}
	}
	return true
}
