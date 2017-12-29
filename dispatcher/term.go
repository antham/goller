package dispatcher

import (
	"fmt"
	"github.com/antham/goller/aggregator"
	"strings"
)

// Term defined a terminal renderer
type Term struct {
	delimiter string
}

// NewTermDispatch creates a new terminal renderer
func NewTermDispatch(delimiter string) Term {
	return Term{delimiter: delimiter}
}

// RenderItems is called to render items in a terminal
func (t Term) RenderItems(aggregators *aggregator.Aggregators) {
	for _, aggregator := range *aggregators {
		datas := []string{}

		for _, data := range aggregator.Datas {
			datas = append(datas, *data)
		}

		fmt.Println(strings.Join(datas, t.delimiter))
	}
}
