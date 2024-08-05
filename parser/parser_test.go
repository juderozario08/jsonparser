package parser

import (
	"jsonparser/tokenizer"
	"testing"
)

type TestStruct struct {
	Value    string
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
	tests := []TestStruct{
		{Value: `{names: ["John", "nelson"]}`, Expected: nil},
	}
	for _, test := range tests {
		result, err := Parser(tokenizer.Tokenizer(test.Value))
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
