package solution

import (
	"fmt"
	"strings"
)

type Result struct {
	Price         int
	PriceDiscount int
}

func (r *Result) String() (result string) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Price: %d\n", r.Price))
	sb.WriteString(fmt.Sprintf("Price with discount: %d\n", r.PriceDiscount))

	return sb.String()
}
