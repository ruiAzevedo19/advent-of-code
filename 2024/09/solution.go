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
	checksum, err := calculateChecksum(fileBlocks)
	if err != nil {
		return nil, err
	}

	compactedFileBlocks := compactFiles(looseFormat)
	checksumCompacted, err := calculateChecksum(compactedFileBlocks)
	if err != nil {
		return nil, err
	}

	return &Result{
		Checksum:          checksum,
		ChecksumCompacted: checksumCompacted,
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
	movedFileBlock = make([]string, len(looseFormat))
	copy(movedFileBlock, looseFormat)

	start, end := 0, len(looseFormat)-1

	for start < end {
		if movedFileBlock[start] != "." {
			start++
			continue
		}
		if movedFileBlock[end] == "." {
			end--
			continue
		}

		movedFileBlock[start], movedFileBlock[end] = movedFileBlock[end], movedFileBlock[start]
	}

	return movedFileBlock
}

func compactFiles(looseFormat []string) (compactedFiles []string) {
	compactedFiles = make([]string, len(looseFormat))
	copy(compactedFiles, looseFormat)

	startFreeSpace, endFreeSpace := 0, 0
	startFileBlock, endFileBlock := len(compactedFiles)-1, len(compactedFiles)-1

	for startFileBlock > 0 {
		// Start of free space.
		if compactedFiles[startFreeSpace] != "." {
			startFreeSpace++
			continue
		}
		// End of free space.
		if endFreeSpace < len(compactedFiles)-1 && (compactedFiles[endFreeSpace] != "." || compactedFiles[endFreeSpace+1] == ".") {
			endFreeSpace++
			continue
		}
		// End of file block.
		if compactedFiles[endFileBlock] == "." {
			endFileBlock--
			continue
		}
		// Start of file block.
		if compactedFiles[startFileBlock-1] == compactedFiles[endFileBlock] {
			startFileBlock--
			continue
		}

		// If all free spaces are checked move to the next file block and start from the beggining.
		if startFreeSpace > startFileBlock {
			startFileBlock, endFileBlock = startFileBlock-1, startFileBlock-1
			startFreeSpace, endFreeSpace = 0, 0
			continue
		}

		fileBlockSpace := endFileBlock - startFileBlock + 1
		freeSpace := endFreeSpace - startFreeSpace + 1

		// If there is no free space for the file blocks find the next free space.
		if fileBlockSpace > freeSpace {
			startFreeSpace, endFreeSpace = endFreeSpace+1, endFreeSpace+1
			continue
		}

		// If there is free space for the file blocks copy them.
		for i := 0; i < fileBlockSpace; i++ {
			compactedFiles[startFreeSpace+i], compactedFiles[startFileBlock+i] = compactedFiles[startFileBlock+i], compactedFiles[startFreeSpace+i]
		}
		startFileBlock, endFileBlock = startFileBlock-1, startFileBlock-1
		startFreeSpace, endFreeSpace = 0, 0
	}

	return compactedFiles
}

func calculateChecksum(fileBlocks []string) (checksum int, err error) {
	for i, block := range fileBlocks {
		if block == "." {
			continue
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
