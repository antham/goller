package cli

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/antham/goller/dsl"
	"github.com/antham/goller/parser"
	"github.com/antham/goller/transformer"
	"gopkg.in/alecthomas/kingpin.v2"
	"strconv"
	"strings"
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
		if stmts.Position == 0 {
			return errors.New("You cannot add a transformer to position 0")
		}

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

// ExtractPositions split positions fields from string
func ExtractPositions(fields string, size int) ([]int, error) {
	var positions []int

	if fields != "" {
		positionDups := make(map[int]bool, 0)

		for _, value := range strings.Split(fields, ",") {
			if position, err := strconv.Atoi(value); err == nil {
				if _, ok := positionDups[position]; ok == true {
					return []int{}, fmt.Errorf("This element is duplicated : %d", position)
				}

				if position >= size+1 {
					return []int{}, fmt.Errorf("Position %d is greater or equal than maximum position %d", position, size+1)
				}

				positionDups[position] = true
				positions = append(positions, position)
			} else {
				return []int{}, fmt.Errorf("%s is not a number", value)
			}
		}
	} else {
		return []int{}, fmt.Errorf("At least 1 element is required")
	}

	return positions, nil
}
