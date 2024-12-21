package solution

import (
	"advent-of-code/challenge"
)

func init() {
	challenge.Register("2024", "7", BridgeRepair)
}

func BridgeRepair(filePath string) (result challenge.Result, err error) {
	equations, err := Parse(filePath)
	if err != nil {
		return nil, err
	}

	// Cache permutations to reduce execution time.
	cachePermutations := map[int][][]Operator{}
	var totalCalibrationResult int
	for _, equation := range equations {
		n := len(equation.Numbers) - 1
		permutations, ok := cachePermutations[n]
		if !ok {
			permutations = CartesianProduct(n)
			cachePermutations[n] = permutations
		}

		if equation.IsValid(permutations) {
			totalCalibrationResult += equation.TestValue
		}
	}

	return &Result{
		TotalCalibrationResult: totalCalibrationResult,
	}, nil
}
