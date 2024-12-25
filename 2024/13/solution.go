package solution

import (
	"advent-of-code/challenge"
	"advent-of-code/util"
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

	return &Result{
		Tokens: clawContraption(machines),
	}, nil
}

func clawContraption(machines []*Machine) (tokens int) {
LOOP:
	for _, machine := range machines {
		for a := 0; a < 100; a++ {
			for b := 0; b < 100; b++ {
				x := a*machine.ButtonA.Row + b*machine.ButtonB.Row
				y := a*machine.ButtonA.Column + b*machine.ButtonB.Column

				if x == machine.Prize.Row && y == machine.Prize.Column {
					tokens += 3*a + b

					continue LOOP
				}
			}
		}
	}

	return tokens
}
