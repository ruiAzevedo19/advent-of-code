package parse

import (
	"errors"
	"os"
	"print-queue/model"
	"strconv"
	"strings"
)

var ErrInvalidRuleSyntax = errors.New("invalid rule syntax")

// Parse parses the given file and returns the rules and updates.
func Parse(filePath string) (rules []*model.Rule, updates [][]int, err error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, nil, err
	}
	lines := strings.Split(string(fileContent), "\n")

	rules = []*model.Rule{}
	updates = [][]int{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "|") {
			rule, err := handleRule(line)
			if err != nil {
				return nil, nil, err
			}
			rules = append(rules, rule)
		} else if strings.Contains(line, ",") {
			update, err := handleUpdate(line)
			if err != nil {
				return nil, nil, err
			}
			updates = append(updates, update)
		}
	}

	return rules, updates, nil
}

func handleRule(line string) (rule *model.Rule, err error) {
	values := strings.Split(line, "|")
	if len(values) != 2 {
		return nil, ErrInvalidRuleSyntax
	}

	x, err := strconv.Atoi(strings.TrimSpace(values[0]))
	if err != nil {
		return nil, err
	}
	y, err := strconv.Atoi(strings.TrimSpace(values[1]))
	if err != nil {
		return nil, err
	}
	rule = model.NewRule(x, y)

	return rule, nil
}

func handleUpdate(line string) (update []int, err error) {
	values := strings.Split(line, ",")

	update = []int{}
	for _, value := range values {
		value = strings.TrimSpace(value)
		v, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		update = append(update, v)

	}

	return update, nil
}
