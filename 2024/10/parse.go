package solution

import (
	"advent-of-code/util"
	"strconv"
)

func Parse(filePath string) (grid [][]int, err error) {
	grideRune, err := util.ParseGrid(filePath)
	if err != nil {
		return nil, err
	}

	grid = make([][]int, len(grideRune))
	for i, row := range grideRune {
		grid[i] = make([]int, len(row))
		for j, v := range row {
			value, err := strconv.Atoi(string(v))
			if err != nil {
				return nil, err
			}
			grid[i][j] = value
		}
	}

	return grid, nil
}
