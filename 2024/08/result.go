package solution

import (
	"fmt"
	"strings"
)

type Result struct {
	NumberOfAntinodeLocations int
}

func (r *Result) String() (result string) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Number of antinode locations: %d\n", r.NumberOfAntinodeLocations))

	return sb.String()
}
