package agregator

// Agregator represents a unique log line
type Agregator struct {
	Count        int
	DatasOrdered map[int]*string
	Datas        []*string
}

// Agregators contains a slice of Agregator
type Agregators []*Agregator

// NewAgregators create agregators
func NewAgregators() *Agregators {
	return &Agregators{}
}
