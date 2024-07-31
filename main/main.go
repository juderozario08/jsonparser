package parser

type TokenType int

const (
	TokenBraceOpen TokenType = iota
	TokenBraceClose
	TokenBracketOpen
	TokenBracketClose
	TokenSquareOpen
	TokenSquareClose
	TokenColon
	TokenSemiColon
	TokenComma
	TokenQuote
	TokenSingleQuote
	TokenBackTick
	TokenDollar
)

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
	return false
}
