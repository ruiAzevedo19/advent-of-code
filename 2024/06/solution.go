package solution

import (
	"advent-of-code/challenge"
	"fmt"
)

func init() {
	challenge.Register("2024", "6", GuardGallivant)
}

const (
	BLOCK       = 'O'
	GUARD       = '^'
	OBSTRUCTION = '#'
	VISITED     = 'X'
)

// Let's assume the grid is not empty and has at least one row.
func GuardGallivant(filePath string) (result challenge.Result, err error) {
	grid, err := Parse(filePath)
	if err != nil {
		return nil, err
	}

	guardDistinctLocations, _ := guardDistinctLocations(newGrid(grid))

	return &Result{
		GuardDistinctLocations: guardDistinctLocations,
	}, nil
}

func guardDistinctLocations(grid [][]rune) (distinctLocations int, isBlocked bool) {
	row, column := findGuardInitialLocation(grid)
	if row == -1 && column == -1 {
		return -1, false
	}

	lastRow, lastColumn := row, column
	directionX, directionY := -1, 0
	locations := 0
	isOut, moved := false, false

	for !isOut && !isBlocked {
		lastRow, lastColumn, locations, moved, isOut = moveGuard(grid, lastRow, lastColumn, directionX, directionY)
		distinctLocations += locations

		if locations == 0 && moved {
			isBlocked = true
		}
		directionX, directionY = turnRight(directionX, directionY)
	}

	return distinctLocations, isBlocked
}

func moveGuard(grid [][]rune, row int, column int, directionX int, directionY int) (lastRow int, lastColumn int, locations int, moved bool, isOut bool) {
	rows := len(grid)
	columns := len(grid[0])

	var n int
	if directionX > 0 {
		n = rows - row
	} else if directionX < 0 {
		n = row + 1
	} else if directionY > 0 {
		n = columns - column
	} else {
		n = column + 1
	}

	for dx := 0; dx < n; dx++ {
		lastRow, lastColumn = row+dx*directionX, column+dx*directionY
		if grid[lastRow][lastColumn] == OBSTRUCTION || grid[lastRow][lastColumn] == BLOCK {
			return lastRow - directionX, lastColumn - directionY, locations, moved, false
		}
		if grid[lastRow][lastColumn] != VISITED {
			grid[lastRow][lastColumn] = VISITED
			locations++
		}
		moved = lastRow != row || lastColumn != column
	}

	return lastRow, lastColumn, locations, moved, true
}

func findGuardInitialLocation(grid [][]rune) (row int, column int) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == GUARD {
				return i, j
			}
		}
	}

	return -1, -1
}

func turnRight(directionX int, directionY int) (newDirectionX int, newDirectionY int) {
	if directionX == -1 && directionY == 0 {
		return 0, 1
	}
	if directionX == 0 && directionY == 1 {
		return 1, 0
	}
	if directionX == 1 && directionY == 0 {
		return 0, -1
	}

	return -1, 0
}

func newGrid(grid [][]rune) (newGrid [][]rune) {
	for i := range grid {
		row := []rune{}
		row = append(row, grid[i]...)
		newGrid = append(newGrid, row)
	}

	return newGrid
}

func printGrid(grid [][]rune) {
	for i := range grid {
		fmt.Println(string(grid[i]))
	}
	fmt.Println()
}
