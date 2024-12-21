package solution

import (
	"fmt"
	"strings"
)

type Result struct {
	NumberOfCorrectReports int
	NumberOfFixedReports   int
}

func (r *Result) String() (result string) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Numer of correct reports: %d\n", r.NumberOfCorrectReports))
	sb.WriteString(fmt.Sprintf("Numer of fixed reports: %d\n", r.NumberOfFixedReports))
	sb.WriteString(fmt.Sprintf("Total: %d\n", r.NumberOfCorrectReports+r.NumberOfFixedReports))

	return sb.String()
}
