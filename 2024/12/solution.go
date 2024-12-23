package solution

import (
	"advent-of-code/challenge"
	"advent-of-code/util"
)

func init() {
	challenge.Register("2024", "12", GardenGroups)
}

func GardenGroups(filePath string) (result challenge.Result, err error) {
	garden, err := util.ParseGrid(filePath)
	if err != nil {
		return nil, err
	}

	return &Result{
		Price: calculatePrice(garden),
	}, nil
}

func calculatePrice(garden [][]rune) (price int) {
	visited := util.EmptyGrid[bool](len(garden), len(garden[0]))

	for i := range garden {
		for j := range garden[i] {
			if visited[i][j] {
				continue
			}
			area, perimeter := group(garden, i, j, garden[i][j], visited)
			price += area * perimeter
		}
	}

	return price
}

func group(garden [][]rune, row int, column int, plant rune, visited [][]bool) (area int, perimeter int) {
	if !util.RangeIsValid(garden, row, column) {
		return 0, 0
	}
	if visited[row][column] || garden[row][column] != plant {
		return 0, 0
	}

	visited[row][column] = true
	area, perimeter = calculateAreaPerimeter(garden, row, column)

	areaUp, perimeterUp := group(garden, row-1, column, plant, visited)
	areaRight, perimeterRight := group(garden, row, column+1, plant, visited)
	areaDown, perimeterDown := group(garden, row+1, column, plant, visited)
	areaLeft, perimeterLeft := group(garden, row, column-1, plant, visited)

	area += areaUp + areaRight + areaDown + areaLeft
	perimeter += perimeterUp + perimeterRight + perimeterDown + perimeterLeft

	return area, perimeter
}

func calculateAreaPerimeter(garden [][]rune, row int, column int) (area int, perimeter int) {
	calculatePerimeter := func(dx int, dy int) {
		if !util.RangeIsValid(garden, row+dx, column+dy) {
			perimeter++
		} else if garden[row][column] != garden[row+dx][column+dy] {
			perimeter++
		}
	}

	// Perimeter up.
	calculatePerimeter(-1, 0)
	// Perimeter right.
	calculatePerimeter(0, 1)
	// Perimeter down.
	calculatePerimeter(1, 0)
	// Perimeter right.
	calculatePerimeter(0, -1)

	return 1, perimeter
}
