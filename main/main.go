package main

import (
	"fmt"
	"jsonparser/parser"
)

type TokenizerTest struct {
	input    string
	expected []parser.Token
}

func main() {
	test := TokenizerTest{
		input: `{"id":"120391", "name": "Some Name", "age": "20", "something": [], "boolean": "true", "nullValue": "null"}`,
		expected: []parser.Token{
			{Type: parser.TokenBraceOpen, Value: "{"},
			{Type: parser.TokenString, Value: "id"},
			{Type: parser.TokenColon, Value: ":"},
			{Type: parser.TokenNumber, Value: "120391"},
			{Type: parser.TokenComma, Value: ","},
			{Type: parser.TokenString, Value: "name"},
			{Type: parser.TokenColon, Value: ":"},
			{Type: parser.TokenString, Value: "Some Name"},
			{Type: parser.TokenComma, Value: ","},
			{Type: parser.TokenString, Value: "age"},
			{Type: parser.TokenColon, Value: ":"},
			{Type: parser.TokenNumber, Value: "20"},
			{Type: parser.TokenComma, Value: ","},
			{Type: parser.TokenString, Value: "something"},
			{Type: parser.TokenColon, Value: ":"},
			{Type: parser.TokenSquareOpen, Value: "["},
			{Type: parser.TokenSquareClose, Value: "]"},
			{Type: parser.TokenComma, Value: ","},
			{Type: parser.TokenString, Value: "boolean"},
			{Type: parser.TokenColon, Value: ":"},
			{Type: parser.TokenBool, Value: "true"},
			{Type: parser.TokenComma, Value: ","},
			{Type: parser.TokenString, Value: "nullValue"},
			{Type: parser.TokenColon, Value: ":"},
			{Type: parser.TokenNull, Value: "null"},
			{Type: parser.TokenBraceClose, Value: "}"},
		},
	}
	result := parser.Tokenizer(test.input)
	fmt.Println(result.IsEqual(test.expected))
}
