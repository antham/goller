package cli

import (
	"bytes"
	"errors"
	"github.com/antham/goller/dsl"
	"github.com/antham/goller/transformer"
	"gopkg.in/alecthomas/kingpin.v2"
)

var transformersGlobal *transformer.Transformers

func init() {
	transformersGlobal = transformer.NewTransformers()
}

// Transformers is a map of statement sort by position
type Transformers struct {
	transformers *transformer.Transformers
}

// Set is used to populate statement from string
func (t *Transformers) Set(value string) error {
	var err error

	parser := dsl.NewParser(bytes.NewBufferString(value))

	stmts, err := parser.ParsePositionAndFunctions()

	if err != nil {
		return err
	}

	(*t).transformers = transformersGlobal

	for _, stmt := range stmts.Functions {
		if stmts.Position == 0 {
			return errors.New("You cannot add a transformer to position 0")
		}

		err = (*t).transformers.Append(stmts.Position, stmt.Name, stmt.Args)

		if err != nil {
			return err
		}
	}

	return nil
}

// Get transformers
func (t *Transformers) Get() *transformer.Transformers {
	return t.transformers
}

// String
func (t *Transformers) String() string {
	return ""
}

// IsCumulative is used for repeated flags on cli
func (t *Transformers) IsCumulative() bool {
	return true
}

// TransformersWrapper is used to transform argument from command line
func TransformersWrapper(s kingpin.Settings) (target *Transformers) {
	target = &Transformers{}
	s.SetValue((*Transformers)(target))
	return
}
