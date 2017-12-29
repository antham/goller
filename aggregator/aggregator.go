package aggregator

// Aggregator represents a unique log line
type Aggregator struct {
	Count        int
	DatasOrdered map[int]*string
	Datas        []*string
}

// Aggregators contains a slice of Aggregator
type Aggregators []*Aggregator

// NewAggregators create aggregators
func NewAggregators() *Aggregators {
	return &Aggregators{}
}
