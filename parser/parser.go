package parser

import (
	"jsonparser/tokenizer"

	"github.com/golang-collections/collections/stack"
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

func (n ObjectNode) GetType() string  { return n.Type }
func (n ArrayNode) GetType() string   { return n.Type }
func (n StringNode) GetType() string  { return n.Type }
func (n NumberNode) GetType() string  { return n.Type }
func (n BooleanNode) GetType() string { return n.Type }
func (n NullNode) GetType() string    { return n.Type }

// Implement the Parser function
func Parser(tokens tokenizer.Tokens) (ASTNode, error) {
	return nil, nil
}

// This will handle the logic of parsing all data in a switch statement
func ParseValue() (ASTNode, error) {
	return nil, nil
}

// This will specifically handle how to parse an object
func ParseObject() (ASTNode, error) {
	return nil, nil
}

// This will specifially handle how to parse an array
func ParseArray() (ASTNode, error) {
	return nil, nil
}

func ValidateSyntax(tokens tokenizer.Tokens) bool {
	st := stack.New()
	for _, token := range tokens {
		switch token.Type {
		case tokenizer.TokenBraceOpen, tokenizer.TokenBracketOpen:
			st.Push(token.Value)
		case tokenizer.TokenSquareOpen:

		}
	}
	return true
}

func ValidateArraySyntax() bool {
	return false
}

func ValidateObject() bool {
	return false
}
