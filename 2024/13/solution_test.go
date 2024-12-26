package solution

import (
	"path/filepath"
	"testing"

	challengetesting "advent-of-code/challenge/testing"
)

func TestClawContraption(t *testing.T) {
	validate := func(t *testing.T, tc *challengetesting.TestCaseChallenge) {
		tc.Validate(t)
	}

	validate(t, &challengetesting.TestCaseChallenge{
		Name: "Simple",

		FilePath: filepath.Join("testdata", "simple.txt"),
		Solution: ClawContraption,

		ExpectedResult: &Result{
			Tokens: 480,
		},
	})
}