package parser

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
