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
func NewTermDispatcher(delimiter string) Term {
	return Term{delimiter: delimiter}
}

// RenderItems is called to render items in a terminal
func (t Term) RenderItems(agregators *agregator.Agregators) {
	for _, agregator := range *agregators {
		fmt.Println(strings.Join(agregator.Datas, t.delimiter))
	}

}
