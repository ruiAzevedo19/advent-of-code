package solution

import (
	"advent-of-code/2024/02/parse"
	"advent-of-code/challenge"
)

func init() {
	challenge.Register("2024", "2", RedNosedReports)
}

func RedNosedReports(filePath string) (result challenge.Result, err error) {
	reports, err := parse.Parse(filePath)
	if err != nil {
		return nil, err
	}

	var numberOfCorrectReports int
	var numberOfFixedReports int
	for _, report := range reports {
		if report.IsSafe() {
			numberOfCorrectReports++
		} else if report.CanBeSafe() {
			numberOfFixedReports++
		}
	}

	return &Result{
		NumberOfCorrectReports: numberOfCorrectReports,
		NumberOfFixedReports:   numberOfFixedReports,
	}, nil
}
