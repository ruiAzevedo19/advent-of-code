package util

import "fmt"

type Coordinate struct {
	Row    int
	Column int
}

func (c *Coordinate) Equals(coordinate *Coordinate) bool {
	return c.Row == coordinate.Row && c.Column == coordinate.Column
}

func (c *Coordinate) IsValid(grid [][]rune) bool {
	if len(grid) == 0 {
		return false
	}

	if c.Row < 0 || c.Row >= len(grid) {
		return false
	}
	if c.Column < 0 || c.Column >= len(grid[0]) {
		return false
	}

	return true
}

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
