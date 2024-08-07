package parser

import (
	"errors"
	"fmt"
	"jsonparser/tokenizer"
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
		fmt.Println(token)
		if tokenizer.TokenComma == token.Type {
			continue
		} else if token.Type != tokenizer.TokenString ||
			tokens[i+1].Type != tokenizer.TokenColon {
			return nil, errors.New("KEY must be a string followed by a colon")
		} else {
			node, err := ParseAndValidate(&tokens, &i)
			if err != nil {
				return nil, err
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
		st := stack.New()
		arrayTokens := make(tokenizer.Tokens, 0)
		for ; st.Len() > 0 && *i < len(*tokens)-1; *i++ {
			err := BracketCheck((*tokens)[*i], &st)
			if err != nil {
				return nil, err
			}
			arrayTokens = append(arrayTokens, (*tokens)[*i])
		}
		value, err := ParseAndValidateArray(arrayTokens)
		if err != nil {
			return nil, err
		}
		return ArrayNode{Key: key, Value: value}, nil

	case tokenizer.TokenBraceOpen:
		st := stack.New()
		objectTokens := make(tokenizer.Tokens, 0)
		for ; st.Len() > 0 && *i < len(*tokens)-1; *i++ {
			err := BracketCheck((*tokens)[*i], &st)
			if err != nil {
				return nil, err
			}
			objectTokens = append(objectTokens, (*tokens)[*i])
		}
		value, err := ParseAndValidateObject(objectTokens)
		if err != nil {
			return nil, err
		}
		return ObjectNode{Key: key, Value: value}, nil

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
	return nil, errors.New("Something went wrong when validating and parsing")
}

func ParseAndValidateObject(tokens tokenizer.Tokens) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	return res, nil
}

func ParseAndValidateArray(tokens tokenizer.Tokens) ([]interface{}, error) {
	res := make([]interface{}, 0)
	return res, nil
}

func BracketCheck(token tokenizer.Token, st **stack.Stack) error {
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
			return errors.New("Wrong syntax for object. Square brackets do not match")
		}
	case tokenizer.TokenBraceClose:
		popped := (*st).Pop()
		if popped == nil || popped != tokenizer.TokenBraceOpen {
			return errors.New("Wrong syntax for object. Curly braces do not match")
		}
	case tokenizer.TokenBracketClose:
		popped := (*st).Pop()
		if popped == nil || popped != tokenizer.TokenBracketOpen {
			return errors.New("Wrong syntax for object. Brackets do not match")
		}
	}
	return nil
}

func CheckComma(tokens *tokenizer.Tokens, i *int) {

}
