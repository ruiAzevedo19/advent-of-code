package solution

import (
	"advent-of-code/challenge"
)

func init() {
	challenge.Register("2024", "5", PrintQueue)
}

func PrintQueue(filePath string) (result challenge.Result, err error) {
	parsedRules, updates, err := Parse(filePath)
	if err != nil {
		return nil, err
	}
	rules := NewRules(parsedRules)

	middlePageSum := func(updates [][]int) (sum int) {
		for _, update := range updates {
			sum += update[len(update)/2]
		}

		return sum
	}

	var validUpdates, invalidUpdates [][]int
	for _, update := range updates {
		if !rules.UpdateIsValid(update) {
			invalidUpdates = append(invalidUpdates, update)
		} else {
			validUpdates = append(validUpdates, update)
		}
	}
	middlePageSumValidUpdates := middlePageSum(validUpdates)

	var fixedUpdates [][]int
	for _, update := range invalidUpdates {
		fixedUpdates = append(fixedUpdates, rules.FixUpdate(update))
	}
	middlePageSumFixedUpdates := middlePageSum(fixedUpdates)

	return &Result{
		ValidUpdates:              validUpdates,
		InvalidUpdates:            invalidUpdates,
		MiddlePageSumValidUpdates: middlePageSumValidUpdates,

		FixedUpdates:              fixedUpdates,
		MiddlePageSumFixedUpdates: middlePageSumFixedUpdates,
	}, nil
}
