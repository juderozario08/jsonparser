package parser

import (
	"fmt"
	"strings"
)

type TokenType int

const (
	TokenBraceOpen TokenType = iota
	TokenBraceClose
	TokenBracketOpen
	TokenBracketClose
	TokenSquareOpen
	TokenSquareClose
	TokenString
	TokenNumber
	TokenComma
	TokenColon
	TokenTrue
	TokenFalse
	TokenNull
)

var tokenVal = map[string]TokenType{
	"{":     TokenBraceOpen,
	"}":     TokenBraceClose,
	"[":     TokenSquareOpen,
	"]":     TokenSquareClose,
	"(":     TokenBracketOpen,
	")":     TokenBracketClose,
	",":     TokenComma,
	":":     TokenColon,
	"true":  TokenTrue,
	"false": TokenFalse,
	"null":  TokenNull,
}

/*
Expected Input:
{
	id: "some-string",
	name: "some-string",
	age: 20,
	something: [],
	boolean: true,
	nullValue: null
}
Expected output:
[
	{ Type: TokenBraceOpen, Value: "{" },
	{ Type: TokenString, Value: "id" },
	{ Type: TokenColon, Value: ":" },
	{ Type: TokenString, Value: "some-string
	{ Type: TokenComma, Value: "," },
	{ Type: TokenString, Value: "name" },
	{ Type: TokenColon, Value: ":" },
	{ Type: TokenString, Value: "some-string
	{ Type: TokenComma, Value: "," },
	{ Type: TokenString, Value: "age" },
	{ Type: TokenColon, Value: ":" },
	{ Type: TokenNumber, Value: "20" },
	{ Type: TokenComma, Value: "," },
	{ Type: TokenString, Value: "something" },
	{ Type: TokenColon, Value: ":" },
	{ Type: TokenSquareOpen, Value: "[" },
	{ Type: TokenSquareClose, Value: "]" },
	{ Type: TokenComma, Value: "," },
	{ Type: TokenString, Value: "boolean" },
	{ Type: TokenColon, Value: ":" },
	{ Type: TokenTrue, Value: "true" },
	{ Type: TokenComma, Value: "," },
	{ Type: TokenString, Value: "nullValue" },
	{ Type: TokenColon, Value: ":" },
	{ Type: TokenNull, Value: "null" },
	{ Type: TokenBraceClose, Value: "}" }
]
*/

type Token struct {
	Type  TokenType
	Value string
}

func ExtractKeywords(tokens []string) string {
	return ""
}

func Tokenize(input string) string {
	return ""
}

func ValidateSyntax(tokens []string) bool {
	stack := make([]string, 0)
	for _, token := range tokens {
		switch token {
		case "{":
			stack = append(stack, token)
		case "[":
			stack = append(stack, token)
		case "(":
			stack = append(stack, token)
		case ")":
			if string(stack[len(stack)-1]) != "(" {
				return false
			}
			stack = stack[:len(stack)-1]
		case "]":
			if string(stack[len(stack)-1]) != "[" {
				return false
			}
			stack = stack[:len(stack)-1]
		case "}":
			if string(stack[len(stack)-1]) != "{" {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return true
}

func Tokenizer(input string) []Token {
	for i := 0; i < len(input); {
		character := string(input[i])
		fmt.Println(character)
		i++
	}
	return []Token{}
}
