package solution

import (
	"path/filepath"
	"testing"

	challengetesting "advent-of-code/challenge/testing"
)

func TestHistoriansHysteria(t *testing.T) {
	validate := func(t *testing.T, tc *challengetesting.TestCaseChallenge) {
		tc.Validate(t)
	}

	validate(t, &challengetesting.TestCaseChallenge{
		Name: "Simple",

		FilePath: filepath.Join("testdata", "simple.txt"),
		Solution: HistorianHysteria,

		ExpectedResult: &Result{
			TotalDistance:   11,
			SimilarityScore: 31,
		},
	})
}
