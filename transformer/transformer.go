package transformer

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// transformer reprents a transformer function to apply to a string
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

	switch len(args) {
	case 0:
		switch fun {
		case "low":
			function = strings.ToLower
		case "upp":
			function = strings.ToUpper
		case "len":
			function = func(input string) string {
				return strconv.Itoa(len(input))
			}
		}
	case 1:
		switch fun {
		case "trim":
			function = func(input string) string {
				return strings.Trim(input, args[0])
			}
		case "triml":
			function = func(input string) string {
				return strings.TrimLeft(input, args[0])
			}
		case "trimr":
			function = func(input string) string {
				return strings.TrimRight(input, args[0])
			}
		case "dell":
			size, err := strconv.Atoi(args[0])

			if err != nil {
				return fmt.Errorf("Argument must be an integer, \"%s\" given", args[0])
			}

			function = func(input string) string {
				if len(input) < size {
					return ""
				}

				return input[size:]
			}
		case "delr":
			size, err := strconv.Atoi(args[0])

			if err != nil {
				return fmt.Errorf("Argument must be an integer, \"%s\" given", args[0])
			}

			function = func(input string) string {
				if len(input) < size {
					return ""
				}

				return input[:len(input)-size]
			}
		case "catr":
			function = func(input string) string {
				return input + args[0]
			}
		case "catl":
			function = func(input string) string {
				return args[0] + input
			}
		case "match":
			reg, err := regexp.Compile(args[0])

			if err != nil {
				return fmt.Errorf("An error occured when parsing regexp : \"%s\"", err)
			}

			function = func(input string) string {
				return strconv.FormatBool(reg.MatchString(input))
			}
		case "add":
			leftOp, err := strconv.Atoi(args[0])

			if err != nil {
				return fmt.Errorf("Argument must be an integer, \"%s\" given", args[0])
			}

			function = func(input string) string {
				rightOp, err := strconv.Atoi(input)

				if err != nil {
					log.Fatalf("Argument must be an integer, \"%s\" given", input)
				}

				return strconv.Itoa(rightOp + leftOp)
			}
		case "sub":
			leftOp, err := strconv.Atoi(args[0])

			if err != nil {
				return fmt.Errorf("Argument must be an integer, \"%s\" given", args[0])
			}

			function = func(input string) string {
				rightOp, err := strconv.Atoi(input)

				if err != nil {
					log.Fatalf("Argument must be an integer, \"%s\" given", input)
				}

				return strconv.Itoa(rightOp - leftOp)
			}
		}
	case 2:
		switch fun {
		case "repl":
			function = func(input string) string {
				return strings.Replace(input, args[0], args[1], -1)
			}
		}
	}

	if function != nil {
		(*t)[position] = append((*t)[position], function)

		return nil
	}

	return fmt.Errorf("\"%s\" doesn't exists or number of argument is wrong", fun)
}

// Apply transformers to a string
func (t *Transformers) Apply(position int, input string) string {
	result := input

	if transformers, ok := (*t)[position]; ok == true {
		for _, transformer := range transformers {
			result = transformer(result)
		}
	}

	return result
}
