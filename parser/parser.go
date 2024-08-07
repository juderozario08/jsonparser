package parser

import (
	"errors"
	"jsonparser/tokenizer"
	"runtime"
	"strconv"

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

func Parser(tokens tokenizer.Tokens) (map[string]interface{}, error) {

	result := make(map[string]interface{})

	if tokens[0].Type != tokenizer.TokenBraceOpen ||
		tokens[len(tokens)-1].Type != tokenizer.TokenBraceClose {
		return nil, errors.New("JSON brace not closed properly")
	}

	for i := 1; i < len(tokens)-1; i++ {

		token := tokens[i]
		if token.Type != tokenizer.TokenString &&
			tokens[i+1].Type == tokenizer.TokenColon {
			return nil, errors.New("KEY must be a string followed by a colon")
		} else if token.Type == tokenizer.TokenString &&
			tokens[i+1].Type != tokenizer.TokenColon {
			return nil, errors.New("JSON Syntax is invalid")
		} else {
			if tokenizer.TokenComma == token.Type {
				continue
			}
			node, err := ParseAndValidate(&tokens, &i)
			if err != nil {
				return nil, errors.New(err.Error())
			}
			result[node.GetKey()] = node.GetValue()
		}

	}
	return result, nil
}

func ParseAndValidate(tokens *tokenizer.Tokens, i *int) (ASTNode, error) {
	key := (*tokens)[*i].Value
	*i += 2
	switch (*tokens)[*i].Type {

	case tokenizer.TokenSquareOpen:
		value, err := ParseAndValidateArray(tokens, i)
		if err != nil {
			_, _, line, _ := runtime.Caller(0)
			return nil, errors.New(err.Error() + " called at: " + strconv.Itoa(line))
		}
		return ArrayNode{Key: key, Value: value}, nil

	case tokenizer.TokenBraceOpen:
		value, err := ParseAndValidateObject(tokens, i)
		if err != nil {
			_, _, line, _ := runtime.Caller(0)
			return nil, errors.New(err.Error() + " called at: " + strconv.Itoa(line))
		}
		return ObjectNode{Key: key, Value: value}, nil

	case tokenizer.TokenString:
		return StringNode{Key: key, Value: (*tokens)[*i].Value}, nil

	case tokenizer.TokenNumber:
		num, err := strconv.ParseFloat((*tokens)[*i].Value, 64)
		if err != nil {
			_, _, line, _ := runtime.Caller(0)
			return nil, errors.New(err.Error() + " " + strconv.Itoa(line))
		}
		return NumberNode{Key: key, Value: num}, nil

	case tokenizer.TokenBool:
		return BooleanNode{Key: key, Value: (*tokens)[*i].Value == "true"}, nil

	case tokenizer.TokenNull:
		return NullNode{Key: key}, nil
	}
	_, _, line, _ := runtime.Caller(0)
	return nil, errors.New("Something went wrong when validating and parsing: " + strconv.Itoa(line))
}

func ParseAndValidateObject(tokens *tokenizer.Tokens, i *int) (map[string]interface{}, error) {
	st := stack.New()
	st.Push((*tokens)[*i].Type)
	tkns := make(tokenizer.Tokens, 0)
	tkns = append(tkns, (*tokens)[*i])
	*i++
	for ; st.Len() > 0 && *i < len(*tokens)-1; *i++ {
		tkns = append(tkns, (*tokens)[*i])
		switch (*tokens)[*i].Type {
		case tokenizer.TokenBraceOpen:
			st.Push(tokenizer.TokenBraceOpen)
		case tokenizer.TokenBraceClose:
			if st.Pop() == nil {
				_, _, line, _ := runtime.Caller(0)
				return nil, errors.New("JSON Brace not closed properly line: " + strconv.Itoa(line))
			}
		}
	}
	val, err := Parser(tkns)
	if err != nil {
		_, _, line, _ := runtime.Caller(0)
		return nil, errors.New(err.Error() + " " + strconv.Itoa(line))
	}
	return val, nil
}

func ParseAndValidateArray(tokens *tokenizer.Tokens, i *int) ([]interface{}, error) {
	st := stack.New()
	st.Push((*tokens)[*i])
	*i++
	res := make([]interface{}, 0)

	for ; st.Len() > 0 && (*tokens)[*i].Type != tokenizer.TokenSquareClose && *i < len(*tokens)-1; *i++ {
		switch (*tokens)[*i].Type {
		case tokenizer.TokenComma:
			continue
		case tokenizer.TokenString:
			res = append(res, (*tokens)[*i].Value)
		case tokenizer.TokenNumber:
			num, _ := strconv.ParseFloat((*tokens)[*i].Value, 64)
			res = append(res, num)
		case tokenizer.TokenBool:
			res = append(res, (*tokens)[*i].Value == "true")
		case tokenizer.TokenNull:
			res = append(res, nil)
		case tokenizer.TokenSquareOpen:
			val, err := ParseAndValidateArray(tokens, i)
			if err != nil {
				return nil, errors.New(err.Error())
			}
			res = append(res, val)
		case tokenizer.TokenBraceOpen:
			value, err := ParseAndValidateObject(tokens, i)
			if err != nil {
				return nil, errors.New(err.Error())
			}
			res = append(res, value)
		}
	}

	if st.Pop() == nil || st.Len() > 0 {
		_, _, line, _ := runtime.Caller(0)
		return nil, errors.New("Final bracket syntax for array is invalid " + strconv.Itoa(line))
	}
	return res, nil

}
