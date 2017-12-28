package parser

import (
	"regexp"
	"strings"
)

var functions = []map[string]parserEntry{
	{
		"whi": whitespace,
		"clf": clf,
	},
	{
		"spl": split,
		"reg": regexpParser,
	},
}

// parserEntry describe a map entry
type parserEntry func([]string) (Parser, error)

// Parser represents a function to explode a string in sub part
type Parser func(string) []string

// NewParser from specified arguments
func NewParser(fun string, args []string) (*Parser, error) {
	var p Parser
	var err error
	var ok bool
	i := len(args)

	if _, ok = functions[i][fun]; ok {
		p, err = functions[i][fun](args)
	}

	return &p, err
}

// split lines following given string
func split(args []string) (Parser, error) {
	return func(input string) []string {
		return strings.Split(input, args[0])
	}, nil
}

// whitespace split lines following whitespaces
func whitespace(args []string) (Parser, error) {
	return func(input string) []string {
		return strings.Fields(input)
	}, nil
}

// regexp implements regexp parser
func regexpParser(args []string) (Parser, error) {
	return func(input string) []string {
		re := regexp.MustCompile(args[0])

		matches := re.FindStringSubmatch(input)

		if len(matches) > 1 {
			return matches[1:]
		}

		return []string{}
	}, nil
}

// clf implements Common Log Format (NCSA Common log format) parser
func clf(args []string) (Parser, error) {
	return func(input string) []string {
		re := regexp.MustCompile(`(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})\s+(.+?)\s+(.+?)\s+\[(.+?)\]\s+"(.+?)"\s+(\d{3})\s+(\d+)`)

		matches := re.FindStringSubmatch(input)

		if len(matches) == 8 {
			return matches[1:]
		}

		return []string{}
	}, nil
}
