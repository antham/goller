package transformer

import (
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
		}
	case 1:
		switch trans {
		case "trim":
			function = func(input string) string {
				return strings.Trim(input, args[0])
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