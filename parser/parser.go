package parser

import (
	"errors"
	"strconv"

	"jsonparser/tokenizer"

	"github.com/golang-collections/collections/stack"
)

// Abstract Syntax Tree defining each Node must have a Key and Value
type ASTNode interface {
	GetKey() string
	GetValue() interface{}
}

type ObjectNode struct {
	Key   string
	Value map[string]interface{}
}

type ArrayNode struct {
	Key   string
	Value []interface{}
}

type StringNode struct {
	Key   string
	Value string
}

type NumberNode struct {
	Key   string
	Value float64
}

type BooleanNode struct {
	Key   string
	Value bool
}

type NullNode struct {
	Key string
}

func (o ObjectNode) GetKey() string  { return o.Key }
func (o ArrayNode) GetKey() string   { return o.Key }
func (o StringNode) GetKey() string  { return o.Key }
func (o NumberNode) GetKey() string  { return o.Key }
func (o BooleanNode) GetKey() string { return o.Key }
func (o NullNode) GetKey() string    { return o.Key }

func (o ObjectNode) GetValue() interface{}  { return o.Value }
func (o ArrayNode) GetValue() interface{}   { return o.Value }
func (o StringNode) GetValue() interface{}  { return o.Value }
func (o NumberNode) GetValue() interface{}  { return o.Value }
func (o BooleanNode) GetValue() interface{} { return o.Value }
func (o NullNode) GetValue() interface{}    { return nil }

// Parse parses tokens into a map[string]interface{}. Root must be a JSON object.
func Parse(tokens tokenizer.Tokens) (Value map[string]interface{}, Error error) {
	if len(tokens) < 2 {
		return nil, errors.New("PARSE: JSON brace not closed properly")
	}

	if tokens[0].Type != tokenizer.TokenBraceOpen ||
		tokens[len(tokens)-1].Type != tokenizer.TokenBraceClose {
		return nil, errors.New("PARSE: JSON brace not closed properly")
	}

	return ParseObject(tokens)
}

// ToValue extracts an ASTNode for a key-value pair starting at index *i
func ToValue(tokens *tokenizer.Tokens, i *int) (Node ASTNode, Error error) {
	if *i >= len(*tokens) {
		return nil, errors.New("TO-VALUE: Unexpected end of tokens")
	}
	key := (*tokens)[*i].Value
	*i += 2
	if *i >= len(*tokens) {
		return nil, errors.New("TO-VALUE: Unexpected end of tokens after colon")
	}

	switch tk := (*tokens)[*i].Type; tk {
	case tokenizer.TokenBraceOpen, tokenizer.TokenSquareOpen:
		tkns, err := IsolateArrayAndObject(tokens, i)
		if err != nil {
			return nil, err
		}
		if tk == tokenizer.TokenSquareOpen {
			value, err := ParseArray(tkns)
			if err != nil {
				return nil, err
			}
			return ArrayNode{Key: key, Value: value}, nil
		} else {
			value, err := ParseObject(tkns)
			if err != nil {
				return nil, err
			}
			return ObjectNode{Key: key, Value: value}, nil
		}

	case tokenizer.TokenString:
		return StringNode{Key: key, Value: (*tokens)[*i].Value}, nil

	case tokenizer.TokenNumber:
		num, err := strconv.ParseFloat((*tokens)[*i].Value, 64)
		if err != nil {
			return nil, err
		}
		return NumberNode{Key: key, Value: num}, nil

	case tokenizer.TokenBool:
		return BooleanNode{Key: key, Value: (*tokens)[*i].Value == "true"}, nil

	case tokenizer.TokenNull:
		return NullNode{Key: key}, nil
	}
	return nil, errors.New("TO-VALUE: Something went wrong when validating and parsing")
}

// ParseObject parses tokens representing an object enclosed in { ... }
func ParseObject(tokens tokenizer.Tokens) (Value map[string]interface{}, Error error) {
	if len(tokens) < 2 || tokens[0].Type != tokenizer.TokenBraceOpen || tokens[len(tokens)-1].Type != tokenizer.TokenBraceClose {
		return nil, errors.New("PARSE-OBJECT: Object must start with { and end with }")
	}

	res := make(map[string]interface{})
	if len(tokens) == 2 {
		return res, nil
	}

	needComma := false
	for i := 1; i < len(tokens)-1; i++ {
		token := tokens[i]

		if needComma {
			if token.Type != tokenizer.TokenComma {
				return nil, errors.New("PARSE-OBJECT: Comma needed")
			}
			needComma = false
			continue
		} else if token.Type == tokenizer.TokenComma {
			return nil, errors.New("PARSE-OBJECT: Unexpected comma")
		}

		if token.Type != tokenizer.TokenString {
			return nil, errors.New("PARSE-OBJECT: Key must be a string")
		}
		key := token.Value

		if i+1 >= len(tokens)-1 || tokens[i+1].Type != tokenizer.TokenColon {
			return nil, errors.New("PARSE-OBJECT: Colon needed")
		}

		i += 2
		if i >= len(tokens)-1 {
			return nil, errors.New("PARSE-OBJECT: Value expected after colon")
		}

		switch valToken := tokens[i]; valToken.Type {
		case tokenizer.TokenString, tokenizer.TokenNull, tokenizer.TokenNumber, tokenizer.TokenBool:
			val, err := SimpleValues(valToken)
			if err != nil {
				return nil, err
			}
			res[key] = val
			needComma = true

		case tokenizer.TokenBraceOpen, tokenizer.TokenSquareOpen:
			tkns, err := IsolateArrayAndObject(&tokens, &i)
			if err != nil {
				return nil, err
			}
			if tkns[0].Type == tokenizer.TokenBraceOpen {
				obj, err := ParseObject(tkns)
				if err != nil {
					return nil, err
				}
				res[key] = obj
			} else {
				arr, err := ParseArray(tkns)
				if err != nil {
					return nil, err
				}
				res[key] = arr
			}
			needComma = true

		default:
			return nil, errors.New("PARSE-OBJECT: Unexpected token for value")
		}
	}

	if !needComma {
		return nil, errors.New("PARSE-OBJECT: Trailing comma not allowed")
	}

	return res, nil
}

// ParseArray parses tokens representing an array enclosed in [ ... ]
func ParseArray(tokens tokenizer.Tokens) (Value []interface{}, Error error) {
	if len(tokens) < 2 || tokens[0].Type != tokenizer.TokenSquareOpen || tokens[len(tokens)-1].Type != tokenizer.TokenSquareClose {
		return nil, errors.New("PARSE-ARRAY: Array must start with [ and end with ]")
	}

	res := make([]interface{}, 0)
	if len(tokens) == 2 {
		return res, nil
	}

	needComma := false
	for i := 1; i < len(tokens)-1; i++ {
		token := tokens[i]

		if needComma {
			if token.Type != tokenizer.TokenComma {
				return nil, errors.New("PARSE-ARRAY: Comma needed")
			}
			needComma = false
			continue
		} else if token.Type == tokenizer.TokenComma {
			return nil, errors.New("PARSE-ARRAY: Unexpected comma")
		}

		switch token.Type {
		case tokenizer.TokenString, tokenizer.TokenNumber, tokenizer.TokenNull, tokenizer.TokenBool:
			val, err := SimpleValues(token)
			if err != nil {
				return nil, err
			}
			res = append(res, val)
			needComma = true

		case tokenizer.TokenBraceOpen, tokenizer.TokenSquareOpen:
			tkns, err := IsolateArrayAndObject(&tokens, &i)
			if err != nil {
				return nil, err
			}
			if tkns[0].Type == tokenizer.TokenBraceOpen {
				obj, err := ParseObject(tkns)
				if err != nil {
					return nil, err
				}
				res = append(res, obj)
			} else {
				arr, err := ParseArray(tkns)
				if err != nil {
					return nil, err
				}
				res = append(res, arr)
			}
			needComma = true

		default:
			return nil, errors.New("PARSE-ARRAY: Unexpected token")
		}
	}

	if !needComma {
		return nil, errors.New("PARSE-ARRAY: Trailing comma not allowed")
	}

	return res, nil
}

// BracketCheck validates that closing brackets/braces match opening brackets/braces using the stack
func BracketCheck(token tokenizer.Token, st **stack.Stack) (Error error) {
	if st == nil || *st == nil {
		return errors.New("BRACKET-CHECK: Stack is nil")
	}
	switch token.Type {
	case tokenizer.TokenSquareOpen, tokenizer.TokenBraceOpen, tokenizer.TokenBracketOpen:
		(*st).Push(token.Type)
	case tokenizer.TokenSquareClose:
		popped := (*st).Pop()
		if popped == nil || popped != tokenizer.TokenSquareOpen {
			return errors.New("BRACKET-CHECK: Wrong syntax for object. Square brackets do not match")
		}
	case tokenizer.TokenBraceClose:
		popped := (*st).Pop()
		if popped == nil || popped != tokenizer.TokenBraceOpen {
			return errors.New("BRACKET-CHECK: Wrong syntax for object. Curly braces do not match")
		}
	case tokenizer.TokenBracketClose:
		popped := (*st).Pop()
		if popped == nil || popped != tokenizer.TokenBracketOpen {
			return errors.New("BRACKET-CHECK: Wrong syntax for object. Brackets do not match")
		}
	}
	return nil
}

// IsolateArrayAndObject isolates all tokens that form a single balanced array or object
func IsolateArrayAndObject(tokens *tokenizer.Tokens, i *int) (Tokens tokenizer.Tokens, Error error) {
	if *i >= len(*tokens) {
		return nil, errors.New("ISOLATE: Index out of range")
	}
	st := stack.New()
	st.Push((*tokens)[*i].Type)
	tkns := make(tokenizer.Tokens, 0)
	tkns = append(tkns, (*tokens)[*i])
	*i++
	for st.Len() > 0 && *i < len(*tokens) {
		token := (*tokens)[*i]
		err := BracketCheck(token, &st)
		if err != nil {
			return nil, err
		}
		tkns = append(tkns, token)
		if st.Len() == 0 {
			break
		}
		*i++
	}
	if st.Len() > 0 {
		return nil, errors.New("ISOLATE: Unclosed object or array")
	}
	return tkns, nil
}

// SimpleValues handles scalar values (strings, numbers, booleans, and null)
func SimpleValues(token tokenizer.Token) (Value interface{}, Error error) {
	switch token.Type {
	case tokenizer.TokenString:
		return token.Value, nil
	case tokenizer.TokenNumber:
		num, err := strconv.ParseFloat(token.Value, 64)
		if err != nil {
			return nil, err
		}
		return num, nil
	case tokenizer.TokenBool:
		return token.Value == "true", nil
	case tokenizer.TokenNull:
		return nil, nil
	default:
		return nil, errors.New("SIMPLE-VALUES: Token is not a simple value")
	}
}
