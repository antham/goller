package dispatcher

import (
	"github.com/antham/goller/v2/aggregator"
)

// Dispatcher interface has to be implemented to render results
type Dispatcher interface {
	RenderItems(aggregators *aggregator.Aggregators)
}
