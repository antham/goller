package transformer

import (
	"bytes"
	"github.com/antham/goller/dsl"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// TransformersMap is a map of statement sort by position
type TransformersMap map[int]Transformers

// Set is used to populate statement from string
func (t *TransformersMap) Set(value string) error {
	parser := dsl.NewParser(bytes.NewBufferString(value))

	stmts, err := parser.Parse()

	if err != nil {
		return err
	}

	trans := &Transformers{}

	for _, stmt := range stmts.Functions {
		trans.Append(stmt.Name, stmt.Args)

		(*t)[stmts.Position] = *trans
	}

	return nil
}

// String
func (t *TransformersMap) String() string {
	return ""
}

// IsCumulative is used for repeated flags on cli
func (t *TransformersMap) IsCumulative() bool {
	return true
}

// TransformersWrapper is used to transform argument from command line
func TransformersWrapper(s kingpin.Settings) (target *TransformersMap) {
	target = &TransformersMap{}
	s.SetValue((*TransformersMap)(target))
	return
}

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
