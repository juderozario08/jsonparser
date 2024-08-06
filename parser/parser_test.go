package parser

import (
	"jsonparser/tokenizer"
	"testing"
)

type TestStruct struct {
	Value    string
	Expected ASTNode
}

type TestArrayStruct struct {
	Value    tokenizer.Tokens
	Expected ASTNode
}

func TestParser(t *testing.T) {
	tests := []TestStruct{
		{Value: "", Expected: nil},
		{Value: "", Expected: nil},
	}
	for _, test := range tests {
		result, err := Parser(tokenizer.Tokenizer(test.Value))
		if err != nil {
			t.Errorf("Test Failed for %v. Want \n%v and got \n%v", test.Value, result, test.Expected)
		}
	}
}

func TestArrayParser(t *testing.T) {
	tests := []TestArrayStruct{
		{Value: []tokenizer.Token{
			{Type: tokenizer.TokenSquareOpen, Value: "["},
			{Type: tokenizer.TokenComma, Value: ","},
			{Type: tokenizer.TokenString, Value: "Jude"},
			{Type: tokenizer.TokenComma, Value: ","},
			{Type: tokenizer.TokenString, Value: "Sara"},
			{Type: tokenizer.TokenComma, Value: ","},
			{Type: tokenizer.TokenNumber, Value: "2"},
			{Type: tokenizer.TokenComma, Value: ","},
			{Type: tokenizer.TokenBool, Value: "true"},
			{Type: tokenizer.TokenComma, Value: ","},
			{Type: tokenizer.TokenNull, Value: "nil"},
			{Type: tokenizer.TokenComma, Value: ","},
			{Type: tokenizer.TokenBool, Value: "false"},
			{Type: tokenizer.TokenSquareClose, Value: "]"},
		}, Expected: nil},
	}
	i := 0
	for _, test := range tests {
		tokens := test.Value
		result, err := ParseAndValidateArray(&tokens, &i)
		if err != nil {
			t.Errorf("Test failed, for %v. Want \n%v and got \n%v", test.Value, result, test.Expected)
		}
	}
}

func TestObjectParser(t *testing.T) {
	tests := []TestStruct{
		{Value: `
			{
				"person": {
					"name": "Jude"
					"age": "20"
				}
			}`, Expected: nil},
	}
	for _, test := range tests {
		result, err := Parser(tokenizer.Tokenizer(test.Value))
		if err != nil {
			t.Errorf("Test failed, for %v. Want \n%v and got \n%v", test.Value, result, test.Expected)
		}
	}
}
