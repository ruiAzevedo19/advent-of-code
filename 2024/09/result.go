package solution

import (
	"fmt"
	"strings"
)

type Result struct {
	Checksum int
}

func (r *Result) String() (result string) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Checksum: %d\n", r.Checksum))

	return sb.String()
}
