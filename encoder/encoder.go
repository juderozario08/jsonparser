package encoder

import (
	"strconv"
)

var depth int

func Encoder(decoded map[string]interface{}) (string, error) {
	result := "{"
	maxLength := len(decoded)
	i := 0
	for k, v := range decoded {
		result += "\n    "
		for j := 0; j < depth; j++ {
			result += "    "
		}
		result += `"` + k + `"`
		switch t := v.(type) {
		case float64:
			str := strconv.FormatFloat(t, 'f', 16, 64)
			index := 0
			isInt := false
			for j := len(str) - 1; j >= 0; j-- {
				if string(str[j]) != "0" {
					index = j
					if string(str[j]) == "." {
						isInt = true
					}
					break
				}
			}
			res := ""
			if isInt {
				for j := 0; j < index; j++ {
					res += string(str[j])
				}
			} else {
				for j := 0; j <= index; j++ {
					res += string(str[j])
				}
			}
			result += `: "` + res + `"`
		case string:
			result += `: "` + t + `"`
		case bool:
			if t {
				result += `: "true"`
			} else {
				result += `: "false"`
			}
		case []interface{}:
			val, err := EncodeArray(t)
			if err != nil {
				return "", err
			}
			result += ": " + val
		case map[string]interface{}:
			depth++
			val, err := Encoder(t)
			if err != nil {
				return "", err
			}
			result += ": " + val
		default:
			result += `: "nil"`
		}
		if i < maxLength-1 {
			result += ","
		}
		i++
	}
	result += "\n"
	for j := 0; j < depth; j++ {
		result += "    "
	}
	depth--
	return result + "}", nil
}

func EncodeArray(array []interface{}) (string, error) {
	result := "["
	maxLength := len(array)
	i := 0
	for _, el := range array {
		switch t := el.(type) {
		case float64:
			str := strconv.FormatFloat(t, 'f', 16, 64)
			index := 0
			isInt := false
			for j := len(str) - 1; j >= 0; j-- {
				if string(str[j]) != "0" {
					index = j
					if string(str[j]) == "." {
						isInt = true
					}
					break
				}
			}
			res := ""
			if isInt {
				for j := 0; j < index; j++ {
					res += string(str[j])
				}
			} else {
				for j := 0; j <= index; j++ {
					res += string(str[j])
				}
			}
			result += `"` + res + `"`
		case string:
			result += `"` + t + `"`
		case bool:
			if t {
				result += `"true"`
			} else {
				result += `"false"`
			}
		case []interface{}:
			val, err := EncodeArray(t)
			if err != nil {
				return "", err
			}
			result += val
		case map[string]interface{}:
			depth += 2
			val, err := Encoder(t)
			if err != nil {
				return "", err
			}
			depth--
			result += val
		default:
			result += `"nil"`
		}
		if i < maxLength-1 {
			result += ","
		}
		i++
	}
	return result + "]", nil
}
