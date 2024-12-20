package solution

import (
	"advent-of-code/2024/02/parse"
)

func RedNosedReports(filePath string) (result *Result, err error) {
	reports, err := parse.Parse(filePath)
	if err != nil {
		return nil, err
	}

	result = &Result{}
	for _, report := range reports {
		if report.IsSafe() {
			result.NumberOfCorrectReports++
		} else if report.CanBeSafe() {
			result.NumberOfFixedReports++
		}
	}

	return result, nil
}
