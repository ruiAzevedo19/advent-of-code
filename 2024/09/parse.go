package solution

import (
	"os"
	"strings"
)

func Parse(filePath string) (diskMap string, err error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(fileContent)), nil
}
