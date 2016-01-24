package cli

import (
	"github.com/antham/goller/agregator"
	"github.com/antham/goller/dispatcher"
	"github.com/antham/goller/reader"
	"github.com/antham/goller/tokenizer"
)

// group is tied to group command
type group struct {
	tokenizer  tokenizer.Tokenizer
	agrBuilder agregator.Builder
	dispatcher dispatcher.Dispatcher
	agregators *agregator.Agregators
	reader     reader.Reader
	positions  []int
	args       *groupCommand
}

// NewGroup create an object related to group command
func NewGroup(args *groupCommand) *group {
	return &group{
		tokenizer:  *tokenizer.NewTokenizer(args.parser.Get()),
		agrBuilder: *agregator.NewBuilder(),
		dispatcher: dispatcher.NewTermDispatch(*args.delimiter),
		reader:     reader.NewStdinReader(),
		args:       args,
	}
}

// Consume tokenize every line from reader
func (g *group) Consume() {
	g.reader.Read(func(line string) {
		tokens, err := g.tokenizer.Tokenize(line)

		checkFatalError(err)

		if len(g.positions) == 0 {
			var err error
			g.positions, err = extractPositions(*g.args.positions, len(tokens))

			checkFatalError(err)
		}

		g.agrBuilder.Agregate(g.positions, &tokens, g.args.transformers.Get())
	})

	g.agregators = g.agrBuilder.Get()
}

// Sort, if defined, order tokenized lines
func (g *group) Sort() {
	if g.args.sorters.Get() != nil {
		g.args.sorters.Get().Sort(g.agregators)
	}
}

// Dispatch call appropriate renderer
func (g *group) Dispatch() {
	g.dispatcher.RenderItems(g.agregators)
}
