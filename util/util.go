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

func EmptyGrid(rows int, columns int) (grid [][]rune) {
	grid = make([][]rune, rows)
	for row := range grid {
		grid[row] = make([]rune, columns)
	}

	return grid
}

func PrintGrid(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println()
}
