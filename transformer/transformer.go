package transformer

import (
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
func (t *Transformers) Append(position int, fun string, args []string) {
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
			function = func(input string) string {
				size, err := strconv.Atoi(args[0])

				if err != nil {
					log.Fatalf("Argument must be an integer : %s given", args[0])
				}

				if len(input) < size {
					return ""
				}

				return input[size:]
			}
		case "delr":
			function = func(input string) string {
				size, err := strconv.Atoi(args[0])

				if err != nil {
					log.Fatalf("Argument must be an integer : %s given", args[0])
				}

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
			function = func(input string) string {
				result, err := regexp.MatchString(args[0], input)

				if err != nil {
					log.Fatalf("An error occured when parsing regexp : %s", err)
				}

				return strconv.FormatBool(result)
			}
		case "add":
			function = func(input string) string {
				rightOp, err := strconv.Atoi(input)

				if err != nil {
					log.Fatalf("Argument must be an integer %s given", input)
				}

				leftOp, err := strconv.Atoi(args[0])

				if err != nil {
					log.Fatalf("Argument must be an integer : %s given", input)
				}

				return strconv.Itoa(rightOp + leftOp)
			}
		case "sub":
			function = func(input string) string {
				rightOp, err := strconv.Atoi(input)

				if err != nil {
					log.Fatalf("Argument must be an integer %s given", input)
				}

				leftOp, err := strconv.Atoi(args[0])

				if err != nil {
					log.Fatalf("Argument must be an integer : %s given", input)
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
	}
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
