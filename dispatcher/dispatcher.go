package dispatcher

import (
	"github.com/antham/goller/aggregator"
)

// Dispatcher interface has to be implemented to render results
type Dispatcher interface {
	RenderItems(aggregators *aggregator.Aggregators)
}
