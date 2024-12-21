package util

import (
	"os"
	"strings"
)

func ParseGrid(filePath string) (grid [][]rune, err error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	content := strings.TrimSpace(string(fileContent))

	lines := strings.Split(content, "\n")
	for _, line := range lines {
		l := []rune{}
		for _, r := range line {
			l = append(l, r)
		}
		grid = append(grid, l)
	}

	return grid, nil
}
