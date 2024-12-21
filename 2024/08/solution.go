package solution

import (
	"advent-of-code/challenge"
	"advent-of-code/util"
)

const (
	ANTINODE = '#'
	EMPTY    = '.'
)

func init() {
	challenge.Register("2024", "8", ResonantCollinearity)
}

func ResonantCollinearity(filePath string) (result challenge.Result, err error) {
	grid, err := util.ParseGrid(filePath)
	if err != nil {
		return nil, err
	}

	numberOfAntinodeLocations := resonantCollinearity(grid)

	return &Result{
		NumberOfAntinodeLocations: numberOfAntinodeLocations,
	}, nil
}

func resonantCollinearity(grid [][]rune) (numberOfAntinodeLocations int) {
	antinodes := util.EmptyGrid(len(grid), len(grid[0]))

	for row := range grid {
		for column := range grid[row] {
			if grid[row][column] == EMPTY {
				continue
			}

			currentCoordinate := &util.Coordinate{
				Row:    row,
				Column: column,
			}

			coordinates := findFrequencies(grid, grid[row][column])
			for _, coordinate := range coordinates {
				if currentCoordinate.Equals(coordinate) {
					continue
				}
				numberOfAntinodeLocations += createAntinodes(antinodes, currentCoordinate, coordinate)
			}
		}
	}

	return numberOfAntinodeLocations
}

func findFrequencies(grid [][]rune, frequency rune) (coordinates []*util.Coordinate) {
	for row := range grid {
		for column := range grid[row] {
			if grid[row][column] == frequency {
				coordinates = append(coordinates, &util.Coordinate{
					Row:    row,
					Column: column,
				})
			}
		}
	}

	return coordinates
}

func createAntinodes(antinodes [][]rune, x *util.Coordinate, y *util.Coordinate) (numberOfAntinodeLocations int) {
	dx := x.Row - y.Row
	dy := x.Column - y.Column

	antinodeX := &util.Coordinate{
		Row:    x.Row + dx,
		Column: x.Column + dy,
	}
	antinodeY := &util.Coordinate{
		Row:    y.Row - dx,
		Column: y.Column - dy,
	}

	if antinodeX.IsValid(antinodes) && antinodes[antinodeX.Row][antinodeX.Column] != ANTINODE {
		antinodes[antinodeX.Row][antinodeX.Column] = ANTINODE
		numberOfAntinodeLocations++
	}
	if antinodeY.IsValid(antinodes) && antinodes[antinodeY.Row][antinodeY.Column] != ANTINODE {
		antinodes[antinodeY.Row][antinodeY.Column] = ANTINODE
		numberOfAntinodeLocations++
	}

	return numberOfAntinodeLocations
}
