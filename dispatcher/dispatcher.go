package dispatcher

import (
	"github.com/antham/goller/agregator"
)

// dispatch interface has to be implemented to render results
type dispatch interface {
	RenderItems(agregators map[[20]byte]*agregator.Agregator)
}
