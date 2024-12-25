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

	price, priceDiscount := calculatePrice(garden)

	return &Result{
		Price:         price,
		PriceDiscount: priceDiscount,
	}, nil
}

func calculatePrice(garden [][]rune) (price int, priceDiscount int) {
	visited := util.EmptyGrid[bool](len(garden), len(garden[0]))

	for i := range garden {
		for j := range garden[i] {
			if visited[i][j] {
				continue
			}
			area, perimeter, corners := group(garden, i, j, garden[i][j], visited)
			price += area * perimeter
			priceDiscount += area * corners
		}
	}

	return price, priceDiscount
}

func group(garden [][]rune, row int, column int, plant rune, visited [][]bool) (area int, perimeter int, corners int) {
	if !util.RangeIsValid(garden, row, column) {
		return 0, 0, 0
	}
	if visited[row][column] || garden[row][column] != plant {
		return 0, 0, 0
	}

	visited[row][column] = true
	area, perimeter, corners = calculateAreaPerimeter(garden, row, column)

	areaUp, perimeterUp, cornersUp := group(garden, row-1, column, plant, visited)
	areaRight, perimeterRight, cornersRight := group(garden, row, column+1, plant, visited)
	areaDown, perimeterDown, cornersDown := group(garden, row+1, column, plant, visited)
	areaLeft, perimeterLeft, cornersLeft := group(garden, row, column-1, plant, visited)

	area += areaUp + areaRight + areaDown + areaLeft
	perimeter += perimeterUp + perimeterRight + perimeterDown + perimeterLeft
	corners += cornersUp + cornersRight + cornersDown + cornersLeft

	return area, perimeter, corners
}

func calculateAreaPerimeter(garden [][]rune, row int, column int) (area int, perimeter int, corners int) {
	calculatePerimeter := func(dx int, dy int) {
		// Calculate perimenter.
		if !util.RangeIsValid(garden, row+dx, column+dy) {
			perimeter++
		} else if garden[row][column] != garden[row+dx][column+dy] {
			perimeter++
		}
	}
	calculateCorners := func(dx int, dy int) {
		ia, ja := row+dx, column
		ib, jb := row, column+dy
		ic, jc := row+dx, column+dy

		if !util.RangeIsValid(garden, ia, ja) && !util.RangeIsValid(garden, ib, jb) {
			corners++

			return
		}

		if util.RangeIsValid(garden, ia, ja) && util.RangeIsValid(garden, ib, jb) && util.RangeIsValid(garden, ic, jc) {
			if garden[row][column] != garden[ia][ja] && garden[row][column] != garden[ib][jb] {
				corners++
			} else if garden[row][column] == garden[ia][ja] && garden[row][column] == garden[ib][jb] && garden[row][column] != garden[ic][jc] {
				corners++
			}
		} else if !util.RangeIsValid(garden, ia, ja) && util.RangeIsValid(garden, ib, jb) {
			if garden[row][column] != garden[ib][jb] {
				corners++
			}
		} else if !util.RangeIsValid(garden, ib, jb) && util.RangeIsValid(garden, ia, ja) {
			if garden[row][column] != garden[ia][ja] {
				corners++
			}
		}
	}

	// Perimeter up.
	calculatePerimeter(-1, 0)
	// Perimeter right.
	calculatePerimeter(0, 1)
	// Perimeter down.
	calculatePerimeter(1, 0)
	// Perimeter left.
	calculatePerimeter(0, -1)

	// Corner left-up.
	calculateCorners(-1, -1)
	// Corner up-right.
	calculateCorners(-1, 1)
	// Corner right-down.
	calculateCorners(1, 1)
	// Corner down-left.
	calculateCorners(1, -1)

	return 1, perimeter, corners
}
