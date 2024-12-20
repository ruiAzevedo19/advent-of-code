package solution

import (
	"fmt"
	"strings"
)

type Result struct {
	ValidUpdates              [][]int
	InvalidUpdates            [][]int
	MiddlePageSumValidUpdates int

	FixedUpdates              [][]int
	MiddlePageSumFixedUpdates int
}

func NewResult() (result *Result) {
	return &Result{}
}

func (r *Result) String() (result string) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("middle page sum: %d\n", r.MiddlePageSumValidUpdates))
	sb.WriteString(fmt.Sprintf("middle page sum fixed: %d\n", r.MiddlePageSumFixedUpdates))

	return sb.String()
}
