package solution

import (
	"advent-of-code/util"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRestroomRedoubt(t *testing.T) {
	type testCase struct {
		Name string

		FilePath      string
		Width         int
		Height        int
		ExecutionTime int

		ExpectedSafetyFactor int
	}
	validate := func(t *testing.T, tc *testCase) {
		t.Run(tc.Name, func(t *testing.T) {
			robots, err := Parse(tc.FilePath)
			require.NoError(t, err)

			actualSafetyFactor := restroomRedoubt(robots, tc.Width, tc.Height, tc.ExecutionTime)

			assert.Equal(t, tc.ExpectedSafetyFactor, actualSafetyFactor)
		})
	}

	validate(t, &testCase{
		Name: "Simple",

		FilePath:      filepath.Join("testdata", "simple.txt"),
		Width:         11,
		Height:        7,
		ExecutionTime: 100,

		ExpectedSafetyFactor: 12,
	})
}

func TestMoveRobot(t *testing.T) {
	type testCase struct {
		Name string

		Robot         *Robot
		Width         int
		Height        int
		ExecutionTime int

		ExpectedPosition *util.Coordinate
	}

	validate := func(t *testing.T, tc *testCase) {
		t.Run(tc.Name, func(t *testing.T) {
			lastPosition := &util.Coordinate{X: tc.Robot.Position.X, Y: tc.Robot.Position.Y}
			moveRobot(tc.Robot, lastPosition, tc.Width, tc.Height, tc.ExecutionTime)

			assert.Equal(t, tc.ExpectedPosition, lastPosition)
		})
	}

	validate(t, &testCase{
		Name: "Simple",

		Robot: &Robot{
			Position: &util.Coordinate{
				X: 2,
				Y: 4,
			},
			Velocity: &util.Coordinate{
				X: 2,
				Y: -3,
			},
		},
		Width:         11,
		Height:        7,
		ExecutionTime: 5,

		ExpectedPosition: &util.Coordinate{
			X: 1,
			Y: 3,
		},
	})
}
