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

	antinodeLocations, antinodeLocationsUpdatedModel := resonantCollinearity(grid)

	return &Result{
		NumberOfAntinodeLocations:             antinodeLocations,
		NumberOfAntinodeLocationsUpdatedModel: antinodeLocationsUpdatedModel,
	}, nil
}

func resonantCollinearity(grid [][]rune) (antinodeLocations int, antinodeLocationsUpdatedModel int) {
	antinodes := util.EmptyGrid[rune](len(grid), len(grid[0]))
	antinodesUpdatedModel := util.EmptyGrid[rune](len(grid), len(grid[0]))

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
					if antinodesUpdatedModel[row][column] != ANTINODE {
						antinodesUpdatedModel[row][column] = ANTINODE
						antinodeLocationsUpdatedModel++
					}

					continue
				}
				antinodeLocations += createAntinodes(antinodes, currentCoordinate, coordinate)
				antinodeLocationsUpdatedModel += createAntinodesUpdatedModel(antinodesUpdatedModel, currentCoordinate, coordinate)
			}
		}
	}

	return antinodeLocations, antinodeLocationsUpdatedModel
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

func createAntinodesUpdatedModel(antinodes [][]rune, x *util.Coordinate, y *util.Coordinate) (numberOfAntinodeLocations int) {
	dx := x.Row - y.Row
	dy := x.Column - y.Column

	antinodeX := &util.Coordinate{}
	for i := 1; ; i++ {
		antinodeX.Row = x.Row + i*dx
		antinodeX.Column = x.Column + i*dy

		if !antinodeX.IsValid(antinodes) {
			break
		}

		if antinodes[antinodeX.Row][antinodeX.Column] != ANTINODE {
			antinodes[antinodeX.Row][antinodeX.Column] = ANTINODE
			numberOfAntinodeLocations++
		}
	}

	antinodeY := &util.Coordinate{}
	for i := 1; ; i++ {
		antinodeY.Row = y.Row - i*dx
		antinodeY.Column = y.Column - i*dy

		if !antinodeY.IsValid(antinodes) {
			break
		}

		if antinodes[antinodeY.Row][antinodeY.Column] != ANTINODE {
			antinodes[antinodeY.Row][antinodeY.Column] = ANTINODE
			numberOfAntinodeLocations++
		}
	}

	return numberOfAntinodeLocations
}
