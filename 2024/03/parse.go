package solution

import "os"

func Parse(filePath string) (corruptedMemory string, err error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return "", nil
	}

	return string(fileContent), nil
}
