package model

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCartesianProduct(t *testing.T) {
	type testCase struct {
		N int

		ExpectedOperations [][]Operator
	}

	validate := func(t *testing.T, tc *testCase) {
		name := fmt.Sprintf("N=%d", tc.N)
		t.Run(name, func(t *testing.T) {
			actualOperations := CartesianProduct(tc.N)

			assert.Equal(t, len(actualOperations), int(math.Pow(float64(len(allOperators)), float64(tc.N))))
			assert.ElementsMatch(t, tc.ExpectedOperations, actualOperations)
		})
	}

	validate(t, &testCase{
		N: 1,

		ExpectedOperations: [][]Operator{
			{OperationMul},
			{OperationSum},
			{OperationCombine},
		},
	})
	validate(t, &testCase{
		N: 2,

		ExpectedOperations: [][]Operator{
			{OperationMul, OperationMul},
			{OperationMul, OperationSum},
			{OperationMul, OperationCombine},
			{OperationSum, OperationMul},
			{OperationSum, OperationSum},
			{OperationSum, OperationCombine},
			{OperationCombine, OperationMul},
			{OperationCombine, OperationSum},
			{OperationCombine, OperationCombine},
		},
	})
}
