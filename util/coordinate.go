package util

type Coordinate struct {
	X int
	Y int
}

func (c *Coordinate) Equals(coordinate *Coordinate) bool {
	return c.X == coordinate.X && c.Y == coordinate.Y
}

func (c *Coordinate) IsValid(grid [][]rune) bool {
	if len(grid) == 0 {
		return false
	}

	if c.X < 0 || c.X >= len(grid) {
		return false
	}
	if c.Y < 0 || c.Y >= len(grid[0]) {
		return false
	}

	return true
}
