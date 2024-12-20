package solution

import (
	"advent-of-code/2024/01/model"
	"advent-of-code/2024/01/parse"
	"math"
)

func HistorianHysteria(filePath string) (result *Result, err error) {
	locations, err := parse.Parse(filePath)
	if err != nil {
		return nil, err
	}
	numberOfLocations := len(locations.GroupOneLocationIDs)

	groupOneHeap := model.Heapify(locations.GroupOneLocationIDs)
	groupTwoHeap := model.Heapify(locations.GroupTwoLocationIDs)

	totalDistance := 0
	for i := 0; i < numberOfLocations; i++ {
		groupOneLocationID, err := groupOneHeap.Pop()
		if err != nil {
			return nil, err
		}
		groupTwoLocationID, err := groupTwoHeap.Pop()
		if err != nil {
			return nil, err
		}

		totalDistance += int(math.Abs(float64(groupOneLocationID) - float64(groupTwoLocationID)))
	}

	frequencies := map[int]int{}
	for _, value := range locations.GroupTwoLocationIDs {
		frequencies[value]++
	}

	similarityScore := 0
	for _, value := range locations.GroupOneLocationIDs {
		similarityScore += value * frequencies[value]
	}

	return &Result{
		TotalDistance:   totalDistance,
		SimilarityScore: similarityScore,
	}, nil
}
