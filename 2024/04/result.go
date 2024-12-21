package solution

import (
	"fmt"
	"strings"
)

type Result struct {
	NumberOfWords  int
	NumberOfXWords int
}

func (r *Result) String() (result string) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Number of XMAS words: %d\n", r.NumberOfWords))
	sb.WriteString(fmt.Sprintf("Number of X-MAS words: %d\n", r.NumberOfXWords))

	return sb.String()
}
