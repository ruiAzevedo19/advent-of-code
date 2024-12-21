package solution

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlutonianPebbles(t *testing.T) {
	type testCase struct {
		Stones []int
		Blinks int

		ExpectedStones []int
	}

	validate := func(t *testing.T, tc *testCase) {
		name := fmt.Sprintf("blinks=%d", tc.Blinks)
		t.Run(name, func(t *testing.T) {
			actualStones := plutonianPebbles(tc.Stones, tc.Blinks)

			assert.Equal(t, tc.ExpectedStones, actualStones)
		})
	}

	validate(t, &testCase{
		Stones:         []int{125, 17},
		Blinks:         1,
		ExpectedStones: []int{253000, 1, 7},
	})
	validate(t, &testCase{
		Stones:         []int{125, 17},
		Blinks:         2,
		ExpectedStones: []int{253, 0, 2024, 14168},
	})
	validate(t, &testCase{
		Stones:         []int{125, 17},
		Blinks:         3,
		ExpectedStones: []int{512072, 1, 20, 24, 28676032},
	})
	validate(t, &testCase{
		Stones:         []int{125, 17},
		Blinks:         4,
		ExpectedStones: []int{512, 72, 2024, 2, 0, 2, 4, 2867, 6032},
	})
	validate(t, &testCase{
		Stones:         []int{125, 17},
		Blinks:         5,
		ExpectedStones: []int{1036288, 7, 2, 20, 24, 4048, 1, 4048, 8096, 28, 67, 60, 32},
	})
	validate(t, &testCase{
		Stones:         []int{125, 17},
		Blinks:         6,
		ExpectedStones: []int{2097446912, 14168, 4048, 2, 0, 2, 4, 40, 48, 2024, 40, 48, 80, 96, 2, 8, 6, 7, 6, 0, 3, 2},
	})
}
