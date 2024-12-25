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
					Row:    94,
					Column: 34,
				},
				ButtonB: &util.Coordinate{
					Row:    22,
					Column: 67,
				},
				Prize: &util.Coordinate{
					Row:    8400,
					Column: 5400,
				},
			},
			{
				ButtonA: &util.Coordinate{
					Row:    26,
					Column: 66,
				},
				ButtonB: &util.Coordinate{
					Row:    67,
					Column: 21,
				},
				Prize: &util.Coordinate{
					Row:    12748,
					Column: 12176,
				},
			},
			{
				ButtonA: &util.Coordinate{
					Row:    17,
					Column: 86,
				},
				ButtonB: &util.Coordinate{
					Row:    84,
					Column: 37,
				},
				Prize: &util.Coordinate{
					Row:    7870,
					Column: 6450,
				},
			},
			{
				ButtonA: &util.Coordinate{
					Row:    69,
					Column: 23,
				},
				ButtonB: &util.Coordinate{
					Row:    27,
					Column: 71,
				},
				Prize: &util.Coordinate{
					Row:    18641,
					Column: 10279,
				},
			},
		},
	})
}
