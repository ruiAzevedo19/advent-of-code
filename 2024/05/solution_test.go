package solution

import (
	"path/filepath"
	"testing"

	challengetesting "advent-of-code/challenge/testing"
)

func TestPrintQueue(t *testing.T) {
	validate := func(t *testing.T, tc *challengetesting.TestCaseChallenge) {
		tc.Validate(t)
	}

	validate(t, &challengetesting.TestCaseChallenge{
		Name: "Simple",

		FilePath: filepath.Join("testdata", "simple.txt"),
		Solution: PrintQueue,

		ExpectedResult: &Result{
			ValidUpdates: [][]int{
				{75, 47, 61, 53, 29},
				{97, 61, 53, 29, 13},
				{75, 29, 13},
			},
			InvalidUpdates: [][]int{
				{75, 97, 47, 61, 53},
				{61, 13, 29},
				{97, 13, 75, 29, 47},
			},
			FixedUpdates: [][]int{
				{97, 75, 47, 61, 53},
				{61, 29, 13},
				{97, 75, 47, 29, 13},
			},
			MiddlePageSumValidUpdates: 143,
			MiddlePageSumFixedUpdates: 123,
		},
	})
}
