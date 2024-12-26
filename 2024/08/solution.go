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
				X: row,
				Y: column,
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
					X: row,
					Y: column,
				})
			}
		}
	}

	return coordinates
}

func createAntinodes(antinodes [][]rune, x *util.Coordinate, y *util.Coordinate) (numberOfAntinodeLocations int) {
	dx := x.X - y.X
	dy := x.Y - y.Y

	antinodeX := &util.Coordinate{
		X: x.X + dx,
		Y: x.Y + dy,
	}
	antinodeY := &util.Coordinate{
		X: y.X - dx,
		Y: y.Y - dy,
	}

	if antinodeX.IsValid(antinodes) && antinodes[antinodeX.X][antinodeX.Y] != ANTINODE {
		antinodes[antinodeX.X][antinodeX.Y] = ANTINODE
		numberOfAntinodeLocations++
	}
	if antinodeY.IsValid(antinodes) && antinodes[antinodeY.X][antinodeY.Y] != ANTINODE {
		antinodes[antinodeY.X][antinodeY.Y] = ANTINODE
		numberOfAntinodeLocations++
	}

	return numberOfAntinodeLocations
}

func createAntinodesUpdatedModel(antinodes [][]rune, x *util.Coordinate, y *util.Coordinate) (numberOfAntinodeLocations int) {
	dx := x.X - y.X
	dy := x.Y - y.Y

	antinodeX := &util.Coordinate{}
	for i := 1; ; i++ {
		antinodeX.X = x.X + i*dx
		antinodeX.Y = x.Y + i*dy

		if !antinodeX.IsValid(antinodes) {
			break
		}

		if antinodes[antinodeX.X][antinodeX.Y] != ANTINODE {
			antinodes[antinodeX.X][antinodeX.Y] = ANTINODE
			numberOfAntinodeLocations++
		}
	}

	antinodeY := &util.Coordinate{}
	for i := 1; ; i++ {
		antinodeY.X = y.X - i*dx
		antinodeY.Y = y.Y - i*dy

		if !antinodeY.IsValid(antinodes) {
			break
		}

		if antinodes[antinodeY.X][antinodeY.Y] != ANTINODE {
			antinodes[antinodeY.X][antinodeY.Y] = ANTINODE
			numberOfAntinodeLocations++
		}
	}

	return numberOfAntinodeLocations
}
