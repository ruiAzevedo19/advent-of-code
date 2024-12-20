package solution

import (
	"fmt"
	"strings"
)

type Result struct {
	TotalCalibrationResult int
}

func (r *Result) String() (result string) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Total calibration result: %d\n", r.TotalCalibrationResult))

	return sb.String()
}
