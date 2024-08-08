package main

import (
	"fmt"
)

func main() {
	/* tokens := tokenizer.Tokenizer(`
		{
			"id": "6108snoa821601",
			"age": "20",
			"boolean": "true",
		"nullValue": {"name": "Jude", "age": "20", "obj": {"null": "false", "values": ["20",
		"30"]}},
			"arr":[["Jude", "Sara"], {"name": "Jude", "age": "20", "null": "null"}, "2.0", "true", "null", "false"],
		}
	`)
	result, err := parser.Parse(tokens)
	if err != nil {
		panic(err.Error())
	}
	for key, value := range result {
		fmt.Println(key, value)
	} */
	a := make([]interface{}, 0)
	b := []interface{}{1, 2, []interface{}{3, 4}}
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, []interface{}{3, 4})
	fmt.Println(equalSlices(a, b))
}

func equalSlices(a []interface{}, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		switch a[i].(type) {
		case []interface{}:
			if !equalSlices(a[i].([]interface{}), b[i].([]interface{})) {
				return false
			}
		default:
			if a[i] != b[i] {
				return false
			}
		}
	}
	return true
}
