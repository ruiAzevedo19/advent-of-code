package solution

import (
	"fmt"
	"strings"
)

type Result struct {
	Trails int
}

func (r *Result) String() (result string) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Trails: %d\n", r.Trails))

	return sb.String()
}
