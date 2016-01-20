package cli

import (
	"bytes"
	"github.com/antham/goller/dsl"
	"github.com/antham/goller/parser"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Parser represent a parser statement
type Parser struct {
	parser *parser.Parser
}

// Set is used to populate statement from string
func (p *Parser) Set(value string) error {
	stmt, err := dsl.NewParser(bytes.NewBufferString(value)).ParseFunction()

	if err != nil {
		return err
	}

	p.parser = parser.NewParser(stmt.Name, stmt.Args)

	return nil
}

// Get parser
func (p *Parser) Get() *parser.Parser {
	return p.parser
}

// String
func (p *Parser) String() string {
	return ""
}

// ParserWrapper is used to transform argument from command line
func ParserWrapper(s *kingpin.ArgClause) (target *Parser) {
	target = new(Parser)
	s.SetValue((*Parser)(target))
	return
}
