package solution

import (
	"advent-of-code/challenge"
	"math"
)

func init() {
	challenge.Register("2024", "1", HistorianHysteria)
}

func HistorianHysteria(filePath string) (result challenge.Result, err error) {
	locations, err := Parse(filePath)
	if err != nil {
		return nil, err
	}
	numberOfLocations := len(locations.GroupOneLocationIDs)

	groupOneHeap := Heapify(locations.GroupOneLocationIDs)
	groupTwoHeap := Heapify(locations.GroupTwoLocationIDs)

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
