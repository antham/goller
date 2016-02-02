package dispatcher

import (
	"fmt"
	"github.com/antham/goller/agregator"
	"strings"
)

// Term defined a terminal renderer
type Term struct {
	delimiter string
}

// NewTermDispatcher create a new terminal renderer
func NewTermDispatch(delimiter string) Term {
	return Term{delimiter: delimiter}
}

// RenderItems is called to render items in a terminal
func (t Term) RenderItems(agregators *agregator.Agregators) {
	for _, agregator := range *agregators {
		datas := []string{}

		for _, data := range agregator.Datas {
			datas = append(datas, *data)
		}

		fmt.Println(strings.Join(datas, t.delimiter))
	}
}
