package main

import (
	"fmt"

	"jsonparser/parser"
	"jsonparser/tokenizer"
)

func main() {
	json := `
		{
			"name": "Jude",
			"age": "20",
			"hobbies": ["coding", "gaming", "gym"],
			"love": {
				"name": "Sara",
				"age": 19
			}
		}`
	result, err := parser.Parse(tokenizer.Tokenizer(json))
	if err != nil {
		fmt.Println(err.Error())
	}
	for k, v := range result {
		fmt.Println(k+" :", v)
	}
}
