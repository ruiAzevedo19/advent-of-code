package challenge

import (
	"fmt"
)

// allChallenges holds all the challenge solutions grouped by year and day.
var allChallenges = map[string]map[string]Solution{}

// Register regists a new challenge.
func Register(year string, day string, solution Solution) {
	_, ok := allChallenges[year]
	if !ok {
		allChallenges[year] = map[string]Solution{}
	}

	if _, ok = allChallenges[year][day]; ok {
		panic(fmt.Sprintf("There is already a registered solution for %s/%s", year, day))
	}

	allChallenges[year][day] = solution
}

// SolutionForYearDay returns the solution for the given year and day.
func SolutionForYearDay(year string, day string) (solution Solution, err error) {
	challengesByDay, ok := allChallenges[year]
	if !ok {
		return nil, fmt.Errorf("there is no solutions for year %s", year)
	}

	solution, ok = challengesByDay[day]
	if !ok {
		return nil, fmt.Errorf("there is no solution for %s/%s", year, day)
	}

	return solution, nil
}
