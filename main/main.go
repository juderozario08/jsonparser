package main

import (
	"fmt"

	"jsonparser/encoder"
	"jsonparser/parser"
	"jsonparser/tokenizer"
)

func main() {
	decoded := map[string]interface{}{
		"name": "Jude",
		"age":  20.0123,
		"courses": []interface{}{
			map[string]interface{}{
				"name": "Jude",
				"age":  20.01,
			}, map[string]interface{}{
				"name": "Sara",
				"age":  20.0,
			},
		},
		"person": map[string]interface{}{
			"NAME":    "JUDE",
			"Partner": "Sara",
			"Person": map[string]interface{}{
				"NAME":    "JUDE",
				"Partner": "Sara",
				"people":  []interface{}{"Jude", "Sara"},
			},
		},
	}
	result, err := encoder.Encoder(decoded)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)
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
	res, err := parser.Parse(tokenizer.Tokenizer(json))
	if err != nil {
		fmt.Println(err.Error())
	}
	for k, v := range res {
		fmt.Println(k+" :", v)
	}
}
