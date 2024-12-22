package solution

import (
	"advent-of-code/challenge"
	"advent-of-code/util"
)

func init() {
	challenge.Register("2024", "10", HoofIt)
}

func HoofIt(filePath string) (result challenge.Result, err error) {
	grid, err := Parse(filePath)
	if err != nil {
		return nil, err
	}

	path := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	scores, ratings := hoofIt(grid, path)

	return &Result{
		Scores:  scores,
		Ratings: ratings,
	}, nil
}

func hoofIt(grid [][]int, path []int) (scores, ratings int) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != 0 {
				continue
			}

			visited := util.EmptyGrid[bool](len(grid), len(grid[0]))
			score, rating := findPath(visited, grid, i, j, path)
			scores += score
			ratings += rating
		}
	}

	return scores, ratings
}

func findPath(visited [][]bool, grid [][]int, row int, column int, path []int) (scores, ratings int) {
	if row < 0 || row >= len(grid) || column < 0 || column >= len(grid[0]) {
		return 0, 0
	} else if grid[row][column] != path[0] {
		return 0, 0
	}

	if len(path) == 1 {
		if !visited[row][column] {
			visited[row][column] = true
			scores++
		}

		return scores, 1
	}

	scoresUp, ratingsUp := findPath(visited, grid, row-1, column, path[1:])
	scoresDown, ratingsDown := findPath(visited, grid, row+1, column, path[1:])
	scoresLeft, ratingsLeft := findPath(visited, grid, row, column-1, path[1:])
	scoresRight, ratingsRight := findPath(visited, grid, row, column+1, path[1:])

	scores = scoresUp + scoresDown + scoresLeft + scoresRight
	ratings = ratingsUp + ratingsDown + ratingsLeft + ratingsRight

	return scores, ratings
}
