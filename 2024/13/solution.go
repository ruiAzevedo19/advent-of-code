package solution

import (
	"advent-of-code/challenge"
	"advent-of-code/util"
	"math"
)

func init() {
	challenge.Register("2024", "13", ClawContraption)
}

type Machine struct {
	ButtonA *util.Coordinate
	ButtonB *util.Coordinate
	Prize   *util.Coordinate
}

func ClawContraption(filePath string) (result challenge.Result, err error) {
	machines, err := Parse(filePath)
	if err != nil {
		return nil, err
	}

	tokens := clawContraption(machines)

	for _, machine := range machines {
		machine.Prize.X += 10000000000000
		machine.Prize.Y += 10000000000000
	}
	tokensHigherPrize := clawContraptionCramerRule(machines)

	return &Result{
		Tokens:            tokens,
		TokensHigherPrize: tokensHigherPrize,
	}, nil
}

func clawContraption(machines []*Machine) (tokens int) {
LOOP:
	for _, machine := range machines {
		for a := 0; a < 100; a++ {
			for b := 0; b < 100; b++ {
				x := a*machine.ButtonA.X + b*machine.ButtonB.X
				y := a*machine.ButtonA.Y + b*machine.ButtonB.Y

				if x == machine.Prize.X && y == machine.Prize.Y {
					tokens += 3*a + b

					continue LOOP
				}
			}
		}
	}

	return tokens
}

func clawContraptionCramerRule(machines []*Machine) (tokens int) {
	for _, machine := range machines {
		ax, ay := float64(machine.ButtonA.X), float64(machine.ButtonA.Y)
		bx, by := float64(machine.ButtonB.X), float64(machine.ButtonB.Y)
		px, py := float64(machine.Prize.X), float64(machine.Prize.Y)

		det := ax*by - ay*bx
		if det == 0.0 {
			continue
		}

		buttonA := (px*by - py*bx) / det
		buttonB := (ax*py - ay*px) / det

		if buttonA == math.Trunc(buttonA) && buttonB == math.Trunc(buttonB) {
			tokens += int(3*buttonA + buttonB)
		}
	}

	return tokens
}
