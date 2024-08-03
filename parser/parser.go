package parser

import (
	"jsonparser/tokenizer"
)

type ASTNode interface {
	GetType() string
}

type ObjectNode struct {
	Type  string
	Value map[string]ASTNode
}
type ArrayNode struct {
	Type  string
	Value []ASTNode
}

type StringNode struct {
	Type  string
	Value string
}

type NumberNode struct {
	Type  string
	Value float64
}

type BooleanNode struct {
	Type  string
	Value bool
}

type NullNode struct {
	Type string
}

var result map[string]interface{}

func (n ObjectNode) GetType() string  { return n.Type }
func (n ArrayNode) GetType() string   { return n.Type }
func (n StringNode) GetType() string  { return n.Type }
func (n NumberNode) GetType() string  { return n.Type }
func (n BooleanNode) GetType() string { return n.Type }
func (n NullNode) GetType() string    { return n.Type }

func Parser(tokens tokenizer.Tokens) (ASTNode, error) {
	for i := 0; i < len(tokens); i++ {
		token := tokens[i].Type
	}
	return nil, nil
}

// This will handle the logic of parsing all data in a switch statement
func ParseAndValidate() (ASTNode, error) {
	return nil, nil
}

// This will specifically handle how to parse an object
func ParseAndValidateObject() (ASTNode, error) {
	return nil, nil
}

// This will specifially handle how to parse an array
func ParseAndValidateArray() (ASTNode, error) {
	return nil, nil
}
