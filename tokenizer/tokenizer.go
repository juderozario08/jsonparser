package tokenizer

import (
	"strconv"
	"strings"

	"github.com/golang-collections/collections/stack"
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
	TokenBool
	TokenNull
)

type Token struct {
	Type  TokenType
	Value string
}

type Tokens []Token

func Tokenizer(input string) Tokens {
	st := stack.New()
	tokens := make([]Token, 0)
	for i := 0; i < len(input); {
		character := string(input[i])
		switch character {
		case "{":
			tokens = append(tokens, Token{Type: TokenBraceOpen, Value: character})
			st.Push(character)
		case "}":
			tokens = append(tokens, Token{Type: TokenBraceClose, Value: character})
			if st.Pop() == nil {
				panic("Invalid JSON")
			}
		case "[":
			tokens = append(tokens, Token{Type: TokenSquareOpen, Value: character})
			st.Push(character)
		case "]":
			tokens = append(tokens, Token{Type: TokenSquareClose, Value: character})
			if st.Pop() == nil {
				panic("Invalid JSON")
			}
		case "(":
			tokens = append(tokens, Token{Type: TokenBracketOpen, Value: character})
			st.Push(character)
		case ")":
			tokens = append(tokens, Token{Type: TokenBracketClose, Value: character})
			if st.Pop() == nil {
				panic("Invalid JSON")
			}
		case ",":
			tokens = append(tokens, Token{Type: TokenComma, Value: character})
		case ":":
			tokens = append(tokens, Token{Type: TokenColon, Value: character})
		case `"`:
			val := ""
			for i += 1; string(input[i]) != "\""; i++ {
				val += string(input[i])
			}
			_, err := strconv.ParseFloat(val, 64)
			if err != nil {
				switch strings.ToLower(val) {
				case "false", "true":
					tokens = append(tokens, Token{Type: TokenBool, Value: val})
				case "null":
					tokens = append(tokens, Token{Type: TokenNull, Value: val})
				default:
					tokens = append(tokens, Token{Type: TokenString, Value: val})
				}
			} else {
				tokens = append(tokens, Token{Type: TokenNumber, Value: val})
			}
		}
		i++
	}
	return tokens
}

func (tokens *Tokens) IsEqual(otherTokens Tokens) bool {
	if len(*tokens) != len(otherTokens) {
		return false
	}
	for i, token := range *tokens {
		if token != otherTokens[i] {
			return false
		}
	}
	return true
}
