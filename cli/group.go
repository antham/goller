package cli

import (
	"fmt"
	"github.com/antham/goller/v2/aggregator"
	"github.com/antham/goller/v2/dispatcher"
	"github.com/antham/goller/v2/reader"
	"github.com/antham/goller/v2/tokenizer"
)

// Group is tied to group command
type Group struct {
	tokenizer   tokenizer.Tokenizer
	agrBuilder  aggregator.Builder
	dispatcher  dispatcher.Dispatcher
	aggregators *aggregator.Aggregators
	reader      reader.Reader
	ignore      *bool
	positions   *[]int
	args        *groupCommand
}

// NewGroup creates an object related to group command
func NewGroup(args *groupCommand) *Group {
	return &Group{
		tokenizer:  *tokenizer.NewTokenizer(args.parser.Get()),
		agrBuilder: *aggregator.NewBuilder(),
		dispatcher: dispatcher.NewTermDispatch(*args.delimiter),
		reader:     reader.NewStdinReader(),
		positions:  args.positions.Get(),
		ignore:     args.ignore,
		args:       args,
	}
}

// Consume tokenizes every line from reader
func (g *Group) Consume() {
	err := g.reader.Read(func(line []byte) error {
		err := g.tokenizer.Tokenize(line)

		size := g.tokenizer.Length()

		if err != nil && *g.ignore {
			return nil
		} else if err != nil {
			return err
		} else if size == 0 {
			err = fmt.Errorf("No tokens found")
		} else if positionsOutOfBoundary(g.positions, size) {
			err = fmt.Errorf("A position is greater or equal than maximum position %d", size)
		}

		if err != nil {
			return err
		}

		g.agrBuilder.Aggregate(*g.positions, g.tokenizer.Get(), g.args.transformers.Get())

		return nil
	})

	if err != nil {
		triggerFatalError(err)
	}

	g.agrBuilder.SetCounterIfAny()
	g.aggregators = g.agrBuilder.Get()
}

// Sort orders tokenized lines
func (g *Group) Sort() {
	if g.args.sorters.Get() != nil {
		g.args.sorters.Get().Sort(g.aggregators)
	}
}

// Dispatch calls appropriate renderer
func (g *Group) Dispatch() {
	g.dispatcher.RenderItems(g.aggregators)
}
