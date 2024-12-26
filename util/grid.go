package util

import "fmt"

func EmptyGrid[T any](rows int, columns int) (grid [][]T) {
	grid = make([][]T, rows)
	for row := range grid {
		grid[row] = make([]T, columns)
	}

	return grid
}

func RangeIsValid[T any](grid [][]T, row int, column int) (ok bool) {
	return row >= 0 && row < len(grid) && column >= 0 && column < len(grid[0])
}

func PrintGrid[T any](grid [][]T) {
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println()
}
