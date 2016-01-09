package dsl

import (
	"fmt"
	"io"
	"strconv"
)

// FunctionStatement represents one parsed function
type FunctionStatement struct {
	Name string
	Args []string
}

// Statement represents all parsed functions from command line
type Statement struct {
	Position  int
	Functions []FunctionStatement
}

// Parser represents a parser.
type Parser struct {
	s   *Scanner
	buf struct {
		tok Token
		lit string
		n   int
	}
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

// Parse extract tokens from string
func (p *Parser) Parse() (*Statement, error) {
	tok, lit := p.scan()

	if tok != NUMBER {
		return nil, fmt.Errorf("found %q, expected a number", lit)
	}

	pos, err := strconv.Atoi(lit)

	if err != nil {
		return nil, err
	}

	if tok, lit = p.scan(); tok != COLON {
		return nil, fmt.Errorf("found %q, expected a colon", lit)
	}

	functionsStmt, err := p.parseFunctions()

	if err != nil {
		return nil, err
	}

	return &Statement{
		Position:  pos,
		Functions: functionsStmt,
	}, nil
}

// parseFuntions extract all functions from string
func (p *Parser) parseFunctions() ([]FunctionStatement, error) {
	functionStmt := []FunctionStatement{}

	for {
		var lit string
		var tok Token

		tok, lit = p.scan()

		if tok != ALNUM {
			return []FunctionStatement{}, fmt.Errorf("found %q, function must have letter and number only", lit)
		}

		args, err := p.parseFuncArgs()

		if err != nil {
			return []FunctionStatement{}, err
		}

		functionStmt = append(functionStmt, FunctionStatement{
			Name: lit,
			Args: args,
		})

		tok, lit = p.scan()

		if tok == EOF {
			break
		}

		if tok != PIPE {
			return []FunctionStatement{}, fmt.Errorf("found %q, function delimiter is a pipe", lit)
		}
	}

	return functionStmt, nil
}

// parseFuncArgs parse all function arguments
func (p *Parser) parseFuncArgs() ([]string, error) {
	args := []string{}

	tok, _ := p.scan()

	if tok != OPAREN {
		p.unscan()
		return args, nil
	}

	for {
		var lit string
		var err error
		var arg string

		arg, err = p.parseFuncArg()

		if err != nil {
			return []string{}, err
		}

		args = append(args, arg)

		tok, lit = p.scan()

		if tok != COMMA && tok != CPAREN {
			return []string{}, fmt.Errorf("found %q, must be a comma or close paren", lit)
		}

		if tok == CPAREN {
			return args, nil
		}
	}
}

// parseFuncArg parse one function argument
func (p *Parser) parseFuncArg() (string, error) {
	var tok Token
	var lit string

	tok, lit = p.scan()

	if tok != DQUOTE {
		return "", fmt.Errorf("found %q, arg must start with a quote", lit)
	}

	var buf string

	for {
		tok, lit = p.scan()

		if tok == EOF {
			return "", fmt.Errorf("found %q, arg must end with a quote", buf)
		}

		if tok != DQUOTE {
			buf += lit
		} else {
			return buf, nil
		}
	}
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok Token, lit string) {
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	tok, lit = p.s.Scan()

	p.buf.tok, p.buf.lit = tok, lit
	return
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() { p.buf.n = 1 }
