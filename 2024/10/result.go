package solution

import (
	"fmt"
	"strings"
)

type Result struct {
	Scores  int
	Ratings int
}

func (r *Result) String() (result string) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Scores: %d\n", r.Scores))
	sb.WriteString(fmt.Sprintf("Ratings: %d\n", r.Ratings))

	return sb.String()
}
