package solution

import (
	"fmt"
	"strings"
)

type Result struct {
	SafetyFactor int
}

func (r *Result) String() (result string) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Safety factor: %d\n", r.SafetyFactor))

	return sb.String()
}
