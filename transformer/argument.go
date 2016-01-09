package transformer

import (
	"bytes"
	"github.com/antham/goller/dsl"
	"gopkg.in/alecthomas/kingpin.v2"
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
