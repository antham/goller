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
	positions    *string
}

// command list all available cli commands
type command map[string]*kingpin.CmdClause

var (
	app = kingpin.New("goller", "Aggregate log fields and count occurences")

	cmd = map[string]*kingpin.CmdClause{
		"group": app.Command("group", "Group occurence of field"),
	}

	groupArgs = &groupCommand{
		delimiter:    cmd["group"].Flag("delimiter", "Separator between results").Short('d').Default(" | ").String(),
		transformers: TransformersWrapper(cmd["group"].Flag("transformer", "Transformers applied to every fields").Short('t')),
		sorters:      SortersWrapper(cmd["group"].Flag("sort", "Sort lines").Short('s')),
		parser:       ParserWrapper(cmd["group"].Arg("parser", "Log line parser to use").Required()),
		positions:    cmd["group"].Arg("positions", "Field positions").Required().String(),
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
