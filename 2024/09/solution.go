package solution

import (
	"advent-of-code/challenge"
	"strconv"
)

func init() {
	challenge.Register("2024", "9", DiskFragmenter)
}

func DiskFragmenter(filePath string) (result challenge.Result, err error) {
	diskMap, err := Parse(filePath)
	if err != nil {
		return nil, err
	}

	looseFormat, err := looseFormat(diskMap)
	if err != nil {
		return nil, err
	}
	fileBlocks := moveFileBlocks(looseFormat)

	checksum, err := checksum(fileBlocks)
	if err != nil {
		return nil, err
	}

	return &Result{
		Checksum: checksum,
	}, nil
}

func looseFormat(diskMap string) (looseFormat []string, err error) {
	for i, d := range diskMap {
		blocks, err := strconv.Atoi(string(d))
		if err != nil {
			return nil, err
		}

		if i%2 == 0 {
			fileID := i / 2
			looseFormat = append(looseFormat, replicate(strconv.Itoa(fileID), blocks)...)
		} else {
			looseFormat = append(looseFormat, replicate(".", blocks)...)
		}
	}

	return looseFormat, nil
}

func moveFileBlocks(looseFormat []string) (movedFileBlock []string) {
	start, end := 0, len(looseFormat)-1

	for start < end {
		if looseFormat[start] != "." {
			start++
			continue
		}
		if looseFormat[end] == "." {
			end--
			continue
		}

		looseFormat[start], looseFormat[end] = looseFormat[end], looseFormat[start]
	}

	return looseFormat
}

func checksum(fileBlocks []string) (checksum int, err error) {
	for i, block := range fileBlocks {
		if block == "." {
			return checksum, nil
		}

		fileBlock, err := strconv.Atoi(block)
		if err != nil {
			return 0, err
		}

		checksum += i * fileBlock
	}

	return checksum, nil
}

func replicate(s string, n int) (result []string) {
	for range n {
		result = append(result, s)
	}

	return result
}
