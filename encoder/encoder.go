package encoder

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Encoder encodes a Go map[string]interface{} into a formatted JSON string.
func Encoder(decoded map[string]interface{}) (string, error) {
	return encodeObject(decoded, 0)
}

// EncodeArray encodes a Go slice []interface{} into a formatted JSON string.
func EncodeArray(array []interface{}) (string, error) {
	return encodeArray(array, 0)
}

func encodeValue(v interface{}, depth int) (string, error) {
	if v == nil {
		return "null", nil
	}

	switch val := v.(type) {
	case string:
		return strconv.Quote(val), nil

	case bool:
		if val {
			return "true", nil
		}
		return "false", nil

	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64), nil

	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32), nil

	case int:
		return strconv.FormatInt(int64(val), 10), nil
	case int8:
		return strconv.FormatInt(int64(val), 10), nil
	case int16:
		return strconv.FormatInt(int64(val), 10), nil
	case int32:
		return strconv.FormatInt(int64(val), 10), nil
	case int64:
		return strconv.FormatInt(val, 10), nil

	case uint:
		return strconv.FormatUint(uint64(val), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(val), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(val), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(val), 10), nil
	case uint64:
		return strconv.FormatUint(val, 10), nil

	case map[string]interface{}:
		return encodeObject(val, depth)

	case []interface{}:
		return encodeArray(val, depth)

	default:
		return "", fmt.Errorf("unsupported type for JSON encoding: %T", v)
	}
}

func encodeObject(m map[string]interface{}, depth int) (string, error) {
	if len(m) == 0 {
		return "{}", nil
	}

	// Sort keys for deterministic output
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	indent := strings.Repeat("    ", depth+1)
	closingIndent := strings.Repeat("    ", depth)

	var sb strings.Builder
	sb.WriteString("{\n")

	for i, k := range keys {
		valStr, err := encodeValue(m[k], depth+1)
		if err != nil {
			return "", err
		}
		sb.WriteString(indent)
		sb.WriteString(strconv.Quote(k))
		sb.WriteString(": ")
		sb.WriteString(valStr)

		if i < len(keys)-1 {
			sb.WriteString(",")
		}
		sb.WriteString("\n")
	}

	sb.WriteString(closingIndent)
	sb.WriteString("}")
	return sb.String(), nil
}

func encodeArray(arr []interface{}, depth int) (string, error) {
	if len(arr) == 0 {
		return "[]", nil
	}

	// Check if array has complex elements (maps or slices)
	hasComplex := false
	for _, el := range arr {
		switch el.(type) {
		case map[string]interface{}, []interface{}:
			hasComplex = true
		}
	}

	if !hasComplex {
		var parts []string
		for _, el := range arr {
			valStr, err := encodeValue(el, depth)
			if err != nil {
				return "", err
			}
			parts = append(parts, valStr)
		}
		return "[" + strings.Join(parts, ", ") + "]", nil
	}

	indent := strings.Repeat("    ", depth+1)
	closingIndent := strings.Repeat("    ", depth)

	var sb strings.Builder
	sb.WriteString("[\n")
	for i, el := range arr {
		valStr, err := encodeValue(el, depth+1)
		if err != nil {
			return "", err
		}
		sb.WriteString(indent)
		sb.WriteString(valStr)
		if i < len(arr)-1 {
			sb.WriteString(",")
		}
		sb.WriteString("\n")
	}
	sb.WriteString(closingIndent)
	sb.WriteString("]")
	return sb.String(), nil
}
