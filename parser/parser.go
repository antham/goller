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
		case "clf":
			{
				function = clf
			}
		}
	case 1:
		switch fun {
		case "spl":
			{
				function = func(input string) []string {
					return strings.Split(input, args[0])
				}
			}
		case "reg":
			{
				{
					function = reg(args[0])
				}

			}
		}
	}

	return &function
}

// reg implements regexp parser
func reg(r string) Parser {
	return func(input string) []string {
		re := regexp.MustCompile(r)

		matches := re.FindStringSubmatch(input)

		if len(matches) > 1 {
			return matches[1:]
		}

		return []string{}
	}
}

// clf implements Common Log Format (NCSA Common log format) parser
func clf(input string) []string {
	re := regexp.MustCompile(`(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})\s+(.+?)\s+(.+?)\s+\[(.+?)\]\s+"(.+?)"\s+(\d{3})\s+(\d+)`)

	matches := re.FindStringSubmatch(input)

	if len(matches) == 8 {
		return matches[1:]
	}

	return []string{}
}
