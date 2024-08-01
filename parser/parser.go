package parser

import "jsonparser/tokenizer"

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
