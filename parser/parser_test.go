package parser

import (
	"testing"

	"jsonparser/tokenizer"
)

type TestStruct struct {
	Name        string
	Value       string
	Expected    map[string]interface{}
	ExpectError bool
}

type TestArrayStruct struct {
	Name        string
	Value       string
	Expected    []interface{}
	ExpectError bool
}

type TestObjectStruct struct {
	Name        string
	Value       string
	Expected    map[string]interface{}
	ExpectError bool
}

func TestParser(t *testing.T) {
	tests := []TestStruct{
		{
			Name:  "Simple object",
			Value: `{"name":"Jude","age":"20"}`,
			Expected: map[string]interface{}{
				"name": "Jude",
				"age":  20.0,
			},
			ExpectError: false,
		},
		{
			Name:        "Empty object",
			Value:       `{}`,
			Expected:    map[string]interface{}{},
			ExpectError: false,
		},
		{
			Name:  "Object with unquoted primitives",
			Value: `{"name": "Sara", "age": 19, "active": true, "extra": null}`,
			Expected: map[string]interface{}{
				"name":   "Sara",
				"age":    19.0,
				"active": true,
				"extra":  nil,
			},
			ExpectError: false,
		},
		{
			Name:  "Nested object and array",
			Value: `{"user": {"name": "Jude"}, "tags": ["go", "json"]}`,
			Expected: map[string]interface{}{
				"user": map[string]interface{}{
					"name": "Jude",
				},
				"tags": []interface{}{"go", "json"},
			},
			ExpectError: false,
		},
		{
			Name:        "Empty string error",
			Value:       "",
			Expected:    nil,
			ExpectError: true,
		},
		{
			Name:        "Unclosed brace error",
			Value:       `{"name": "Jude"`,
			Expected:    nil,
			ExpectError: true,
		},
		{
			Name:        "Trailing comma error",
			Value:       `{"name": "Jude",}`,
			Expected:    nil,
			ExpectError: true,
		},
		{
			Name:        "Missing colon error",
			Value:       `{"name" "Jude"}`,
			Expected:    nil,
			ExpectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result, err := Parse(tokenizer.Tokenizer(test.Value))
			if test.ExpectError {
				if err == nil {
					t.Errorf("Expected error for %q, got nil", test.Value)
				}
			} else {
				if err != nil {
					t.Fatalf("Unexpected error for %q: %v", test.Value, err)
				}
				if !equalObjects(result, test.Expected) {
					t.Errorf("For %q:\nGot:      %v\nExpected: %v", test.Value, result, test.Expected)
				}
			}
		})
	}
}

func TestArrayParser(t *testing.T) {
	tests := []TestArrayStruct{
		{
			Name:     "String array",
			Value:    `["Jude", "Sara"]`,
			Expected: []interface{}{"Jude", "Sara"},
		},
		{
			Name:  "Nested array with quoted numbers",
			Value: `["Jude", ["20", "30"]]`,
			Expected: []interface{}{"Jude", []interface{}{
				20.0, 30.0,
			}},
		},
		{
			Name:  "Array of objects",
			Value: `[{"name":"Jude", "age": "20"},{"name": "Sara", "age": "20"}]`,
			Expected: []interface{}{
				map[string]interface{}{
					"name": "Jude",
					"age":  20.0,
				},
				map[string]interface{}{
					"name": "Sara",
					"age":  20.0,
				},
			},
		},
		{
			Name:     "Empty array",
			Value:    `[]`,
			Expected: []interface{}{},
		},
		{
			Name:     "Array of unquoted numbers and booleans",
			Value:    `[10, 20.5, true, false, null]`,
			Expected: []interface{}{10.0, 20.5, true, false, nil},
		},
		{
			Name:        "Trailing comma error",
			Value:       `[1, 2,]`,
			ExpectError: true,
		},
		{
			Name:        "Missing comma error",
			Value:       `[1 2]`,
			ExpectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			tokens := tokenizer.Tokenizer(test.Value)
			result, err := ParseArray(tokens)
			if test.ExpectError {
				if err == nil {
					t.Errorf("Expected error for %q, got nil", test.Value)
				}
			} else {
				if err != nil {
					t.Fatalf("Unexpected error for %q: %v", test.Value, err)
				}
				if !equalSlices(result, test.Expected) {
					t.Errorf("For %q:\nGot:      %v\nExpected: %v", test.Value, result, test.Expected)
				}
			}
		})
	}
}

func equalSlices(a []interface{}, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		switch va := a[i].(type) {
		case []interface{}:
			vb, ok := b[i].([]interface{})
			if !ok || !equalSlices(va, vb) {
				return false
			}
		case map[string]interface{}:
			vb, ok := b[i].(map[string]interface{})
			if !ok || !equalObjects(va, vb) {
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
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		bv, exists := b[k]
		if !exists {
			return false
		}
		switch va := v.(type) {
		case []interface{}:
			vb, ok := bv.([]interface{})
			if !ok || !equalSlices(va, vb) {
				return false
			}
		case map[string]interface{}:
			vb, ok := bv.(map[string]interface{})
			if !ok || !equalObjects(va, vb) {
				return false
			}
		default:
			if v != bv {
				return false
			}
		}
	}
	return true
}

func TestObjectParser(t *testing.T) {
	tests := []TestObjectStruct{
		{
			Name:  "Valid nested object and array",
			Value: ` {"person":{"name":"Jude","age":"20"}, "people": ["Jude", "Sara"]}`,
			Expected: map[string]interface{}{
				"person": map[string]interface{}{
					"name": "Jude",
					"age":  20.0,
				},
				"people": []interface{}{"Jude", "Sara"},
			},
			ExpectError: false,
		},
		{
			Name:        "Missing comma error",
			Value:       `{"person":{"name":"Jude","age":"20"} "people": ["Jude", "Sara"]}`,
			Expected:    nil,
			ExpectError: true,
		},
		{
			Name:        "Empty object",
			Value:       `{}`,
			Expected:    map[string]interface{}{},
			ExpectError: false,
		},
		{
			Name:  "Unquoted primitives",
			Value: `{"age": 19, "active": true, "balance": null}`,
			Expected: map[string]interface{}{
				"age":     19.0,
				"active":  true,
				"balance": nil,
			},
			ExpectError: false,
		},
		{
			Name:        "Trailing comma error",
			Value:       `{"a": 1,}`,
			Expected:    nil,
			ExpectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			tokens := tokenizer.Tokenizer(test.Value)
			result, err := ParseObject(tokens)
			if test.ExpectError {
				if err == nil {
					t.Errorf("Expected error for %q, got nil", test.Value)
				}
			} else {
				if err != nil {
					t.Fatalf("Unexpected error for %q: %v", test.Value, err)
				}
				if !equalObjects(result, test.Expected) {
					t.Errorf("For %q:\nGot:      %v\nExpected: %v", test.Value, result, test.Expected)
				}
			}
		})
	}
}
