package main

import "regexp"

type TokenType string

const (
	Keyword     TokenType = "keyword"
	Identifier  TokenType = "identifier"
	Operator    TokenType = "operator"
	Number      TokenType = "number"
	String      TokenType = "string"
	Char        TokenType = "char"
	Comment     TokenType = "comment"
	Whitespace  TokenType = "whitespace"
	Punctuation TokenType = "punctuation"
	Unknown     TokenType = "unknown"
)

type Token struct {
	Type  TokenType
	Value string
}

type TokenPattern struct {
	Type TokenType
	Pattern *regexp.Regexp
}
