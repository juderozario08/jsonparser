package encoder

import (
	"strings"
	"testing"
)

func TestEncoderSimple(t *testing.T) {
	data := map[string]interface{}{
		"name":   "Jude",
		"age":    20.5,
		"count":  42,
		"active": true,
		"extra":  nil,
	}

	result, err := Encoder(data)
	if err != nil {
		t.Fatalf("Encoder returned error: %v", err)
	}

	// Verify keys are present with proper JSON formatting
	expectedSubstrings := []string{
		`"name": "Jude"`,
		`"age": 20.5`,
		`"count": 42`,
		`"active": true`,
		`"extra": null`,
	}

	for _, sub := range expectedSubstrings {
		if !strings.Contains(result, sub) {
			t.Errorf("Expected result to contain %q, but got:\n%s", sub, result)
		}
	}
}

func TestEncoderEmpty(t *testing.T) {
	emptyObj, err := Encoder(map[string]interface{}{})
	if err != nil {
		t.Fatalf("Encoder returned error: %v", err)
	}
	if emptyObj != "{}" {
		t.Errorf("Expected '{}', got %q", emptyObj)
	}

	emptyArr, err := EncodeArray([]interface{}{})
	if err != nil {
		t.Fatalf("EncodeArray returned error: %v", err)
	}
	if emptyArr != "[]" {
		t.Errorf("Expected '[]', got %q", emptyArr)
	}
}

func TestEncoderArray(t *testing.T) {
	arr := []interface{}{"apple", "banana", 123, true, nil}
	result, err := EncodeArray(arr)
	if err != nil {
		t.Fatalf("EncodeArray returned error: %v", err)
	}

	expected := `["apple", "banana", 123, true, null]`
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestEncoderNested(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"name": "Sara",
			"age":  19,
		},
		"scores": []interface{}{100, 95},
	}

	result, err := Encoder(data)
	if err != nil {
		t.Fatalf("Encoder returned error: %v", err)
	}

	expectedSubstrings := []string{
		`"user": {`,
		`"name": "Sara"`,
		`"age": 19`,
		`"scores": [100, 95]`,
	}

	for _, sub := range expectedSubstrings {
		if !strings.Contains(result, sub) {
			t.Errorf("Expected result to contain %q, but got:\n%s", sub, result)
		}
	}
}

func TestEncoderEscapedStrings(t *testing.T) {
	data := map[string]interface{}{
		"quote": "He said \"hello\"\nand left.",
	}

	result, err := Encoder(data)
	if err != nil {
		t.Fatalf("Encoder returned error: %v", err)
	}

	if !strings.Contains(result, `\"hello\"`) {
		t.Errorf("Expected escaped quotes in result, got:\n%s", result)
	}
	if !strings.Contains(result, `\n`) {
		t.Errorf("Expected escaped newline in result, got:\n%s", result)
	}
}

func TestEncoderReentrancy(t *testing.T) {
	data := map[string]interface{}{
		"a": map[string]interface{}{
			"b": 1,
		},
	}

	res1, err1 := Encoder(data)
	if err1 != nil {
		t.Fatalf("First call failed: %v", err1)
	}

	res2, err2 := Encoder(data)
	if err2 != nil {
		t.Fatalf("Second call failed: %v", err2)
	}

	if res1 != res2 {
		t.Errorf("Expected identical results on consecutive calls.\nCall 1:\n%s\nCall 2:\n%s", res1, res2)
	}
}
