package main

import (
	"fmt"
	"jsonparser/parser"
	"jsonparser/tokenizer"
)

/* type TokenizerTest struct {
	input    string
	expected []tokenizer.Token
} */

func main() {
	// test := TokenizerTest{
	// 	input: `{"id":"120391", "name": "Some Name", "age": "20", "something": [], "boolean": "true", "nullValue": "null"}`,
	// 	expected: []tokenizer.Token{
	// 		{Type: tokenizer.TokenBraceOpen, Value: "{"},
	// 		{Type: tokenizer.TokenString, Value: "id"},
	// 		{Type: tokenizer.TokenColon, Value: ":"},
	// 		{Type: tokenizer.TokenNumber, Value: "120391"},
	// 		{Type: tokenizer.TokenComma, Value: ","},
	// 		{Type: tokenizer.TokenString, Value: "name"},
	// 		{Type: tokenizer.TokenColon, Value: ":"},
	// 		{Type: tokenizer.TokenString, Value: "Some Name"},
	// 		{Type: tokenizer.TokenComma, Value: ","},
	// 		{Type: tokenizer.TokenString, Value: "age"},
	// 		{Type: tokenizer.TokenColon, Value: ":"},
	// 		{Type: tokenizer.TokenNumber, Value: "20"},
	// 		{Type: tokenizer.TokenComma, Value: ","},
	// 		{Type: tokenizer.TokenString, Value: "something"},
	// 		{Type: tokenizer.TokenColon, Value: ":"},
	// 		{Type: tokenizer.TokenSquareOpen, Value: "["},
	// 		{Type: tokenizer.TokenSquareClose, Value: "]"},
	// 		{Type: tokenizer.TokenComma, Value: ","},
	// 		{Type: tokenizer.TokenString, Value: "boolean"},
	// 		{Type: tokenizer.TokenColon, Value: ":"},
	// 		{Type: tokenizer.TokenBool, Value: "true"},
	// 		{Type: tokenizer.TokenComma, Value: ","},
	// 		{Type: tokenizer.TokenString, Value: "nullValue"},
	// 		{Type: tokenizer.TokenColon, Value: ":"},
	// 		{Type: tokenizer.TokenNull, Value: "null"},
	// 		{Type: tokenizer.TokenBraceClose, Value: "}"},
	// 	},
	// }
	// result := tokenizer.Tokenizer(test.input)
	// fmt.Println(result.IsEqual(test.expected))
	TestArrayParser()
}

func TestArrayParser() {
	tests :=
		tokenizer.Tokens{
			{Type: tokenizer.TokenSquareOpen, Value: "["},
			{Type: tokenizer.TokenString, Value: "Jude"},
			{Type: tokenizer.TokenString, Value: "Sara"},
			{Type: tokenizer.TokenNumber, Value: "2"},
			{Type: tokenizer.TokenBool, Value: "true"},
			{Type: tokenizer.TokenNull, Value: "nil"},
			{Type: tokenizer.TokenBool, Value: "false"},
			{Type: tokenizer.TokenSquareClose, Value: "]"},
		}
	expected := []interface{}{"Jude", "Sara", 2.0, true, nil, false}
	i := 0
	result, err := parser.ParseAndValidateArray(&tests, &i)
	if err != nil {
		fmt.Println(err.Error())
	}
	if !isEqual(result, expected) {
		fmt.Printf("Test failed, for %v. Want \n%v and got \n%v", tests, result, expected)
	}
}

func isEqual(res []interface{}, expected []interface{}) bool {
	if len(res) != len(expected) {
		fmt.Println("Size does not match")
		return false
	}
	for i := 0; i < len(res); i++ {
		if res[i] != expected[i] {
			fmt.Println(res[i], expected[i])
			return false
		}
	}
	return true
}
