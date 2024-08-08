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

func Parse(tokens tokenizer.Tokens) (Value map[string]interface{}, Error error) {
	result := make(map[string]interface{})

	if tokens[0].Type != tokenizer.TokenBraceOpen ||
		tokens[len(tokens)-1].Type != tokenizer.TokenBraceClose {
		return nil, errors.New("PARSE: JSON brace not closed properly")
	}

	needComma := false
	for i := 1; i < len(tokens)-1; i++ {

		token := tokens[i]
		if needComma && token.Type != tokenizer.TokenComma {
			return nil, errors.New("PARSE: Comma is missing")
		}
		if tokenizer.TokenComma == token.Type {
			needComma = false
			continue
		}
		if token.Type == tokenizer.TokenString &&
			tokens[i+1].Type == tokenizer.TokenColon {
			node, err := ToValue(&tokens, &i)
			needComma = true
			if err != nil {
				return nil, err
			}
			result[node.GetKey()] = node.GetValue()
		} else {
			return nil, errors.New("PARSE: Key must be a string followed by a colon")
		}

	}
	return result, nil
}

// This function is responsible for returning the value in the proper format for usability
func ToValue(tokens *tokenizer.Tokens, i *int) (Node ASTNode, Error error) {
	key := (*tokens)[*i].Value
	*i += 2
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

// This function is responsible for parsing the object
func ParseObject(tokens tokenizer.Tokens) (Value map[string]interface{}, Error error) {
	res := make(map[string]interface{})
	needComma := false
	for i := 1; i < len(tokens)-1; i++ {
		if tokens[i].Type != tokenizer.TokenComma && needComma {
			return nil, errors.New("PARSE-OBJECT: Comma needed")
		} else if tokens[i].Type == tokenizer.TokenComma && needComma {
			needComma = false
			continue
		}
		key := tokens[i].Value
		if tokens[i+1].Type != tokenizer.TokenColon {
			return nil, errors.New("PARSE-OBJECT: Colon needed")
		}
		i += 2
		switch token := tokens[i]; token.Type {
		case tokenizer.TokenString, tokenizer.TokenNull, tokenizer.TokenNumber, tokenizer.TokenBool:
			val, err := SimpleValues(token)
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
		}
	}
	return res, nil
}

/* This function is responsible for specifically parsing an array */
func ParseArray(tokens tokenizer.Tokens) (Value []interface{}, Error error) {
	res := make([]interface{}, 0)
	needComma := false
	for i := 1; i < len(tokens)-1; i++ {
		if tokens[i].Type != tokenizer.TokenComma && needComma {
			return nil, errors.New("PARSE-ARRAY: Comma needed")
		} else if tokens[i].Type == tokenizer.TokenComma && needComma {
			needComma = false
			continue
		}
		switch token := tokens[i]; token.Type {
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
		}
	}
	return res, nil
}

// This function is responsible for checking if the bracket syntax is correct from the input
func BracketCheck(token tokenizer.Token, st **stack.Stack) (Error error) {
	switch token.Type {
	case tokenizer.TokenSquareOpen:
		(*st).Push(token.Type)
	case tokenizer.TokenBraceOpen:
		(*st).Push(token.Type)
	case tokenizer.TokenBracketOpen:
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

// This function is used to isolate all tokens that specifically form an array or an object
func IsolateArrayAndObject(tokens *tokenizer.Tokens, i *int) (Tokens tokenizer.Tokens, Error error) {
	st := stack.New()
	st.Push((*tokens)[*i].Type)
	tkns := make(tokenizer.Tokens, 0)
	tkns = append(tkns, (*tokens)[*i])
	*i++
	for ; st.Len() > 0 && *i < len(*tokens)-1; *i++ {
		err := BracketCheck((*tokens)[*i], &st)
		if err != nil {
			return nil, err
		}
		tkns = append(tkns, (*tokens)[*i])
	}
	*i--
	return tkns, nil
}

// This function is responsible for handling simple values like
// Strings, Null, Boolean and Numbers and returning the value
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
	default:
		return nil, nil
	}
}
