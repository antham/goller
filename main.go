package main

import (
	"fmt"
	"github.com/antham/goller/agregator"
	"github.com/antham/goller/cli"
	"github.com/antham/goller/dispatcher"
	"github.com/antham/goller/reader"
	"github.com/antham/goller/tokenizer"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

const version = "1.4.0"

type groupCommand struct {
	cmd          *kingpin.CmdClause
	delimiter    *string
	transformers *cli.Transformers
	parser       *cli.Parser
	sorters      *cli.Sorters
	positions    *string
}

type command map[string]*kingpin.CmdClause

var (
	app = kingpin.New("goller", "Aggregate log fields and count occurences")

	cmd = map[string]*kingpin.CmdClause{
		"group": app.Command("group", "Group occurence of field"),
	}

	groupArgs = &groupCommand{
		delimiter:    cmd["group"].Flag("delimiter", "Separator between results").Short('d').Default(" | ").String(),
		transformers: cli.TransformersWrapper(cmd["group"].Flag("transformer", "Transformers applied to every fields").Short('t')),
		sorters:      cli.SortersWrapper(cmd["group"].Flag("sort", "Sort lines").Short('s')),
		parser:       cli.ParserWrapper(cmd["group"].Arg("parser", "Log line parser to use").Required()),
		positions:    cmd["group"].Arg("positions", "Field positions").Required().String(),
	}
)

// main entry point
func main() {
	app.Version(version)

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case cmd["group"].FullCommand():
		group(groupArgs)
	}
}

// group run field agregation
func group(args *groupCommand) {
	tok := tokenizer.NewTokenizer(args.parser.Get())

	agrBuilder := agregator.NewBuilder()

	var positions []int

	reader.ReadStdin(func(line string) {
		tokens, err := tok.Tokenize(line)

		if err != nil {
			fmt.Println(err)

			os.Exit(1)
		}

		if len(positions) == 0 {
			var err error
			positions, err = cli.ExtractPositions(*args.positions, len(tokens))

			if err != nil {
				fmt.Println(err)

				os.Exit(1)
			}
		}

		agrBuilder.Agregate(positions, &tokens, args.transformers.Get())
	})

	if args.sorters.Get() != nil {
		args.sorters.Get().Sort(agrBuilder.Get())
	}

	var d dispatcher.Dispatcher

	d = dispatcher.NewTermDispatch(*args.delimiter)
	d.RenderItems(agrBuilder.Get())
}
