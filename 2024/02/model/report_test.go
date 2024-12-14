package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReportIsSafe(t *testing.T) {
	type testCase struct {
		Name string

		Report *Report

		ExpectedOk        bool
		ExpectedFailIndex int
	}

	validate := func(t *testing.T, tc *testCase) {
		t.Run(tc.Name, func(t *testing.T) {
			actualOk := tc.Report.IsSafe()

			assert.Equal(t, tc.ExpectedOk, actualOk)
		})
	}

	t.Run("Invalid", func(t *testing.T) {
		validate(t, &testCase{
			Name: "Not all descending order",

			Report: &Report{
				values: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 2},
			},

			ExpectedOk: false,
		})
		validate(t, &testCase{
			Name: "Not all ascending order",

			Report: &Report{
				values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 8},
			},

			ExpectedOk: false,
		})
		validate(t, &testCase{
			Name: "Two adjacent values are the same",

			Report: &Report{
				values: []int{1, 2, 3, 3, 4, 5},
			},

			ExpectedOk: false,
		})
		validate(t, &testCase{
			Name: "Two adjacent values differ more than three",

			Report: &Report{
				values: []int{1, 2, 3, 10},
			},

			ExpectedOk: false,
		})
	})
	t.Run("Valid", func(t *testing.T) {
		validate(t, &testCase{
			Name: "Single",

			Report: &Report{
				values: []int{1},
			},

			ExpectedOk: true,
		})
		validate(t, &testCase{
			Name: "Ascending",

			Report: &Report{
				values: []int{1, 2, 5, 6, 7, 10},
			},

			ExpectedOk: true,
		})
		validate(t, &testCase{
			Name: "Descending",

			Report: &Report{
				values: []int{9, 8, 5, 4, 1},
			},

			ExpectedOk: true,
		})
	})
}
