package transformer

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

type transformer func(string) string

// Transformers list
type Transformers struct {
	transformers []transformer
}

// Append transformer to transformer list
func (t *Transformers) Append(trans string, args []string) {
	var function transformer

	switch len(args) {
	case 0:
		switch trans {
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
		switch trans {
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
		case "rcat":
			function = func(input string) string {
				return input + args[0]
			}
		case "lcat":
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
		switch trans {
		case "repl":
			function = func(input string) string {
				return strings.Replace(input, args[0], args[1], -1)
			}
		}
	}

	if function != nil {
		t.transformers = append(t.transformers, function)
	}
}

// Apply transformers to a string
func (t *Transformers) Apply(input string) string {

	result := input

	for _, transformer := range t.transformers {
		result = transformer(result)
	}

	return result
}
