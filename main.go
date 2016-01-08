package main

import (
	"fmt"
	"github.com/antham/goller/agregator"
	"github.com/antham/goller/dispatcher"
	"github.com/antham/goller/reader"
	"github.com/antham/goller/tokenizer"
	"github.com/antham/goller/transformer"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	app = kingpin.New("goller", "Aggregate log fields and count occurences")

	counter             = app.Command("counter", "Count occurence of field")
	counterDelimiter    = counter.Flag("delimiter", "Separator bewteen results").Short('d').Default(" | ").String()
	counterTransformers = transformer.TransformersWrapper(counter.Flag("trans", "Transformers applied to every fields").Short('t'))
	counterPositions    = counter.Arg("positions", "Field positions").Required().String()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case counter.FullCommand():
		count(*counterPositions, *counterDelimiter, *counterTransformers)
	}
}

func count(positionsString string, delimiter string, trans transformer.TransformersMap) {
	tokenizer.Init()

	agregators := agregator.NewAgregators()

	reader.ReadStdin(func(line string) {
		tokens := tokenizer.Tokenize(line)

		positions, err := agregator.ExtractPositions(positionsString, len(tokens))

		if err != nil {
			fmt.Println(err)

			os.Exit(1)
		}

		agregators.Agregate(positions, &tokens, trans)
	})

	d := dispatcher.NewTermDispatcher(delimiter)
	d.RenderItems(agregators.Get())
}
