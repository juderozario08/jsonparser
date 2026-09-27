package tokenizer

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
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

func (t TokenType) String() string {
	switch t {
	case TokenBraceOpen:
		return "TokenBraceOpen"
	case TokenBraceClose:
		return "TokenBraceClose"
	case TokenBracketOpen:
		return "TokenBracketOpen"
	case TokenBracketClose:
		return "TokenBracketClose"
	case TokenSquareOpen:
		return "TokenSquareOpen"
	case TokenSquareClose:
		return "TokenSquareClose"
	case TokenString:
		return "TokenString"
	case TokenNumber:
		return "TokenNumber"
	case TokenComma:
		return "TokenComma"
	case TokenColon:
		return "TokenColon"
	case TokenBool:
		return "TokenBool"
	case TokenNull:
		return "TokenNull"
	default:
		return fmt.Sprintf("TokenType(%d)", t)
	}
}

type Token struct {
	Type  TokenType
	Value string
}

type Tokens []Token

func (tokens Tokens) IsEqual(otherTokens Tokens) bool {
	if len(tokens) != len(otherTokens) {
		return false
	}
	for i, token := range tokens {
		if token != otherTokens[i] {
			return false
		}
	}
	return true
}

func Tokenizer(input string) Tokens {
	tokens := make(Tokens, 0)
	n := len(input)

	for i := 0; i < n; {
		c := input[i]

		// Skip whitespace
		if c == ' ' || c == '\t' || c == '\n' || c == '\r' {
			i++
			continue
		}

		switch c {
		case '{':
			tokens = append(tokens, Token{Type: TokenBraceOpen, Value: "{"})
			i++
		case '}':
			tokens = append(tokens, Token{Type: TokenBraceClose, Value: "}"})
			i++
		case '[':
			tokens = append(tokens, Token{Type: TokenSquareOpen, Value: "["})
			i++
		case ']':
			tokens = append(tokens, Token{Type: TokenSquareClose, Value: "]"})
			i++
		case '(':
			tokens = append(tokens, Token{Type: TokenBracketOpen, Value: "("})
			i++
		case ')':
			tokens = append(tokens, Token{Type: TokenBracketClose, Value: ")"})
			i++
		case ',':
			tokens = append(tokens, Token{Type: TokenComma, Value: ","})
			i++
		case ':':
			tokens = append(tokens, Token{Type: TokenColon, Value: ":"})
			i++
		case '"':
			// Quoted string
			i++ // Skip opening quote
			var sb strings.Builder
			for i < n && input[i] != '"' {
				if input[i] == '\\' && i+1 < n {
					i++
					switch input[i] {
					case '"':
						sb.WriteByte('"')
					case '\\':
						sb.WriteByte('\\')
					case '/':
						sb.WriteByte('/')
					case 'b':
						sb.WriteByte('\b')
					case 'f':
						sb.WriteByte('\f')
					case 'n':
						sb.WriteByte('\n')
					case 'r':
						sb.WriteByte('\r')
					case 't':
						sb.WriteByte('\t')
					case 'u':
						if i+4 < n {
							hex := input[i+1 : i+5]
							if code, err := strconv.ParseInt(hex, 16, 32); err == nil {
								sb.WriteRune(rune(code))
								i += 4
							} else {
								sb.WriteString("\\u" + hex)
							}
						} else {
							sb.WriteString("\\u")
						}
					default:
						sb.WriteByte(input[i])
					}
				} else {
					sb.WriteByte(input[i])
				}
				i++
			}
			if i < n && input[i] == '"' {
				i++ // Skip closing quote
			}

			val := sb.String()
			// Check if the quoted string represents a number, boolean, or null for compatibility
			if _, err := strconv.ParseFloat(val, 64); err == nil {
				tokens = append(tokens, Token{Type: TokenNumber, Value: val})
			} else {
				switch strings.ToLower(val) {
				case "false", "true":
					tokens = append(tokens, Token{Type: TokenBool, Value: strings.ToLower(val)})
				case "null", "nil":
					tokens = append(tokens, Token{Type: TokenNull, Value: val})
				default:
					tokens = append(tokens, Token{Type: TokenString, Value: val})
				}
			}

		default:
			// Unquoted number (e.g. 19, -20.5, 1e-4)
			if c == '-' || unicode.IsDigit(rune(c)) {
				start := i
				if input[i] == '-' {
					i++
				}
				for i < n && unicode.IsDigit(rune(input[i])) {
					i++
				}
				if i < n && input[i] == '.' {
					i++
					for i < n && unicode.IsDigit(rune(input[i])) {
						i++
					}
				}
				if i < n && (input[i] == 'e' || input[i] == 'E') {
					i++
					if i < n && (input[i] == '+' || input[i] == '-') {
						i++
					}
					for i < n && unicode.IsDigit(rune(input[i])) {
						i++
					}
				}
				numStr := input[start:i]
				tokens = append(tokens, Token{Type: TokenNumber, Value: numStr})
			} else if unicode.IsLetter(rune(c)) {
				// Unquoted identifier: true, false, null, nil
				start := i
				for i < n && unicode.IsLetter(rune(input[i])) {
					i++
				}
				word := input[start:i]
				switch strings.ToLower(word) {
				case "true", "false":
					tokens = append(tokens, Token{Type: TokenBool, Value: strings.ToLower(word)})
				case "null", "nil":
					tokens = append(tokens, Token{Type: TokenNull, Value: strings.ToLower(word)})
				default:
					tokens = append(tokens, Token{Type: TokenString, Value: word})
				}
			} else {
				// Unknown character, skip safely
				i++
			}
		}
	}

	return tokens
}
