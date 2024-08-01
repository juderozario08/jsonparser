package main

import (
	"fmt"
	"jsonparser/tokenizer"
)

type TokenizerTest struct {
	input    string
	expected []tokenizer.Token
}

func main() {
	test := TokenizerTest{
		input: `{"id":"120391", "name": "Some Name", "age": "20", "something": [], "boolean": "true", "nullValue": "null"}`,
		expected: []tokenizer.Token{
			{Type: tokenizer.TokenBraceOpen, Value: "{"},
			{Type: tokenizer.TokenString, Value: "id"},
			{Type: tokenizer.TokenColon, Value: ":"},
			{Type: tokenizer.TokenNumber, Value: "120391"},
			{Type: tokenizer.TokenComma, Value: ","},
			{Type: tokenizer.TokenString, Value: "name"},
			{Type: tokenizer.TokenColon, Value: ":"},
			{Type: tokenizer.TokenString, Value: "Some Name"},
			{Type: tokenizer.TokenComma, Value: ","},
			{Type: tokenizer.TokenString, Value: "age"},
			{Type: tokenizer.TokenColon, Value: ":"},
			{Type: tokenizer.TokenNumber, Value: "20"},
			{Type: tokenizer.TokenComma, Value: ","},
			{Type: tokenizer.TokenString, Value: "something"},
			{Type: tokenizer.TokenColon, Value: ":"},
			{Type: tokenizer.TokenSquareOpen, Value: "["},
			{Type: tokenizer.TokenSquareClose, Value: "]"},
			{Type: tokenizer.TokenComma, Value: ","},
			{Type: tokenizer.TokenString, Value: "boolean"},
			{Type: tokenizer.TokenColon, Value: ":"},
			{Type: tokenizer.TokenBool, Value: "true"},
			{Type: tokenizer.TokenComma, Value: ","},
			{Type: tokenizer.TokenString, Value: "nullValue"},
			{Type: tokenizer.TokenColon, Value: ":"},
			{Type: tokenizer.TokenNull, Value: "null"},
			{Type: tokenizer.TokenBraceClose, Value: "}"},
		},
	}
	result := tokenizer.Tokenizer(test.input)
	fmt.Println(result.IsEqual(test.expected))
}
