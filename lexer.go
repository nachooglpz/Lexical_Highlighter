package main

import "regexp"

var tokenPatterns = []TokenPattern{
	{Comment, regexp.MustCompile(`\/\*[\s\S]*?\*\/|\/\/.*`)}, //Reference: https://regexr.com/3aeb7. DeanM
	{Keyword, regexp.MustCompile(`\b(class|this|public|static|void|int|if|else|for|while|return|char|boolean|String)\b`)},
	{String, regexp.MustCompile(`\"([^\\"]|\\.)*\"`)},
	{Char, regexp.MustCompile(`\'([^\\']|\\.)*\'`)},
	{Number, regexp.MustCompile(`\b\d+(\.\d+)?\b`)},
	{Operator, regexp.MustCompile(`[+\-*/=<>!&|]+`)},
	{Punctuation, regexp.MustCompile(`[.,;:(){}[\]]`)},
	{Whitespace, regexp.MustCompile(`\s+`)},
	{Identifier, regexp.MustCompile(`\b[a-zA-Z_][a-zA-Z0-9_]*\b`)},
	// {Unknown, regexp.MustCompile(`.`)},
}

func getTokensFromLine(line string) []Token {
	tokens := []Token{}
	input := line

	for len(input) > 0 {
		matched := false

		for _, tokenPattern := range tokenPatterns {
			if loc := tokenPattern.Pattern.FindStringIndex(input); loc != nil && loc[0] == 0 {
				match := input[loc[0]:loc[1]]
				tokens = append(tokens, Token{Type: tokenPattern.Type, Value: match})
				input = input[loc[1]:]
				matched = true
				break

			}
		}

		// Unknown / invalid characters
		if !matched {
			tokens = append(
				tokens,
				Token{Type: Unknown, Value: string(input[0])},
			)
			input = input[1:]
		}
	}

	return tokens
}
