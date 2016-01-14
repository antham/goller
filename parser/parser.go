package parser

import (
	"regexp"
	"strings"
)

// Parser represents a function to explode a string in sub part
type Parser func(string) []string

// NewParser from specified arguments
func NewParser(fun string, args []string) *Parser {
	var function Parser

	switch len(args) {
	case 0:
		switch fun {
		case "whi":
			{
				function = strings.Fields
			}
		}
	case 1:
		switch fun {
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

	return &function
}
