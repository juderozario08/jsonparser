package parser

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

import (
	"errors"
	"fmt"
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

var tokenVal = map[any]TokenType{
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
		value, exists := tokenVal[token]
		// Change this to an if condition and fix the test case not working
		if exists {
			switch value {
			case TokenBraceOpen | TokenSquareOpen | TokenBracketOpen:
				Push(&stack, token)
				fmt.Println(stack)
			case TokenBracketClose:
				elem, err := LastElement(&stack)
				if err != nil || tokenVal[elem] != TokenBracketOpen {
					return false
				}
				_, _ = Pop(&stack)
				fmt.Println(stack)
			case TokenSquareClose:
				elem, err := LastElement(&stack)
				if err != nil || tokenVal[elem] != TokenSquareOpen {
					return false
				}
				_, _ = Pop(&stack)
				fmt.Println(stack)
			case TokenBraceClose:
				elem, err := LastElement(&stack)
				if err != nil || tokenVal[elem] != TokenBracketOpen {
					return false
				}
				_, _ = Pop(&stack)
				fmt.Println(stack)
			}
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

func GetValueFromString() {

}

func Push(stack *[]string, value string) {
	*stack = append(*stack, value)
}

func Pop(stack *[]string) (string, error) {
	if len(*stack) > 0 {
		value := (*stack)[len(*stack)-1]
		*stack = (*stack)[:len(*stack)-1]
		return value, nil
	}
	return "", errors.New("ARRAY CONTAINS 0 ELEMENTS")
}

func LastElement(stack *[]string) (string, error) {
	arr := *stack
	if len(arr) != 0 {
		return arr[len(arr)-1], nil
	}
	return "", errors.New("INVALID INDEX")
}
