package cli

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/antham/goller/dsl"
	"github.com/antham/goller/transformer"
	"gopkg.in/alecthomas/kingpin.v2"
)

var transformersGlobal *transformer.Transformers

// Transformers is a map of statement sort by position
type Transformers struct {
	transformers *transformer.Transformers
}

// Init all starting states
func (t *Transformers) Init() {
	transformersGlobal = transformer.NewTransformers()
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

// ValidatePositions against extracted transformers
func (t *Transformers) ValidatePositions(positions *[]int) error {
	if t.transformers != nil {
		for transPosition := range *t.transformers {
			positionMatch := false

			for _, position := range *positions {
				if transPosition == position {
					positionMatch = true
				}
			}

			if !positionMatch {
				return fmt.Errorf("Transformer is wrong : position %d doesn't exist", transPosition)
			}
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
	target.Init()
	s.SetValue(target)
	return
}
