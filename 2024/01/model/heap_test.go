package model

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHeapify(t *testing.T) {
	type testCase struct {
		Name string

		Input []int
	}

	isSortedAscending := func(s []int) (ok bool) {
		for i := 0; i < len(s)-1; i++ {
			if s[i] > s[i+1] {
				return false
			}
		}

		return true
	}

	random := func(n int) (r []int) {
		for i := 0; i < n; i++ {
			r = append(r, rand.Int())
		}

		return r
	}

	validateOrder := func(t *testing.T, tc *testCase) {
		t.Run(tc.Name, func(t *testing.T) {
			heap := Heapify(tc.Input)

			var values []int
			for range tc.Input {
				value, err := heap.Pop()
				require.NoError(t, err)
				values = append(values, value)
			}
			assert.True(t, isSortedAscending(values))
		})
	}

	validateOrder(t, &testCase{
		Name: "Small",

		Input: random(20),
	})
	validateOrder(t, &testCase{
		Name: "Large",

		Input: random(1000000),
	})
}
