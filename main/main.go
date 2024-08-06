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
	/* tests := tokenizer.Tokens{
		{Type: tokenizer.TokenSquareOpen, Value: "["},
		{Type: tokenizer.TokenSquareOpen, Value: "["},
		{Type: tokenizer.TokenSquareOpen, Value: "["},
		{Type: tokenizer.TokenString, Value: "Jude"},
		{Type: tokenizer.TokenString, Value: "Sara"},
		{Type: tokenizer.TokenSquareClose, Value: "]"},
		{Type: tokenizer.TokenSquareOpen, Value: "["},
		{Type: tokenizer.TokenNumber, Value: "2"},
		{Type: tokenizer.TokenBool, Value: "true"},
		{Type: tokenizer.TokenSquareClose, Value: "]"},
		{Type: tokenizer.TokenSquareOpen, Value: "["},
		{Type: tokenizer.TokenNull, Value: "nil"},
		{Type: tokenizer.TokenBool, Value: "false"},
		{Type: tokenizer.TokenSquareClose, Value: "]"},
		{Type: tokenizer.TokenSquareClose, Value: "]"},
		{Type: tokenizer.TokenSquareOpen, Value: "["},
		{Type: tokenizer.TokenSquareOpen, Value: "["},
		{Type: tokenizer.TokenString, Value: "Jude"},
		{Type: tokenizer.TokenString, Value: "Sara"},
		{Type: tokenizer.TokenSquareClose, Value: "]"},
		{Type: tokenizer.TokenSquareOpen, Value: "["},
		{Type: tokenizer.TokenNumber, Value: "2"},
		{Type: tokenizer.TokenBool, Value: "true"},
		{Type: tokenizer.TokenSquareClose, Value: "]"},
		{Type: tokenizer.TokenSquareOpen, Value: "["},
		{Type: tokenizer.TokenNull, Value: "nil"},
		{Type: tokenizer.TokenBool, Value: "false"},
		{Type: tokenizer.TokenSquareClose, Value: "]"},
		{Type: tokenizer.TokenSquareClose, Value: "]"},
		{Type: tokenizer.TokenSquareOpen, Value: "["},
		{Type: tokenizer.TokenSquareOpen, Value: "["},
		{Type: tokenizer.TokenString, Value: "Jude"},
		{Type: tokenizer.TokenString, Value: "Sara"},
		{Type: tokenizer.TokenSquareClose, Value: "]"},
		{Type: tokenizer.TokenSquareOpen, Value: "["},
		{Type: tokenizer.TokenNumber, Value: "2"},
		{Type: tokenizer.TokenBool, Value: "true"},
		{Type: tokenizer.TokenSquareClose, Value: "]"},
		{Type: tokenizer.TokenSquareOpen, Value: "["},
		{Type: tokenizer.TokenNull, Value: "nil"},
		{Type: tokenizer.TokenBool, Value: "false"},
		{Type: tokenizer.TokenSquareClose, Value: "]"},
		{Type: tokenizer.TokenSquareClose, Value: "]"},
		// {Type: tokenizer.TokenSquareClose, Value: "]"},
	} */
	// test := tokenizer.Tokenizer(`{
	// 	"name":[
	// 			[
	// 				["Jude" "Sara"], ["2", "true"], ["null", "false"]
	// 			],
	// 			[
	// 				["Jude" "Sara"], ["2", "true"], ["null", "false"]
	// 			],
	// 			[
	// 				["Jude" "Sara"], ["2", "true"], ["null", "false"]
	// 			]
	// 		]
	// 	}`)
	// expected := []interface{}{
	// 	[]interface{}{
	// 		[]interface{}{"Jude", "Sara"},
	// 		[]interface{}{2.0, true},
	// 		[]interface{}{nil, false},
	// 	},
	// 	[]interface{}{
	// 		[]interface{}{"Jude", "Sara"},
	// 		[]interface{}{2.0, true},
	// 		[]interface{}{nil, false},
	// 	},
	// 	[]interface{}{
	// 		[]interface{}{"Jude", "Sara"},
	// 		[]interface{}{2.0, true},
	// 		[]interface{}{nil, false},
	// 	},
	// }
	result, err := parser.Parser(tokenizer.Tokenizer(`{
		"id": "6108snoa821601",
		"arr":[
				[
					["Jude" "Sara"], ["2", "true"], ["null", "false"]
				],
				[
					["Jude" "Sara"], ["2", "true"], ["null", "false"]
				],
				[
					["Jude" "Sara"], ["2", "true"], ["null", "false"]
				]
			],
		"age": "20",
		"something": [],
		"nullValue": "null"
		"boolean": "true"
		}`))
	if err != nil {
		fmt.Println(err.Error())
	}
	for k, v := range result {
		fmt.Println(k, v)
	}
	// if IsEqual(result, expected) {
	// 	fmt.Printf("Want \n%v \nand got \n%v\n", result, expected)
	// }
}

func IsEqual(res []interface{}, exp []interface{}) bool {
	if len(res) != len(exp) {
		return false
	}
	for i := 0; i < len(res); i++ {
		switch res[i].(type) {
		case []interface{}:
			if !IsEqual(res[i].([]interface{}), exp[i].([]interface{})) {
				return false
			}
		default:
			if res[i] != exp[i] {
				return false
			}
		}
	}
	return true
}
