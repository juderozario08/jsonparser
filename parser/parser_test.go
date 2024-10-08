package parser

import (
	"testing"

	"jsonparser/tokenizer"
)

type TestStruct struct {
	Value    string
	Expected interface{}
}

type TestArrayStruct struct {
	Value    string
	Expected []interface{}
}

type TestObjectStruct struct {
	Value    string
	Expected map[string]interface{}
}

func TestParser(t *testing.T) {
	tests := []TestStruct{
		{Value: "", Expected: nil},
		{Value: "", Expected: nil},
	}
	for _, test := range tests {
		result, err := Parse(tokenizer.Tokenizer(test.Value))
		if err != nil {
			t.Errorf("Test Failed for %v. Want \n%v and got \n%v", test.Value, result, test.Expected)
		}
	}
}

func TestArrayParser(t *testing.T) {
	tests := []TestArrayStruct{
		{Value: `["Jude", "Sara"]`, Expected: []interface{}{"Jude", "Sara"}},
		{Value: `["Jude", ["20", "30"]]`, Expected: []interface{}{"Jude", []interface{}{
			20.0, 30.0,
		}}},
		{Value: `[{"name":"Jude", "age": "20"},{"name": "Sara", "age": "20"}]`, Expected: []interface{}{
			map[string]interface{}{
				"name": "Jude",
				"age":  20.0,
			},
			map[string]interface{}{
				"name": "Sara",
				"age":  20.0,
			},
		}},
	}
	for _, test := range tests {
		tokens := tokenizer.Tokenizer(test.Value)
		result, err := ParseArray(tokens)
		if err != nil {
			t.Errorf(err.Error())
		}
		if !equalSlices(result, test.Expected) {
			t.Errorf("\nTest:     %v. \nResult:   %v\nExpected: %v\n", test.Value, result, test.Expected)
		}
	}
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
		case map[string]interface{}:
			if !equalObjects(a[i].(map[string]interface{}), b[i].(map[string]interface{})) {
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

func equalObjects(a map[string]interface{}, b map[string]interface{}) bool {
	for k, v := range a {
		switch a[k].(type) {
		case []interface{}:
			if !equalSlices(v.([]interface{}), b[k].([]interface{})) {
				return false
			}
		case map[string]interface{}:
			if !equalObjects(v.(map[string]interface{}), b[k].(map[string]interface{})) {
				return false
			}
		default:
			if v != b[k] {
				return false
			}
		}
	}
	return true
}

func TestObjectParser(t *testing.T) {
	tests := []TestObjectStruct{
		{
			Value: ` {"person":{"name":"Jude","age":"20"}, "people": ["Jude", "Sara"]}`,
			Expected: map[string]interface{}{
				"person": map[string]interface{}{
					"name": "Jude",
					"age":  20.0,
				},
				"people": []interface{}{"Jude", "Sara"},
			},
		},
		{
			Value:    `{"person":{"name":"Jude","age":"20"} "people": ["Jude", "Sara"]}`,
			Expected: nil,
		},
	}
	for _, test := range tests {
		tokens := tokenizer.Tokenizer(test.Value)
		result, err := ParseObject(tokens)
		if err == nil && test.Expected == nil {
			t.Errorf("Program error expected")
		} else if err != nil && test.Expected != nil || !equalObjects(result, test.Expected) {
			t.Errorf("Input: %v\nResult: %v\nExpected: %v\n", test.Value, result, test.Expected)
		}
	}
}
