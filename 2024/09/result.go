package solution

import (
	"fmt"
	"strings"
)

type Result struct {
	Checksum          int
	ChecksumCompacted int
}

func (r *Result) String() (result string) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Checksum: %d\n", r.Checksum))
	sb.WriteString(fmt.Sprintf("Checksum compacted: %d\n", r.ChecksumCompacted))

	return sb.String()
}
