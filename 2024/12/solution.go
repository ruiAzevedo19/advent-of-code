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

	for _, direction := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		a, p, c := group(garden, row+direction[0], column+direction[1], plant, visited)
		area, perimeter, corners = area+a, perimeter+p, corners+c
	}

	return area, perimeter, corners
}

func calculateAreaPerimeter(garden [][]rune, row int, column int) (area int, perimeter int, corners int) {
	differs := func(x, y int) bool {
		return garden[row][column] != garden[x][y]
	}

	calculatePerimeter := func(dx int, dy int) {
		if !util.RangeIsValid(garden, row+dx, column+dy) {
			perimeter++
		} else if differs(row+dx, column+dy) {
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
			if differs(ia, ja) && differs(ib, jb) {
				corners++
			} else if !differs(ia, ja) && !differs(ib, jb) && differs(ic, jc) {
				corners++
			}
		} else if !util.RangeIsValid(garden, ia, ja) && util.RangeIsValid(garden, ib, jb) {
			if differs(ib, jb) {
				corners++
			}
		} else if !util.RangeIsValid(garden, ib, jb) && util.RangeIsValid(garden, ia, ja) {
			if differs(ia, ja) {
				corners++
			}
		}
	}

	for _, direction := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		calculatePerimeter(direction[0], direction[1])
	}

	for _, direction := range [][]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}} {
		calculateCorners(direction[0], direction[1])
	}

	return 1, perimeter, corners
}
