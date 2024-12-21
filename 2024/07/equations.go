package solution

import "strconv"

type Equation struct {
	TestValue int
	Numbers   []int
}

func NewEquation(testValue int, numbers []int) (equation *Equation) {
	return &Equation{
		TestValue: testValue,
		Numbers:   numbers,
	}
}

func (e *Equation) IsValid(operatorPermutations [][]Operator) (ok bool) {
	for _, operatorPermutation := range operatorPermutations {
		testValue := e.Numbers[0]
		for i := 1; i < len(e.Numbers); i++ {
			if operatorPermutation[i-1] == OperationMul {
				testValue *= e.Numbers[i]
			} else if operatorPermutation[i-1] == OperationSum {
				testValue += e.Numbers[i]
			} else {
				testValue = concateDigits(testValue, e.Numbers[i])
			}
		}
		if e.TestValue == testValue {
			return true
		}
	}

	return false
}

func concateDigits(x int, y int) (result int) {
	xx := strconv.Itoa(x)
	yy := strconv.Itoa(y)
	// It's safe to ignore the error because "xx" and "yy" are numbers.
	r, _ := strconv.Atoi(xx + yy)

	return r
}
