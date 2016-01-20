package cli

import (
	"bytes"
	"github.com/antham/goller/dsl"
	"github.com/antham/goller/sorter"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Sorter is a map of statement sort by position
type Sorters struct {
	sorters *sorter.Sorters
}

// Set is used to populate statement from string
func (s *Sorters) Set(value string) error {
	parser := dsl.NewParser(bytes.NewBufferString(value))

	stmts, err := parser.ParsePositionsAndFunctions()

	if err != nil {
		return err
	}

	(*s).sorters = sorter.NewSorters()
	var previousPosition int

	for _, stmt := range *stmts {
		(*s).sorters.Append(previousPosition, stmt.Position, stmt.Functions[0].Name, stmt.Functions[0].Args)
		previousPosition = stmt.Position
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
	s.SetValue((*Sorters)(target))
	return
}
