package solution

import (
	"advent-of-code/util"
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var robotRE = regexp.MustCompile(`p=(.*),(.*) v=(.*),(.*)`)

var ErrMalformedData = errors.New("malformed data")

func Parse(filePath string) (robots []*Robot, err error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	content := strings.TrimSpace(string(fileContent))

	for _, line := range strings.Split(content, "\n") {
		matches := robotRE.FindStringSubmatch(line)
		if len(matches) != 5 {
			return nil, ErrMalformedData
		}

		px, err := strconv.Atoi(matches[1])
		if err != nil {
			return nil, err
		}
		py, err := strconv.Atoi(matches[2])
		if err != nil {
			return nil, err
		}
		vx, err := strconv.Atoi(matches[3])
		if err != nil {
			return nil, err
		}
		vy, err := strconv.Atoi(matches[4])
		if err != nil {
			return nil, err
		}

		robot := &Robot{
			Position: &util.Coordinate{
				X: px,
				Y: py,
			},
			Velocity: &util.Coordinate{
				X: vx,
				Y: vy,
			},
		}
		robots = append(robots, robot)
	}

	return robots, nil
}
