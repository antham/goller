package dsl

// Token represents a entity extracted from string parsing
type Token int

const (
	ILLEGAL Token = iota
	EOF

	NUMBER
	ALNUM
	STRING
	OPAREN
	CPAREN
	COMMA
	COLON
	DQUOTE
	EDQUOTE
	PIPE
	DOT
)
