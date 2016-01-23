package dispatcher

import (
	"github.com/antham/goller/agregator"
)

// Dispatcher interface has to be implemented to render results
type Dispatcher interface {
	RenderItems(agregators *agregator.Agregators)
}

