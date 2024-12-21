package solution

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

var (
	ErrInvalidInput        = errors.New("invalid input")
	ErrListsDoNotMatchSize = errors.New("lists do not match size")
)

func Parse(filePath string) (locations *Locations, err error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(fileContent), "\n")

	locations = NewLocations()
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		locationIDs := strings.Fields(line)
		if len(locationIDs) != 2 {
			return nil, ErrInvalidInput
		}

		locationID, err := strconv.Atoi(locationIDs[0])
		if err != nil {
			return nil, err
		}
		locations.AddGroupOneLocationID(locationID)

		locationID, err = strconv.Atoi(locationIDs[1])
		if err != nil {
			return nil, err
		}
		locations.AddGroupTwoLocationID(locationID)
	}

	if len(locations.GroupOneLocationIDs) != len(locations.GroupTwoLocationIDs) {
		return nil, ErrListsDoNotMatchSize
	}

	return locations, nil
}
