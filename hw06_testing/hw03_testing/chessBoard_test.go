package board_test

import (
	"fmt"
	"testing"

	board "github.com/AlexSH61/homework_basic/hw06_testing/hw03_testing"
)

func TestBoardSiz(t *testing.T) {
	tests := []struct {
		size int
		want string
	}{
		{size: 1, want: "# \n"},
		{size: 3, want: "#  # \n #  \n#  # \n"},
		{size: 4, want: "#  #  \n #  # \n#  #  \n #  # \n"},
	}
	for _, size := range tests {
		t.Run(fmt.Sprintf("size%d", size.size), func(t *testing.T) {
			result := board.Boardsize(size.size)
			if result != size.want {
				t.Errorf("size %d: %q %q", size.size, size.want, result)
			}
		})
	}
}
