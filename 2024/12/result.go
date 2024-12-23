package solution

import (
	"fmt"
	"strings"
)

type Result struct {
	Price int
}

func (r *Result) String() (result string) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Price: %d\n", r.Price))

	return sb.String()
}
