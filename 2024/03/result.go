package solution

import (
	"fmt"
	"strings"
)

type Result struct {
	SummedMultiplications   int
	SummedMultiplicationsDo int
}

func (r *Result) String() (result string) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Summed multiplications: %d\n", r.SummedMultiplications))
	sb.WriteString(fmt.Sprintf("Summed multiplication with do-instructions: %d\n", r.SummedMultiplicationsDo))

	return sb.String()
}
