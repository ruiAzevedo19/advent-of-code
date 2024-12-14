package model

import "math"

type Operator string

var allOperators []Operator

func registerOperator(operator string) Operator {
	o := Operator(operator)
	allOperators = append(allOperators, o)

	return o
}

var (
	OperationCombine = registerOperator("||")
	OperationMul     = registerOperator("*")
	OperationSum     = registerOperator("+")
)

func CartesianProduct(n int) (permutations [][]Operator) {
	// The number of elements is "numberOfOperators**n".
	numberOfElements := int(math.Pow(float64(len(allOperators)), float64(n)))

	for _, operator := range allOperators {
		permutations = append(permutations, []Operator{operator})
	}

	for range n - 1 {
		for _, permutation := range permutations {
			permutations = append(permutations, permute(allOperators, permutation)...)
		}
	}

	return permutations[len(permutations)-numberOfElements:]
}

func permute(operators []Operator, permutations []Operator) (result [][]Operator) {
	for _, operator := range operators {
		result = append(result, append([]Operator{operator}, permutations...))
	}

	return result
}
