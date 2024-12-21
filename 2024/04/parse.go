package solution

import (
	"os"
	"strings"
)

func Parse(filePath string) (grid []string, err error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.TrimSpace(string(fileContent)), "\n")

	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}

	return lines, nil
}
