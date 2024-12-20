package solution

import (
	"fmt"
	"strings"
)

type Result struct {
	GuardDistinctLocations int
}

func (r *Result) String() (result string) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Number of distinct locations: %d\n", r.GuardDistinctLocations))

	return sb.String()
}
