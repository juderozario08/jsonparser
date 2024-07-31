package main

import (
	"fmt"
	"jsonparser"
)

func main() {
	tests := []struct {
		tokens   []string
		expected bool
	}{
		{[]string{"{", "(", ")", "}"}, true},
		{[]string{"{", "(", "]", "}"}, false},
	}
	for _, test := range tests {
		fmt.Println(parser.ValidateSyntax(test.tokens))
	}
}
