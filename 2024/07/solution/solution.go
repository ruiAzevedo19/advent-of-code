package solution

import (
	"advent-of-code/2024/07/model"
	"advent-of-code/2024/07/parse"
	"advent-of-code/challenge"
)

func init() {
	challenge.Register("2024", "7", BridgeRepair)
}

func BridgeRepair(filePath string) (result challenge.Result, err error) {
	equations, err := parse.Parse(filePath)
	if err != nil {
		return nil, err
	}

	// Cache permutations to reduce execution time.
	cachePermutations := map[int][][]model.Operator{}
	var totalCalibrationResult int
	for _, equation := range equations {
		n := len(equation.Numbers) - 1
		permutations, ok := cachePermutations[n]
		if !ok {
			permutations = model.CartesianProduct(n)
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
