package solution

import (
	"ceres-search/parse"
)

// Let's assume the grid is not empty and always has at least one row.
func CeresSearch(filePath string) (result *Result, err error) {
	grid, err := parse.Parse(filePath)
	if err != nil {
		return nil, err
	}

	return &Result{
		NumberOfWords:  findWord(grid, "XMAS"),
		NumberOfXWords: findXWord(grid, "MAS"),
	}, nil
}

func findWord(grid []string, word string) (numberOfWords int) {
	signCombinations := [][]int{
		// Diagonal
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
		// Horizontal
		{0, 1}, {0, -1},
		// Vertical
		{1, 0}, {-1, -0},
	}

	for row := range grid {
		for column := range grid[row] {
			for _, sc := range signCombinations {
				if wordExists(grid, word, row, column, sc[0], sc[1]) {
					numberOfWords++
				}
			}
		}
	}

	return numberOfWords
}

func findXWord(grid []string, word string) (numberOfWords int) {
	rows := len(grid)
	columns := len(grid[0])

	wordReversed := reverse(word)
	wordCombinations := [][]string{
		{word, word},
		{word, wordReversed},
		{wordReversed, word},
		{wordReversed, wordReversed},
	}

	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			for _, wc := range wordCombinations {
				if wordExists(grid, wc[0], row, column, 1, 1) && wordExists(grid, wc[1], row, column+len(word)-1, 1, -1) {
					numberOfWords++
				}
			}
		}
	}

	return numberOfWords
}

func reverse(word string) (reversedWord string) {
	reversed := []rune(word)

	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}

	return string(reversed)
}

func wordExists(grid []string, word string, row int, column int, rowSign int, columnSign int) (found bool) {
	for dx, w := range word {
		i, j := row+rowSign*dx, column+columnSign*dx
		if !rangeIsValid(grid, i, j) {
			return false
		}
		if grid[i][j] != byte(w) {
			return false
		}
	}

	return true
}

func rangeIsValid(grid []string, row int, column int) (ok bool) {
	rows := len(grid)
	columns := len(grid[0])

	return row >= 0 && row < rows && column >= 0 && column < columns
}
