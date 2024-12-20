package solution

import (
	"os"
	"strconv"
	"strings"
)

func Parse(filePath string) (reports []*Report, err error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(fileContent), "\n")

	reports = []*Report{}
	for _, line := range lines {
		reportValues := strings.Fields(line)

		report := []int{}
		for _, reportValue := range reportValues {
			reportValue := strings.TrimSpace(reportValue)
			value, err := strconv.Atoi(reportValue)
			if err != nil {
				return nil, err
			}
			report = append(report, value)
		}
		reports = append(reports, NewReport(report))
	}

	return reports, nil
}
