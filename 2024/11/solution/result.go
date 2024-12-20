package solution

import (
	"fmt"
	"strings"
)

type Result struct {
	NumberOfStones25 int
	NumberOfStones75 int
}

func (r *Result) String() (result string) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Number of stones (25 blinks): %d\n", r.NumberOfStones25))
	sb.WriteString(fmt.Sprintf("Number of stones (75 blinks): %d\n", r.NumberOfStones75))

	return sb.String()
}
