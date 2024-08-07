package main

import (
	"fmt"
	"jsonparser/parser"
	"jsonparser/tokenizer"
)

func main() {
	TestObjectParser()
}

func TestArrayParser() {
	result, err := parser.Parser(tokenizer.Tokenizer(`{
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

func TestObjectParser() {
	test := `
		{
			"people": [{ "name": "Jude", "something": [{"hello": "sara"}] }]
		}
	`
	result, err := parser.Parser(tokenizer.Tokenizer(test))
	if err != nil {
		println(err.Error())
	}
	for k, v := range result {
		fmt.Println(k, v)
	}
}
