package parse

import (
	"os"
	"strconv"
	"strings"
)

func Parse(filePath string) (stones []int, err error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	content := strings.TrimSpace(string(fileContent))

	parsedNumbers := strings.Split(content, " ")
	for _, parsedNumber := range parsedNumbers {
		stone, err := strconv.Atoi(strings.TrimSpace(parsedNumber))
		if err != nil {
			return nil, err
		}
		stones = append(stones, stone)
	}

	return stones, nil
}
