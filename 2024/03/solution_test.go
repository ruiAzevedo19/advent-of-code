package solution

import (
	"path/filepath"
	"testing"

	challengetesting "advent-of-code/challenge/testing"
)

func TestMullItOver(t *testing.T) {
	validate := func(t *testing.T, tc *challengetesting.TestCaseChallenge) {
		tc.Validate(t)
	}

	validate(t, &challengetesting.TestCaseChallenge{
		Name: "Simple",

		FilePath: filepath.Join("testdata", "simple.txt"),
		Solution: MullItOver,

		ExpectedResult: &Result{
			SummedMultiplications:   161,
			SummedMultiplicationsDo: 161,
		},
	})
	validate(t, &challengetesting.TestCaseChallenge{
		Name: "Simple Do-instructions",

		FilePath: filepath.Join("testdata", "simple-do.txt"),
		Solution: MullItOver,

		ExpectedResult: &Result{
			SummedMultiplications:   161,
			SummedMultiplicationsDo: 48,
		},
	})
}
