package transformer

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var functions = []map[string]transformerEntry{
	{
		"low": lowercase,
		"upp": uppercase,
		"len": length,
	},
	{
		"trim":  trim,
		"triml": trimLeft,
		"trimr": trimRight,
		"dell":  deleteLeft,
		"delr":  deleteRight,
		"catr":  concatRight,
		"catl":  concatLeft,
		"match": match,
		"add":   add,
		"sub":   subtract,
	},
	{
		"repl": replace,
	},
}

// transformerEntry describe a map entry
type transformerEntry func([]string) (transformer, error)

// transformer represents a transformer function to apply to a string
type transformer func(string) string

// Transformers list
type Transformers map[int][]transformer

// NewTransformers create transformers
func NewTransformers() *Transformers {
	return &Transformers{}
}

// Append transformer to transformer list
func (t *Transformers) Append(position int, fun string, args []string) error {
	var function transformer
	var err error

	if _, ok := functions[len(args)][fun]; ok {
		if function, err = functions[len(args)][fun](args); err != nil {
			return err
		}
	}

	if function == nil {
		return fmt.Errorf(`"%s" doesn't exist or number of arguments "%d" is wrong`, fun, len(args))
	}

	(*t)[position] = append((*t)[position], function)

	return nil
}

// Apply transformers to a string
func (t *Transformers) Apply(position int, input string) string {
	result := input

	if transformers, ok := (*t)[position]; ok {
		for _, transformer := range transformers {
			result = transformer(result)
		}
	}

	return result
}

// lowercase string
func lowercase(args []string) (transformer, error) {
	return strings.ToLower, nil
}

// uppercase string
func uppercase(args []string) (transformer, error) {
	return strings.ToUpper, nil
}

// length give size string
func length(args []string) (transformer, error) {
	return func(input string) string {
		return strconv.Itoa(len(input))
	}, nil
}

// trim remove characters from both side of string
func trim(args []string) (transformer, error) {
	return func(input string) string {
		return strings.Trim(input, args[0])
	}, nil
}

// trimLeft remove characters from left size
func trimLeft(args []string) (transformer, error) {
	return func(input string) string {
		return strings.TrimLeft(input, args[0])
	}, nil
}

// trimRight remove characters from right size
func trimRight(args []string) (transformer, error) {
	return func(input string) string {
		return strings.TrimRight(input, args[0])
	}, nil
}

// deleteLeft remove a given number of element from left side
func deleteLeft(args []string) (transformer, error) {
	size, err := strconv.Atoi(args[0])

	if err != nil {
		return nil, fmt.Errorf("Argument must be an integer, \"%s\" given", args[0])
	}

	return func(input string) string {
		if len(input) < size {
			return ""
		}

		return input[size:]
	}, nil
}

// deleteRight remove a given number of element from right side
func deleteRight(args []string) (transformer, error) {
	size, err := strconv.Atoi(args[0])

	if err != nil {
		return nil, fmt.Errorf("Argument must be an integer, \"%s\" given", args[0])
	}

	return func(input string) string {
		if len(input) < size {
			return ""
		}

		return input[:len(input)-size]
	}, nil
}

// concatRight add a given string to right side of string
func concatRight(args []string) (transformer, error) {
	return func(input string) string {
		return input + args[0]
	}, nil
}

// concatLeft add a given string to left side of string
func concatLeft(args []string) (transformer, error) {
	return func(input string) string {
		return args[0] + input
	}, nil
}

// match return true if regexp match
func match(args []string) (transformer, error) {
	reg, err := regexp.Compile(args[0])

	if err != nil {
		return nil, fmt.Errorf("An error occurred when parsing regexp : \"%s\"", err)
	}

	return func(input string) string {
		return strconv.FormatBool(reg.MatchString(input))
	}, nil
}

// add a given number to input
func add(args []string) (transformer, error) {
	leftOp, err := strconv.Atoi(args[0])

	if err != nil {
		return nil, fmt.Errorf("Argument must be an integer, \"%s\" given", args[0])
	}

	return func(input string) string {
		rightOp, err := strconv.Atoi(input)

		if err != nil {
			log.Fatalf("Argument must be an integer, \"%s\" given", input)
		}

		return strconv.Itoa(rightOp + leftOp)
	}, nil
}

// subtract a given number from input
func subtract(args []string) (transformer, error) {
	leftOp, err := strconv.Atoi(args[0])

	if err != nil {
		return nil, fmt.Errorf("Argument must be an integer, \"%s\" given", args[0])
	}

	return func(input string) string {
		rightOp, err := strconv.Atoi(input)

		if err != nil {
			log.Fatalf("Argument must be an integer, \"%s\" given", input)
		}

		return strconv.Itoa(rightOp - leftOp)
	}, nil
}

// replace a string with another
func replace(args []string) (transformer, error) {
	return func(input string) string {
		return strings.Replace(input, args[0], args[1], -1)
	}, nil
}
