package main

import (
	"fmt"
	"jsonparser/parser"
	"jsonparser/tokenizer"
)

func main() {
	tokens := tokenizer.Tokenizer(`
		{
			"id": "6108snoa821601",
			"age": "20",
			"boolean": "true",
			"nullValue": {"name": "Jude", "age": "20", "null": "null"},
			"arr":[ "Jude", "Sara", "2", "true", "null", "false" ],
		}
	`)
	result, err := parser.Parser(tokens)
	if err != nil {
		panic(err.Error())
	}
	for key, value := range result {
		fmt.Println(key, value)
	}
}
