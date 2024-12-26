package solution

import (
	"fmt"
	"strings"
)

type Result struct {
	Tokens            int
	TokensHigherPrize int
}

func (r *Result) String() (result string) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Tokens: %d\n", r.Tokens))
	sb.WriteString(fmt.Sprintf("Tokens higher prize: %d\n", r.TokensHigherPrize))

	return sb.String()
}
