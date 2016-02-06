package cli

import (
	"fmt"
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
	ignore     *bool
	positions  *[]int
	args       *groupCommand
}

// NewGroup create an object related to group command
func NewGroup(args *groupCommand) *group {
	return &group{
		tokenizer:  *tokenizer.NewTokenizer(args.parser.Get()),
		agrBuilder: *agregator.NewBuilder(),
		dispatcher: dispatcher.NewTermDispatch(*args.delimiter),
		reader:     reader.NewStdinReader(),
		positions:  args.positions.Get(),
		ignore:     args.ignore,
		args:       args,
	}
}

// Consume tokenize every line from reader
func (g *group) Consume() {
	g.reader.Read(func(line string) {
		tokens, err := g.tokenizer.Tokenize(line)

		if err != nil && *g.ignore {
			return
		}

		triggerFatalError(err)

		if positionsOutOfBoundary(g.positions, len(tokens)) {
			triggerFatalError(fmt.Errorf("A position is greater or equal than maximum position %d", len(tokens)))
		}

		g.agrBuilder.Agregate(*g.positions, &tokens, g.args.transformers.Get())
	})

	g.agrBuilder.SetCounterIfAny()
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
