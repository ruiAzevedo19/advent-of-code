package solution

import (
	"advent-of-code/util"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	type testCase struct {
		Name string

		FilePath string

		ExpectedMachines []*Machine
	}

	validate := func(t *testing.T, tc *testCase) {
		t.Run(tc.Name, func(t *testing.T) {
			actualMachines, actualErr := Parse(tc.FilePath)
			require.NoError(t, actualErr)

			assert.Equal(t, tc.ExpectedMachines, actualMachines)
		})
	}

	validate(t, &testCase{
		Name: "Simple",

		FilePath: filepath.Join("testdata", "simple.txt"),

		ExpectedMachines: []*Machine{
			{
				ButtonA: &util.Coordinate{
					X: 94,
					Y: 34,
				},
				ButtonB: &util.Coordinate{
					X: 22,
					Y: 67,
				},
				Prize: &util.Coordinate{
					X: 8400,
					Y: 5400,
				},
			},
			{
				ButtonA: &util.Coordinate{
					X: 26,
					Y: 66,
				},
				ButtonB: &util.Coordinate{
					X: 67,
					Y: 21,
				},
				Prize: &util.Coordinate{
					X: 12748,
					Y: 12176,
				},
			},
			{
				ButtonA: &util.Coordinate{
					X: 17,
					Y: 86,
				},
				ButtonB: &util.Coordinate{
					X: 84,
					Y: 37,
				},
				Prize: &util.Coordinate{
					X: 7870,
					Y: 6450,
				},
			},
			{
				ButtonA: &util.Coordinate{
					X: 69,
					Y: 23,
				},
				ButtonB: &util.Coordinate{
					X: 27,
					Y: 71,
				},
				Prize: &util.Coordinate{
					X: 18641,
					Y: 10279,
				},
			},
		},
	})
}
