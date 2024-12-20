package solution

import (
	"fmt"
	"strings"
)

type Result struct {
	TotalDistance   int
	SimilarityScore int
}

func (r *Result) String() (result string) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Total distance: %d\n", r.TotalDistance))
	sb.WriteString(fmt.Sprintf("Similarity score: %d\n", r.SimilarityScore))

	return sb.String()
}
