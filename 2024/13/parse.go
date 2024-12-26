package solution

import (
	"advent-of-code/util"
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var ErrWrongMachineEntries = errors.New("there must be three entries for each machine")

var (
	buttonARe = buttonRe("Button A")
	buttonBRe = buttonRe("Button B")
	prizeRe   = regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)
)

func buttonRe(button string) *regexp.Regexp {
	return regexp.MustCompile(button + `: X\+(\d+), Y\+(\d+)`)
}

func Parse(filePath string) (machines []*Machine, err error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	content := strings.TrimSpace(string(fileContent))

	for _, line := range strings.Split(content, "\n\n") {
		entries := strings.Split(line, "\n")
		if len(entries) != 3 {
			return nil, ErrWrongMachineEntries
		}

		machine := &Machine{}
		machine.ButtonA, err = extractCoordinates(buttonARe, entries[0])
		if err != nil {
			return nil, err
		}
		machine.ButtonB, err = extractCoordinates(buttonBRe, entries[1])
		if err != nil {
			return nil, err
		}
		machine.Prize, err = extractCoordinates(prizeRe, entries[2])
		if err != nil {
			return nil, err
		}
		machines = append(machines, machine)
	}

	return machines, nil
}

func extractCoordinates(re *regexp.Regexp, entry string) (coordinate *util.Coordinate, err error) {
	matches := re.FindStringSubmatch(entry)

	row, err := strconv.Atoi(matches[1])
	if err != nil {
		return nil, err
	}
	column, err := strconv.Atoi(matches[2])
	if err != nil {
		return nil, err
	}

	return &util.Coordinate{
		X: row,
		Y: column,
	}, nil
}
