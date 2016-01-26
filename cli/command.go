package cli

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

// groupCommand describe all dependencies of a group command
type groupCommand struct {
	cmd          *kingpin.CmdClause
	delimiter    *string
	transformers *Transformers
	parser       *Parser
	sorters      *Sorters
	positions    *[]int
}

// command list all available cli commands
type command map[string]*kingpin.CmdClause

var (
	app = kingpin.New("goller", "Aggregate log fields and count occurences")

	cmd = map[string]*kingpin.CmdClause{
		"group": app.Command("group", "Group occurence of field"),
	}

	parserArg    = ParserWrapper(cmd["group"].Arg("parser", "Log line parser to use").Required())
	positionsArg = PositionsWrapper(cmd["group"].Arg("positions", "Field positions").Required()).Get()

	groupArgs = &groupCommand{
		delimiter:    cmd["group"].Flag("delimiter", "Separator between results").Short('d').Default(" | ").String(),
		transformers: TransformersWrapper(cmd["group"].Flag("transformer", "Transformers applied to every fields").Short('t'), positionsArg),
		sorters:      SortersWrapper(cmd["group"].Flag("sort", "Sort lines").Short('s'), positionsArg),
		parser:       parserArg,
		positions:    positionsArg,
	}
)

// Run commmand line arguments parsing
func Run(version string) {
	app.Version(version)

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case cmd["group"].FullCommand():
		group := NewGroup(groupArgs)
		group.Consume()
		group.Sort()
		group.Dispatch()
	}
}
