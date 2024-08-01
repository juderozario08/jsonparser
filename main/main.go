package main

import (
	"fmt"
	"jsonparser"
)

func main() {
	tests := []string{
		`{
			"key": "value",
			"key2": "100"
			"key3": "null"
			"key4": []
		}
		`,
	}
	for _, test := range tests {
		for _, token := range parser.Tokenizer(test) {
			fmt.Println(token.Value)
		}
	}
}
