package board

import (
	"strings"
)

func Boardsize(size int) string {
	sb := &strings.Builder{}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i+j)%2 == 0 {
				sb.WriteString("#")
			}
			sb.WriteString(" ")
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
