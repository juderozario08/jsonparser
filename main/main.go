package main

import (
	"fmt"

	"jsonparser/parser"
	"jsonparser/tokenizer"
)

func main() {
	test := `{"people":[{"name":"Jude", "age": "20"},{"name": "Sara", "age": "20"}]}`
	tokens := tokenizer.Tokenizer(test)
	result, err := parser.Parse(tokens)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)
}
