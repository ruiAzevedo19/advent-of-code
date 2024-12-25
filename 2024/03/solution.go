package solution

import (
	"advent-of-code/challenge"
	"advent-of-code/util"
	"regexp"
	"strconv"
)

func init() {
	challenge.Register("2024", "3", MullItOver)
}

func MullItOver(filePath string) (result challenge.Result, err error) {
	corruptedMemory, err := util.ParseLine(filePath)
	if err != nil {
		return nil, err
	}

	summedMultiplications, err := mullItOver(corruptedMemory)
	if err != nil {
		return nil, err
	}
	summedMultiplicationsDo, err := mullItOverDo(corruptedMemory)
	if err != nil {
		return nil, err
	}

	return &Result{
		SummedMultiplications:   summedMultiplications,
		SummedMultiplicationsDo: summedMultiplicationsDo,
	}, nil
}

// mulRe is the regex to capture multiplications, e.g., mul(1,2).
var mulRe = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func mullItOver(corruptedMemory string) (summedMultiplications int, err error) {
	for _, match := range mulRe.FindAllStringSubmatch(corruptedMemory, -1) {
		if len(match) != 3 {
			continue
		}

		x, err := strconv.Atoi(match[1])
		if err != nil {
			return 0, err
		}
		y, err := strconv.Atoi(match[2])
		if err != nil {
			return 0, err
		}

		summedMultiplications += x * y
	}

	return summedMultiplications, nil
}

// mulDoRe is the regex to capture multiplications and do-instructions.
var mulDoRe = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)

func mullItOverDo(corruptedMemory string) (summedMultiplications int, err error) {
	isEnabled := true
	for _, match := range mulDoRe.FindAllStringSubmatch(corruptedMemory, -1) {
		if match[0] == "do()" {
			isEnabled = true
		} else if match[0] == "don't()" {
			isEnabled = false
		} else if len(match) == 3 && isEnabled {
			x, err := strconv.Atoi(match[1])
			if err != nil {
				return 0, err
			}
			y, err := strconv.Atoi(match[2])
			if err != nil {
				return 0, err
			}

			summedMultiplications += x * y
		}
	}

	return summedMultiplications, nil
}
