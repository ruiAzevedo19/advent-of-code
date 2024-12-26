package solution

import (
	"advent-of-code/challenge"
	"advent-of-code/util"
)

func init() {
	challenge.Register("2024", "14", RestroomRedoubt)
}

type Robot struct {
	Position *util.Coordinate
	Velocity *util.Coordinate
}

func RestroomRedoubt(filePath string) (result challenge.Result, err error) {
	robots, err := Parse(filePath)
	if err != nil {
		return nil, err
	}

	safetyFactor := restroomRedoubt(robots, 101, 103, 100)

	return &Result{
		SafetyFactor: safetyFactor,
	}, nil
}

func restroomRedoubt(robots []*Robot, width int, height int, executionTime int) int {
	// Holds the number of robots for each quadrant.
	quadrants := map[int]int{}

	for _, robot := range robots {
		lastPosition := &util.Coordinate{X: robot.Position.X, Y: robot.Position.Y}
		moveRobot(robot, lastPosition, width, height, executionTime)
		if quadrant := quadrant(lastPosition, width, height); quadrant != -1 {
			quadrants[quadrant]++
		}
	}

	return safetyFactor(quadrants)
}

func moveRobot(robot *Robot, lastPosition *util.Coordinate, width int, height int, executionTime int) {
	if executionTime == 0 {
		return
	}

	lastPosition.X = (lastPosition.X + robot.Velocity.X + width) % width
	lastPosition.Y = (lastPosition.Y + robot.Velocity.Y + height) % height

	moveRobot(robot, lastPosition, width, height, executionTime-1)
}

func quadrant(coordinate *util.Coordinate, width int, height int) (quadrant int) {
	middleX := width / 2
	middleY := height / 2

	if coordinate.X < middleX {
		if coordinate.Y < middleY {
			return 1
		} else if coordinate.Y >= height-middleY {
			return 3
		}
	} else if coordinate.X >= width-middleX {
		if coordinate.Y < middleY {
			return 2
		} else if coordinate.Y >= height-middleY {
			return 4
		}
	}

	return -1
}

func safetyFactor(quadrants map[int]int) (safetyFactor int) {
	safetyFactor = 1
	for _, n := range quadrants {
		safetyFactor *= n
	}

	return safetyFactor
}
