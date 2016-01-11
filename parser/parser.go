package parser

import (
	"bytes"
	"github.com/antham/goller/dsl"
	"gopkg.in/alecthomas/kingpin.v2"
	"regexp"
	"strings"
)

type parser func(string) []string

// Parser represent a parser statement
type Parser parser

// Set is used to populate statement from string
func (p *Parser) Set(value string) error {
	dslParser := dsl.NewParser(bytes.NewBufferString(value))

	stmt, err := dslParser.ParseFunction()

	if err != nil {
		return err
	}

	(*p).Create(stmt.Name, stmt.Args)

	return nil
}

// Create a new parser from specified arguments
func (p *Parser) Create(pars string, args []string) {
	var function parser

	switch len(args) {
	case 0:
		switch pars {
		case "whi":
			{
				function = strings.Fields
			}
		}
	case 1:
		switch pars {
		case "reg":
			{
				{
					function = func(input string) []string {
						re := regexp.MustCompile(args[0])

						matches := re.FindStringSubmatch(input)

						if len(matches) > 1 {
							return matches[1:]
						}

						return []string{}
					}
				}

			}
		}
	}

	if function != nil {
		*p = (Parser)(function)
	}
}

// Parse string in array of string
func (p *Parser) Parse(input string) []string {
	return (*p)(input)
}

// String
func (p *Parser) String() string {
	return ""
}

// Wrapper is used to transform argument from command line
func Wrapper(s kingpin.Settings) (target *Parser) {
	target = new(Parser)
	s.SetValue((*Parser)(target))
	return
}
