package main

import (
	"fmt"
	"jsonparser/parser"
	"jsonparser/tokenizer"
)

func main() {
	tokens := tokenizer.Tokenizer(`{
		"id": "6108snoa821601",
		"arr":[
				[
					["Jude", "Sara"], ["2", "true"], ["null", "false"]
				],
				[
					["Jude", "Sara"], ["2", "true"], ["null", "false"]
				],
				[
					["Jude", "Sara"], ["2", "true"], ["null", "false"]
				]
			],
		"age": "20",
		"something": [{
			"key": "value",
			"key2": "value2"
		},{
			"key": "value",
			"key2": "value2"
		}],
		"nullValue": "null",
		"boolean": "true"
		}`)
	result, err := parser.Parser(tokens)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result["something"])
}
