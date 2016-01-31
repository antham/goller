package cli

import (
	"bytes"
	"fmt"
	"github.com/antham/goller/dsl"
	"github.com/antham/goller/sorter"
	"gopkg.in/alecthomas/kingpin.v2"
)

var sortersGlobal *sorter.Sorters

// Sorters is a map of statement sort by position
type Sorters struct {
	sorters *sorter.Sorters
}

// Init all starting states
func (s *Sorters) Init() {
	sortersGlobal = sorter.NewSorters()
}

// Set is used to populate statement from string
func (s *Sorters) Set(value string) error {
	parser := dsl.NewParser(bytes.NewBufferString(value))

	stmts, err := parser.ParsePositionsAndFunctions()

	if err != nil {
		return err
	}

	(*s).sorters = sortersGlobal

	for _, stmt := range *stmts {
		(*s).sorters.Append(stmt.Position, stmt.Functions[0].Name, stmt.Functions[0].Args)
	}

	return nil
}

// ValidatePositions against defined sorters
func (s *Sorters) ValidatePositions(positions *[]int) error {
	if s.sorters == nil {
		return nil
	}

	for _, sorter := range *s.sorters {
		positionMatch := false

		for _, position := range *positions {
			if sorter.HasPosition(position) {
				positionMatch = true
			}
		}

		if !positionMatch {
			return fmt.Errorf("Sort is wrong : position %d doesn't exist", sorter.GetPosition())
		}
	}

	return nil
}

// Get sorters
func (s *Sorters) Get() *sorter.Sorters {
	return s.sorters
}

// String
func (s *Sorters) String() string {
	return ""
}

// SortersWrapper is used to transform argument from command line
func SortersWrapper(s kingpin.Settings) (target *Sorters) {
	target = &Sorters{}
	(*Sorters)(target).Init()
	s.SetValue((*Sorters)(target))
	return
}
