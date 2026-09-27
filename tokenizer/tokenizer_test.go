package tokenizer

import (
	"testing"
)

type TokenizerTest struct {
	name     string
	input    string
	expected []Token
}

func TestTokenizer(t *testing.T) {
	tests := []TokenizerTest{
		{
			name:  "Simple Key-Value Object",
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
			name: "Object with Quoted Primitives (backwards compatibility)",
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
		{
			name:  "Standard JSON with Unquoted Numbers, Booleans, Null",
			input: `{"id": 120391, "age": 20, "active": true, "deleted": false, "data": null, "score": -3.14}`,
			expected: []Token{
				{Type: TokenBraceOpen, Value: "{"},
				{Type: TokenString, Value: "id"},
				{Type: TokenColon, Value: ":"},
				{Type: TokenNumber, Value: "120391"},
				{Type: TokenComma, Value: ","},
				{Type: TokenString, Value: "age"},
				{Type: TokenColon, Value: ":"},
				{Type: TokenNumber, Value: "20"},
				{Type: TokenComma, Value: ","},
				{Type: TokenString, Value: "active"},
				{Type: TokenColon, Value: ":"},
				{Type: TokenBool, Value: "true"},
				{Type: TokenComma, Value: ","},
				{Type: TokenString, Value: "deleted"},
				{Type: TokenColon, Value: ":"},
				{Type: TokenBool, Value: "false"},
				{Type: TokenComma, Value: ","},
				{Type: TokenString, Value: "data"},
				{Type: TokenColon, Value: ":"},
				{Type: TokenNull, Value: "null"},
				{Type: TokenComma, Value: ","},
				{Type: TokenString, Value: "score"},
				{Type: TokenColon, Value: ":"},
				{Type: TokenNumber, Value: "-3.14"},
				{Type: TokenBraceClose, Value: "}"},
			},
		},
		{
			name:  "Escaped String Content",
			input: `{"msg": "Hello \"world\"\nLine 2"}`,
			expected: []Token{
				{Type: TokenBraceOpen, Value: "{"},
				{Type: TokenString, Value: "msg"},
				{Type: TokenColon, Value: ":"},
				{Type: TokenString, Value: "Hello \"world\"\nLine 2"},
				{Type: TokenBraceClose, Value: "}"},
			},
		},
		{
			name:  "Empty Structures",
			input: `{"obj": {}, "arr": []}`,
			expected: []Token{
				{Type: TokenBraceOpen, Value: "{"},
				{Type: TokenString, Value: "obj"},
				{Type: TokenColon, Value: ":"},
				{Type: TokenBraceOpen, Value: "{"},
				{Type: TokenBraceClose, Value: "}"},
				{Type: TokenComma, Value: ","},
				{Type: TokenString, Value: "arr"},
				{Type: TokenColon, Value: ":"},
				{Type: TokenSquareOpen, Value: "["},
				{Type: TokenSquareClose, Value: "]"},
				{Type: TokenBraceClose, Value: "}"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Tokenizer(test.input)
			if !result.IsEqual(test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestUnterminatedString(t *testing.T) {
	// Must not panic on unterminated strings
	tokens := Tokenizer(`{"unclosed": "value`)
	if len(tokens) == 0 {
		t.Errorf("Expected tokens to be extracted without panic")
	}
}
