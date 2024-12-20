package solution

import (
	"advent-of-code/2024/11/parse"
	"strconv"
)

func PlutonianPebbles(filePath string) (result *Result, err error) {
	stones, err := parse.Parse(filePath)
	if err != nil {
		return nil, err
	}

	stones25 := plutonianPebbles(stones, 25)
	numberOfStones75 := plutonianPebblesCached(stones, 75)

	return &Result{
		NumberOfStones25: len(stones25),
		NumberOfStones75: numberOfStones75,
	}, nil
}

// REMARK This works for a small number of blinks but not for larger ones since it is space inefficient.
// Check "plutonianPebblesCached" for an efficient version.
func plutonianPebbles(stones []int, blinks int) (changedStones []int) {
	for range blinks {
		changedStones = []int{}
		for _, stone := range stones {
			if stone == 0 {
				changedStones = append(changedStones, 1)
			} else if leftStone, rightStone, ok := hasEvenNumberOfDigits(stone); ok {
				changedStones = append(changedStones, leftStone, rightStone)
			} else {
				changedStones = append(changedStones, stone*2024)
			}
			stones = changedStones
		}
	}

	return stones
}

// plutonianPebblesCached is the cached version of "plutonianPebbles".
func plutonianPebblesCached(stones []int, blinks int) (numberOfStones int) {
	cache := NewCache()

	for _, stone := range stones {
		numberOfStones += blink(cache, stone, blinks)
	}

	return numberOfStones
}

func blink(cache *Cache, stone int, blinks int) (numberOfStones int) {
	if blinks == 0 {
		return 1
	}

	numberOfStones, ok := cache.isCached(stone, blinks)
	if ok {
		return numberOfStones
	}

	if stone == 0 {
		numberOfStones = blink(cache, 1, blinks-1)
	} else if leftStone, rightStone, ok := hasEvenNumberOfDigits(stone); ok {
		numberOfStones = blink(cache, leftStone, blinks-1) + blink(cache, rightStone, blinks-1)
	} else {
		numberOfStones = blink(cache, stone*2024, blinks-1)
	}
	cache.Add(stone, blinks, numberOfStones)

	return numberOfStones
}

func hasEvenNumberOfDigits(x int) (leftStone int, rightStone int, ok bool) {
	if len(strconv.Itoa(x))%2 != 0 {
		return 0, 0, false
	}

	xx := strconv.Itoa(x)
	leftStone, _ = strconv.Atoi(xx[:len(xx)/2])
	rightStone, _ = strconv.Atoi(xx[len(xx)/2:])

	return leftStone, rightStone, true
}
