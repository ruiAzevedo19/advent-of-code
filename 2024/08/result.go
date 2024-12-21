package solution

import (
	"fmt"
	"strings"
)

type Result struct {
	NumberOfAntinodeLocations             int
	NumberOfAntinodeLocationsUpdatedModel int
}

func (r *Result) String() (result string) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Number of antinode locations: %d\n", r.NumberOfAntinodeLocations))
	sb.WriteString(fmt.Sprintf("Number of antinode locations with updated model: %d\n", r.NumberOfAntinodeLocationsUpdatedModel))

	return sb.String()
}
