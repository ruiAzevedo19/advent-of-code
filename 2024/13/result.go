package solution

import (
	"fmt"
	"strings"
)

type Result struct {
	Tokens int
}

func (r *Result) String() (result string) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Tokens: %d\n", r.Tokens))

	return sb.String()
}
