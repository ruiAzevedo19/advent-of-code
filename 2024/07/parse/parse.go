package parse

import (
	"advent-of-code/2024/07/model"
	"errors"
	"os"
	"strconv"
	"strings"
)

var ErrInvalidFormat = errors.New("invalid format")

func Parse(filePath string) (equations []*model.Equation, err error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	content := strings.TrimSpace(string(fileContent))
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		l := strings.Split(line, ":")
		if len(l) != 2 {
			return nil, ErrInvalidFormat
		}

		testValue, err := strconv.Atoi(strings.TrimSpace(l[0]))
		if err != nil {
			return nil, err
		}

		var numbers []int
		parsedNumbers := strings.Split(strings.TrimSpace(l[1]), " ")
		for _, number := range parsedNumbers {
			number = strings.TrimSpace(number)
			n, err := strconv.Atoi(number)
			if err != nil {
				return nil, err
			}
			numbers = append(numbers, n)
		}

		equations = append(equations, model.NewEquation(testValue, numbers))
	}

	return equations, nil
}
