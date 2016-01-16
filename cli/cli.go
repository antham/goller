package cli

import (
	"bytes"
	"github.com/antham/goller/dsl"
	"github.com/antham/goller/parser"
	"github.com/antham/goller/transformer"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Transformers is a map of statement sort by position
type Transformers struct {
	transformers *transformer.Transformers
}

// Set is used to populate statement from string
func (t *Transformers) Set(value string) error {
	parser := dsl.NewParser(bytes.NewBufferString(value))

	stmts, err := parser.ParsePositionAndFunctions()

	if err != nil {
		return err
	}

	(*t).transformers = transformer.NewTransformers()

	for _, stmt := range stmts.Functions {
		(*t).transformers.Append(stmts.Position, stmt.Name, stmt.Args)
	}

	return nil
}

// Get transformers
func (t *Transformers) Get() *transformer.Transformers {
	return t.transformers
}

// String
func (t *Transformers) String() string {
	return ""
}

// IsCumulative is used for repeated flags on cli
func (t *Transformers) IsCumulative() bool {
	return true
}

// TransformersWrapper is used to transform argument from command line
func TransformersWrapper(s kingpin.Settings) (target *Transformers) {
	target = &Transformers{}
	s.SetValue((*Transformers)(target))
	return
}

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
func ParserWrapper(s kingpin.Settings) (target *Parser) {
	target = new(Parser)
	s.SetValue((*Parser)(target))
	return
}
